package models

import "time"

type Packaging struct {
	Id        string
	Size      int
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
