package postgres

import (
	"architecture_go/services/admin/internal/domain/video"
	"architecture_go/services/admin/internal/domain/video/name"
	"architecture_go/services/admin/internal/repository/storage/postgres/dao"
)

func (r Repository) toDomainVideo(dao *dao.Video) (*video.Video, error) {

	result, err := video.NewWithID(
		dao.ID,
		dao.CreatedAt,
		name.VideoName(dao.Name),
	)
	if err != nil {
		return nil, err
	}
	return result, nil
}
