package dto

type PagingInfoOutput struct {
	TotalCount      int
	HasNextPage     bool
	HasPreviousPage bool
	Offset          int
	Limit           int
}
