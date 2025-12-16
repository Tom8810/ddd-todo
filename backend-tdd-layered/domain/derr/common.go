package derr

import "errors"

var (
	ErrInvalidSortDirection = errors.New("invalid sort direction")
)
