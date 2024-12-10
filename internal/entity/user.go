package entity

import (
	"github.com/rs/xid"
)

type User struct {
	ID         xid.ID
	Username   string
	Name       string
	ImageUrl   string
	StatusText string
}

type UserProjection struct {
	ID         xid.ID
	Username   string
	Name       string
	ImageUrl   string
	StatusText string
}
