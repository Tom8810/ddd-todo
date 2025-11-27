package domainport

import (
	"github.com/ddd-todo/project-backend/application/usecase/todo/dto"
	"github.com/ddd-todo/project-backend/domain/dmodel/agg"
	"github.com/ddd-todo/project-backend/domain/drepository"
	domainport "github.com/ddd-todo/project-backend/port/domainport/common"
)

func ToDomainFilterAndPage(input dto.ListTodosInput) (drepository.TodoFilter, drepository.PageInput, error) {
	dFilter, err := ToDomainTodoFilter(input.Filter)
	if err != nil {
		return drepository.TodoFilter{}, drepository.PageInput{}, err
	}

	dPage, err := ToDomainPageInput(input.Page)
	if err != nil {
		return drepository.TodoFilter{}, drepository.PageInput{}, err
	}
	return dFilter, dPage, nil
}

func ToDtoListTodosOutput(todos []*agg.Todo, totalCount int, pageInput dto.TodoPageInput) dto.ListTodosOutput {
	todoList := make([]dto.TodoOutput, len(todos))
	for i, dTodo := range todos {
		todoList[i] = ToDtoTodoOutput(dTodo)
	}

	return dto.ListTodosOutput{
		Todos:      todoList,
		PagingInfo: domainport.ToDtoPagingInfoOutput(pageInput, totalCount),
	}
}
