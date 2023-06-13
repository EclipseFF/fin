package dao

import (
	"time"
)

type Video struct {
	ID        uint      `db:"id"`
	CreatedAt time.Time `db:"created_at"`

	Name string `db:"name"`
}
