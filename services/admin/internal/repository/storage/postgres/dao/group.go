package dao

import (
	"time"
)

type Group struct {
	ID        uint      `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}
