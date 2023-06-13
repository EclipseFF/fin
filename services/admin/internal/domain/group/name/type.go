package name

import "github.com/pkg/errors"

var (
	MaxLength      = 50
	ErrWrongLength = errors.Errorf("name must be less than or equal to %d characters", MaxLength)
)

type GroupName struct {
	value string
}

func New(name string) (GroupName, error) {
	if len([]rune(name)) > MaxLength {
		return GroupName{}, ErrWrongLength
	}
	return GroupName{value: name}, nil
}

func (g GroupName) Value() string {
	return g.value
}
