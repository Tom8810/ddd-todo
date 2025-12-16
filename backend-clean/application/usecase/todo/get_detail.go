package usecase

import (
	"context"

	"github.com/ddd-todo/project-backend/application/usecase/todo/dto"
	"github.com/ddd-todo/project-backend/domain/dmodel/vo"
	domainport "github.com/ddd-todo/project-backend/port/domainport/todo"
)

func (s *TodoUsecase) GetTodoDetail(ctx context.Context, input dto.GetTodoDetailInput) (dto.TodoOutput, error) {
	dID, err := vo.NewTodoID(input.TodoID)
	if err != nil {
		return dto.TodoOutput{}, err
	}

	todo, err := s.Repository.TodoRepository.FindByID(ctx, dID)
	if err != nil {
		return dto.TodoOutput{}, err
	}

	return domainport.ToDtoTodoOutput(todo), nil
}
