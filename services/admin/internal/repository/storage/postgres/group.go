package postgres

import (
	"architecture_go/pkg/tools/transaction"
	"architecture_go/pkg/type/columnCode"
	"architecture_go/pkg/type/context"
	log "architecture_go/pkg/type/logger"
	"architecture_go/pkg/type/queryParameter"
	"architecture_go/services/admin/internal/domain/group"
	"architecture_go/services/admin/internal/repository/storage/postgres/dao"
	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
)

var mappingSortGroup = map[columnCode.ColumnCode]string{
	"id":   "id",
	"name": "name",
}

func (r *Repository) CreateGroup(c context.Context, group *group.Group) (*group.Group, error) {

	ctx := c.CopyWithTimeout(r.options.Timeout)
	defer ctx.Cancel()

	query, args, err := r.genSQL.Insert("slurm.group").
		Columns(
			"id",
			"name",
			"created_at",
		).
		Values(
			group.ID(),
			group.Name().Value(),
			group.CreatedAt()).ToSql()
	if err != nil {
		return nil, log.ErrorWithContext(ctx, err)
	}

	if _, err = r.db.Exec(ctx, query, args...); err != nil {
		return nil, log.ErrorWithContext(ctx, err)
	}
	return group, nil
}

func (r *Repository) DeleteGroup(c context.Context, ID uint) error {

	ctx := c.CopyWithTimeout(r.options.Timeout)
	defer ctx.Cancel()

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return log.ErrorWithContext(ctx, err)
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, tx)

	if err = r.deleteGroupTx(ctx, tx, ID); err != nil {
		return err
	}

	return nil
}

func (r *Repository) deleteGroupTx(ctx context.Context, tx pgx.Tx, ID uint) error {
	query, args, err := r.genSQL.Update("slurm.group").
		Set("is_archived", true).
		Where(squirrel.Eq{
			"id":          ID,
			"is_archived": false,
		}).ToSql()

	if err != nil {
		return log.ErrorWithContext(ctx, err)
	}

	if _, err = tx.Exec(ctx, query, args...); err != nil {
		return log.ErrorWithContext(ctx, err)
	}

	if err = r.clearGroupTx(ctx, tx, ID); err != nil {
		return err
	}

	return nil
}

func (r *Repository) clearGroupTx(ctx context.Context, tx pgx.Tx, groupID uint) error {
	query, args, err := r.genSQL.
		Delete("slurm.contact_in_group").
		Where(squirrel.Eq{"group_id": groupID}).
		ToSql()
	if err != nil {
		return log.ErrorWithContext(ctx, err)
	}

	if _, err = tx.Exec(ctx, query, args...); err != nil {
		return log.ErrorWithContext(ctx, err)
	}

	return nil
}

func (r *Repository) ListGroup(c context.Context, parameter queryParameter.QueryParameter) ([]*group.Group, error) {

	ctx := c.CopyWithTimeout(r.options.Timeout)
	defer ctx.Cancel()

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, log.ErrorWithContext(ctx, err)
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, tx)

	response, err := r.listGroupTx(ctx, tx, parameter)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *Repository) listGroupTx(ctx context.Context, tx pgx.Tx, parameter queryParameter.QueryParameter) ([]*group.Group, error) {
	var result []*group.Group

	var builder = r.genSQL.Select(
		"id",
		"name",
		"created_at",
	).
		From("slurm.group")

	builder = builder.Where(squirrel.Eq{"is_archived": false})

	if len(parameter.Sorts) > 0 {
		builder = builder.OrderBy(parameter.Sorts.Parsing(mappingSortGroup)...)
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
	rows, err := tx.Query(ctx, query, args...)
	if err != nil {
		return nil, log.ErrorWithContext(ctx, err)
	}

	var groups []*dao.Group
	if err = pgxscan.ScanAll(&groups, rows); err != nil {
		return nil, log.ErrorWithContext(ctx, err)
	}

	return result, nil
}
