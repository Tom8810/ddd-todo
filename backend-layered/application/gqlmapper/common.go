package gqlmapper

import (
	"github.com/ddd-todo/project-backend/domain/drepository"
	"github.com/ddd-todo/project-backend/graph/graphmodel"
)

func ToAggPageInput(m *graphmodel.PageInput) (drepository.PageInput, error) {
	var sortField *drepository.TodoSortField
	var err error
	if m.SortKey != nil && *m.SortKey != "" {
		sortField, err = drepository.NewTodoSortField(*m.SortKey)
		if err != nil {
			return drepository.PageInput{}, err
		}
	}

	sortDirection, err := drepository.NewSortDirection(m.SortDirection.String())
	if err != nil {
		return drepository.PageInput{}, err
	}

	sort := drepository.NewTodoSort(sortField, sortDirection)

	paging := drepository.NewPaging(int(m.Limit), int(m.Offset))

	return drepository.NewPageInput(sort, paging), nil
}
