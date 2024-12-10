package option

type Pagination struct {
	Token string
	Size  int
}

type ReturnPost struct {
	IncludeUser bool
}
