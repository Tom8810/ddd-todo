package gqlport

import (
	"github.com/ddd-todo/project-backend/application/usecase/user/dto"
	"github.com/ddd-todo/project-backend/graph/graphmodel"
)

func ToDtoUserUpdateInput(userID string, input string) dto.UserUpdateInput {
	return dto.UserUpdateInput{
		UserID:      userID,
		UpdateField: input,
	}
}

func ToGqlUser(user dto.UserOutput) graphmodel.User {
	return graphmodel.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
