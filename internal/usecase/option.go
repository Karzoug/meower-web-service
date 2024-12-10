package usecase

type PaginationOptions struct {
	Token string
	Size  int
}

type ReturnPostOptions struct {
	IncludeUser bool
}
