package useCase

import (
	"architecture_go/pkg/type/context"
	"architecture_go/pkg/type/queryParameter"
	"architecture_go/services/admin/internal/domain/group"
	"architecture_go/services/admin/internal/domain/video"
	"github.com/google/uuid"
)

type Video interface {
	Delete(c context.Context, ID uint /*Тут можно передавать фильтр*/) error
	VideoReader
}

type VideoReader interface {
	List(c context.Context, parameter queryParameter.QueryParameter) ([]*video.Video, error)
}

type Group interface {
	Create(c context.Context, groupCreate *group.Group) (*group.Group, error)
	Delete(c context.Context, ID uuid.UUID /*Тут можно передавать фильтр*/) error

	GroupReader
}

type GroupReader interface {
	List(c context.Context, parameter queryParameter.QueryParameter) ([]*group.Group, error)
}

type VideosInGroup interface {
	CreateVideoIntoGroup(ctx context.Context, groupID uint, videos ...*video.Video) ([]*video.Video, error)
	DeleteVideoFromGroup(ctx context.Context, groupID, videoID uint) error
	AddVideoToGroup(ctx context.Context, groupID uint, videoID ...uint) error
}
