package gqlport

import "github.com/ddd-todo/project-backend/application/usecase/user/dto"

func ToDtoUserDeleteInput(userID string) dto.DeleteUserInput {
	return dto.DeleteUserInput{
		UserID: userID,
	}
}
