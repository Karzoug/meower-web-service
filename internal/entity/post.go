package entity

import (
	"time"
)

type Post struct {
	ID        string
	Text      string `validate:"required,min=1,max=280"`
	AuthorID  string
	IsDeleted bool
	UpdatedAt time.Time
}

type NewPost struct {
	AuthorID string
	Text     string `validate:"required,min=1,max=280"`
}
