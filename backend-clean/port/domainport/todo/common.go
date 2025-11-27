package domainport

import (
	"github.com/ddd-todo/project-backend/application/usecase/todo/dto"
	"github.com/ddd-todo/project-backend/domain/dmodel/agg"
	"github.com/ddd-todo/project-backend/domain/dmodel/vo"
	"github.com/ddd-todo/project-backend/domain/drepository"
)

func ToDtoTodoOutput(a *agg.Todo) dto.TodoOutput {
	return dto.TodoOutput{
		ID:          a.GetID(),
		Title:       a.GetTitle(),
		Description: a.GetDescription(),
		Status:      a.GetStatus(),
		Deadline:    a.GetDeadline(),
		Priority:    a.GetPriority(),
		UserID:      a.GetUserID(),
		CreatedAt:   a.GetCreatedAt(),
		UpdatedAt:   a.GetUpdatedAt(),
	}
}

func ToDomainTodoFilter(input dto.TodoFilterInput) (drepository.TodoFilter, error) {
	var err error
	var dUserID vo.UserID
	if input.UserID != nil {
		dUserID, err = vo.NewUserID(*input.UserID)
		if err != nil {
			return drepository.TodoFilter{}, err
		}
	}

	var status vo.Status
	if input.Status != nil {
		status, err = vo.NewStatus(*input.Status)
		if err != nil {
			return drepository.TodoFilter{}, err
		}
	}

	var priority vo.Priority
	if input.Priority != nil {
		priority, err = vo.NewPriority(*input.Priority)
		if err != nil {
			return drepository.TodoFilter{}, err
		}
	}

	return drepository.NewTodoFilter(
		input.Keyword,
		&dUserID,
		&status,
		&priority,
		input.DeadlineAfter,
		input.DeadlineBefore,
	)
}

func ToDomainPageInput(input dto.TodoPageInput) (drepository.PageInput, error) {
	var sortField *drepository.TodoSortField
	if input.SortKey != nil {
		f := drepository.TodoSortField(*input.SortKey)
		sortField = &f
	}

	sortDirection, err := drepository.NewSortDirection(string(input.SortDirection))
	if err != nil {
		return drepository.PageInput{}, err
	}

	return drepository.PageInput{
		Sort:   drepository.NewTodoSort(sortField, sortDirection),
		Paging: drepository.NewPaging(input.Limit, input.Offset),
	}, nil
}
