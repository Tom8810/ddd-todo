package usecase

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"github.com/ddd-todo/project-backend/application/usecase/user/dto"
	"github.com/ddd-todo/project-backend/domain/derr"
	"github.com/ddd-todo/project-backend/domain/dmodel/agg"
	"github.com/ddd-todo/project-backend/domain/dmodel/vo"
	domainport "github.com/ddd-todo/project-backend/port/domainport/user"
)

func (s *UserUsecase) Login(ctx context.Context, input dto.LoginInput) (dto.LoginOutput, error) {
	emailVO, err := vo.NewEmail(input.Email)
	if err != nil {
		return dto.LoginOutput{}, err
	}

	dUser, err := s.Repository.UserRepository.FindByEmail(ctx, emailVO)
	if err != nil {
		return dto.LoginOutput{}, err
	}
	if dUser == nil {
		return dto.LoginOutput{}, derr.ErrUserNotFound
	}

	if !s.verifyPassword(dUser, input.Password) {
		return dto.LoginOutput{}, derr.ErrInvalidPassword
	}

	token, err := s.AuthService.GenerateToken(dUser.ID.Value(), dUser.Email.Value())
	if err != nil {
		return dto.LoginOutput{}, err
	}

	user := domainport.ToDtoUserOutput(dUser)

	return domainport.ToDtoLoginOutput(token, &user), nil
}

func (s *UserUsecase) verifyPassword(user *agg.User, plainPassword string) bool {
	hashedPassword := user.Password.Value()
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}
