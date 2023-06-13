package video

import (
	"architecture_go/pkg/type/group"
	"architecture_go/services/admin/internal/domain/video/name"
	"github.com/pkg/errors"
	"time"
)

var (
	ErrGroupRequired = errors.New("Group is required")
)

type Video struct {
	id        uint
	createdAt time.Time

	name name.VideoName
}

func NewWithID(
	id uint,
	createdAt time.Time,
	name name.VideoName,
) (*Video, error) {

	return &Video{
		id:        id,
		createdAt: createdAt.UTC(),
		name:      name,
	}, nil
}

func New(
	id uint,
	group group.Group,
	name name.VideoName,
) (*Video, error) {

	if group.IsEmpty() {
		return nil, ErrGroupRequired
	}

	var timeNow = time.Now().UTC()
	return &Video{
		id:        id,
		createdAt: timeNow,
		name:      name,
	}, nil
}

func (v Video) ID() uint {
	return v.id
}

func (v Video) CreatedAt() time.Time {
	return v.createdAt
}

func (v Video) Name() name.VideoName {
	return v.name
}
