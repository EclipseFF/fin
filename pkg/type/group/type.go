package group

import "strings"

type Group struct {
	value string
}

func (g Group) String() string {
	return g.value
}

func New(group string) *Group {
	return &Group{
		value: getGroup(group),
	}
}

func (g Group) Equal(group Group) bool {
	return g.value == group.value
}

func (g Group) IsEmpty() bool {
	return len(strings.TrimSpace(g.value)) == 0
}
