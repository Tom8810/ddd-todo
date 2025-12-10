package drepository

import (
	"time"

	"github.com/ddd-todo/project-backend/domain/derr"
	"github.com/ddd-todo/project-backend/domain/dmodel/vo"
)

type TodoFilter struct {
	Keyword  *string
	UserID   *vo.UserID
	Status   *vo.Status
	Priority *vo.Priority
	FromDate *time.Time
	ToDate   *time.Time
}

type PageInput struct {
	Sort   *TodoSort
	Paging *Paging
}

type TodoSort struct {
	Field     *TodoSortField
	Direction SortDirection
}

type TodoSortField string

const (
	TodoSortByCreatedAt TodoSortField = "created_at"
	TodoSortByDeadline  TodoSortField = "deadline"
	TodoSortByPriority  TodoSortField = "priority"
)

type SortDirection string

const (
	SortAsc  SortDirection = "ASC"
	SortDesc SortDirection = "DESC"
)

type Paging struct {
	Limit  int
	Offset int
}

func NewTodoFilter(keyword *string, userID *vo.UserID, status *vo.Status, priority *vo.Priority, fromDate, toDate *time.Time) (TodoFilter, error) {
	filter := TodoFilter{
		Keyword:  keyword,
		UserID:   userID,
		Status:   status,
		Priority: priority,
		FromDate: fromDate,
		ToDate:   toDate,
	}
	if err := filter.validateFilter(); err != nil {
		return TodoFilter{}, err
	}
	return filter, nil
}

func (f *TodoFilter) validateFilter() error {
	if f.FromDate != nil && f.ToDate != nil && f.FromDate.After(*f.ToDate) {
		return derr.ErrInvalidTodoFitlerDateRange
	}
	return nil
}

func NewPageInput(sort *TodoSort, paging *Paging) PageInput {
	var limit, offset int
	if paging != nil {
		limit = paging.Limit
		offset = paging.Offset
	} else {
		limit = 20
		offset = 0
	}
	return PageInput{
		Sort:   sort,
		Paging: NewPaging(limit, offset),
	}
}

func NewTodoSort(field *TodoSortField, direction SortDirection) *TodoSort {
	return &TodoSort{
		Field:     field,
		Direction: direction,
	}
}

func NewTodoSortField(value string) (*TodoSortField, error) {
	switch value {
	case string(TodoSortByCreatedAt):
		f := TodoSortByCreatedAt
		return &f, nil
	case string(TodoSortByDeadline):
		f := TodoSortByDeadline
		return &f, nil
	case string(TodoSortByPriority):
		f := TodoSortByPriority
		return &f, nil
	default:
		return nil, derr.ErrInvalidTodoSortField
	}
}

func NewSortDirection(value string) (SortDirection, error) {
	switch value {
	case string(SortAsc):
		d := SortAsc
		return d, nil
	case string(SortDesc):
		d := SortDesc
		return d, nil
	default:
		return "", derr.ErrInvalidSortDirection
	}
}

func NewPaging(limit, offset int) *Paging {
	if limit <= 0 {
		limit = 20
	}
	if offset < 0 {
		offset = 0
	}
	return &Paging{
		Limit:  limit,
		Offset: offset,
	}
}
