package usecase

import (
	"context"

	"github.com/ddd-todo/project-backend/application/usecase/todo/dto"
	domainport "github.com/ddd-todo/project-backend/port/domainport/todo"
)

func (s *TodoUsecase) ListTodos(ctx context.Context, filter dto.TodoFilterInput, pageInput dto.TodoPageInput) (dto.ListTodosOutput, error) {
	dFilter, err := domainport.ToDomainTodoFilter(filter)
	if err != nil {
		return dto.ListTodosOutput{}, err
	}
	dPaging, err := domainport.ToDomainPageInput(pageInput)
	if err != nil {
		return dto.ListTodosOutput{}, err
	}

	todos, err := s.Repository.TodoRepository.List(ctx, dFilter, dPaging)
	if err != nil {
		return dto.ListTodosOutput{}, err
	}

	count, err := s.Repository.TodoRepository.Count(ctx, dFilter)
	if err != nil {
		return dto.ListTodosOutput{}, err
	}

	return domainport.ToDtoListTodosOutput(todos, count, pageInput), nil
}
