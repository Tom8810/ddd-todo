package usecase

import (
	"context"

	"github.com/ddd-todo/project-backend/application/usecase/user/dto"
	"github.com/ddd-todo/project-backend/domain/dmodel/vo"
	domainport "github.com/ddd-todo/project-backend/port/domainport/user"
)

func (s *UserUsecase) GetUserByID(ctx context.Context, input dto.GetUserByIDInput) (dto.UserOutput, error) {
	dID, err := vo.NewUserID(input.UserID)
	if err != nil {
		return dto.UserOutput{}, err
	}
	user, err := s.Repository.UserRepository.FindByID(ctx, dID)
	if err != nil {
		return dto.UserOutput{}, err
	}

	return domainport.ToDtoUserOutput(user), nil
}
