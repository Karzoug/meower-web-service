package entity

import (
	"time"

	"github.com/rs/xid"
)

type Post struct {
	ID        xid.ID
	Text      string `validate:"required,min=1,max=280"`
	AuthorID  xid.ID
	IsDeleted bool
	UpdatedAt time.Time
}

type NewPost struct {
	Text string `validate:"required,min=1,max=280"`
}
