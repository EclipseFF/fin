package postgres

import (
	"architecture_go/pkg/tools/transaction"
	"architecture_go/pkg/type/columnCode"
	"architecture_go/pkg/type/context"
	log "architecture_go/pkg/type/logger"
	"architecture_go/pkg/type/queryParameter"
	"architecture_go/services/admin/internal/domain/video"
	"architecture_go/services/admin/internal/repository/storage/postgres/dao"
	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/opentracing/opentracing-go"
	"time"
)

var mappingSortContact = map[columnCode.ColumnCode]string{
	"id":   "id",
	"name": "name",
}

func (r *Repository) DeleteVideo(c context.Context, ID uint) error {

	ctx := c.CopyWithTimeout(r.options.Timeout)
	defer ctx.Cancel()

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return log.ErrorWithContext(ctx, err)
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, tx)

	if err = r.deleteVideoTx(ctx, tx, ID); err != nil {
		return err
	}

	return nil
}

func (r *Repository) deleteVideoTx(ctx context.Context, tx pgx.Tx, ID uint) error {
	builder := r.genSQL.Update("arlan.video").
		Set("is_archived", true).
		Set("modified_at", time.Now().UTC()).
		Where(squirrel.Eq{"is_archived": false, "id": ID})

	query, args, err := builder.ToSql()
	if err != nil {
		return log.ErrorWithContext(ctx, err)
	}

	rows, err := tx.Query(ctx, query, args...)
	if err != nil {
		return log.ErrorWithContext(ctx, err)
	}

	var daoContacts []*dao.Video
	if err = pgxscan.ScanAll(&daoContacts, rows); err != nil {
		return log.ErrorWithContext(ctx, err)
	}

	return nil
}

func (r *Repository) ListContact(c context.Context, parameter queryParameter.QueryParameter) ([]*video.Video, error) {

	ctx := c.CopyWithTimeout(r.options.Timeout)
	defer ctx.Cancel()

	span, tmp := opentracing.StartSpanFromContext(c, "ListContact")
	defer span.Finish()
	ctx = context.New(tmp)

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, log.ErrorWithContext(ctx, err)
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, tx)

	if parameter.Pagination.Limit == 0 {
		parameter.Pagination.Limit = r.options.DefaultLimit
	}

	contacts, err := r.listContactTx(ctx, tx, parameter)
	if err != nil {
		return nil, err
	}

	return contacts, nil
}

func (r *Repository) listContactTx(ctx context.Context, tx pgx.Tx, parameter queryParameter.QueryParameter) ([]*video.Video, error) {
	var builder = r.genSQL.Select(
		"id",
		"created_at",
		"name",
	).From("slurm.contact")

	builder = builder.Where(squirrel.Eq{"is_archived": false})

	if len(parameter.Sorts) > 0 {
		builder = builder.OrderBy(parameter.Sorts.Parsing(mappingSortContact)...)
	} else {
		builder = builder.OrderBy("created_at DESC")
	}

	if parameter.Pagination.Limit > 0 {
		builder = builder.Limit(parameter.Pagination.Limit)
	}
	if parameter.Pagination.Offset > 0 {
		builder = builder.Offset(parameter.Pagination.Offset)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, log.ErrorWithContext(ctx, err)
	}

	rows, err := tx.Query(ctx, query, args...)
	if err != nil {
		return nil, log.ErrorWithContext(ctx, err)
	}

	var daoVideos []*dao.Video
	if err = pgxscan.ScanAll(&daoVideos, rows); err != nil {
		return nil, log.ErrorWithContext(ctx, err)
	}

	return nil, nil
}
