package group

import (
	"architecture_go/services/admin/internal/domain/group/name"
	"time"
)

type Group struct {
	id        uint
	createdAt time.Time

	name name.GroupName
}

func NewWithID(id uint, createdAt time.Time, name name.GroupName) *Group {
	return &Group{
		id:        id,
		createdAt: createdAt.UTC(),
		name:      name,
	}
}

func New(id uint, name name.GroupName) *Group {
	var timeNow = time.Now().UTC()
	return &Group{
		id:        id,
		createdAt: timeNow,
		name:      name,
	}
}

func (g Group) ID() uint {
	return g.id
}

func (g Group) CreatedAt() time.Time {
	return g.createdAt
}

func (g Group) Name() name.GroupName {
	return g.name
}
