package gqlport

import "github.com/ddd-todo/project-backend/application/usecase/user/dto"

func ToDtoGetUserByIDInput(userID string) dto.GetUserByIDInput {
	return dto.GetUserByIDInput{
		UserID: userID,
	}
}
