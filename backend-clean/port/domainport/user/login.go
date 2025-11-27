package domainport

import (
	"github.com/ddd-todo/project-backend/application/usecase/user/dto"
)

func ToDtoLoginOutput(token string, user *dto.UserOutput) dto.LoginOutput {
	return dto.LoginOutput{
		Token: token,
		User:  user,
	}
}
