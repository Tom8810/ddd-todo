package gqlport

import (
	"github.com/ddd-todo/project-backend/application/usecase/todo/dto"
	"github.com/ddd-todo/project-backend/graph/graphmodel"
	gqlport "github.com/ddd-todo/project-backend/port/gqlport/common"
)

func ToDtoListTodosInput(filter graphmodel.TodoFilter, pageInput graphmodel.PageInput) dto.ListTodosInput {
	return dto.ListTodosInput{
		Filter: ToDtoTodoFilterInput(filter),
		Page:   ToDtoTodoPageInput(pageInput),
	}
}

func ToGqlTodoConnection(output dto.ListTodosOutput) *graphmodel.TodoConnection {
	items := make([]*graphmodel.Todo, len(output.Todos))
	for i, todo := range output.Todos {
		items[i] = ToGqlTodo(todo)
	}

	return &graphmodel.TodoConnection{
		Items:      items,
		PagingInfo: gqlport.ToGqlPagingInfo(output.PagingInfo),
	}
}
