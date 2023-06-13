package name

import "github.com/pkg/errors"

var (
	MaxLength      = 50
	ErrWrongLength = errors.Errorf("name must be less or equal to %d characters", MaxLength)
)

type VideoName string

func (n VideoName) String() string {
	return string(n)
}

func New(name string) (*VideoName, error) {
	if len([]rune(name)) > MaxLength {
		return nil, ErrWrongLength
	}
	n := VideoName(name)
	return &n, nil
}
