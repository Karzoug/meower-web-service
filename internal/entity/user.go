package entity

type User struct {
	ID         string
	Username   string
	Name       string
	ImageUrl   string
	StatusText string
}

type UserProjection struct {
	ID         string
	Username   string
	Name       string
	ImageUrl   string
	StatusText string
}
