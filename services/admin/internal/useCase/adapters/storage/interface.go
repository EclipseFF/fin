package storage

import (
	"architecture_go/pkg/type/context"
	"architecture_go/pkg/type/queryParameter"
	"architecture_go/services/admin/internal/domain/group"
	"architecture_go/services/admin/internal/domain/video"
)

type Storage interface {
	Group
	Video
}

type Video interface {
	DeleteVideo(ctx context.Context, ID uint) error

	VideoReader
}

type VideoReader interface {
	ListVideos(ctx context.Context, parameter queryParameter.QueryParameter) ([]*video.Video, error)
	GetVideoByID(ctx context.Context, ID uint) (response *video.Video)
}

type Group interface {
	CreateGroup(ctx context.Context, group *group.Group) (*group.Group, error)
	DeleteGroup(ctx context.Context, ID uint) error

	GroupReader
	VideosInGroup
}

type GroupReader interface {
	ListGroup(ctx context.Context, parameter queryParameter.QueryParameter) ([]*group.Group, error)
}

type VideosInGroup interface {
	CreateVideoIntoGroup(ctx context.Context, groupID uint, videos ...*video.Video) ([]*video.Video, error)
	DeleteVideoFromGroup(ctx context.Context, groupID, videoID uint) error
	AddVideoToGroup(ctx context.Context, groupID uint, videoID ...uint) error
}
