package gqlport

import (
	"github.com/ddd-todo/project-backend/application/usecase/user/dto"
	"github.com/ddd-todo/project-backend/graph/graphmodel"
)

func ToDtoUserRegisterInput(input graphmodel.UserCreateInput) dto.RegisterInput {
	return dto.RegisterInput{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}
}
