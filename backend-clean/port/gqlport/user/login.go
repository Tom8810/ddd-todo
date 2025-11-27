package gqlport

import (
	"github.com/ddd-todo/project-backend/application/usecase/user/dto"
	"github.com/ddd-todo/project-backend/graph/graphmodel"
)

func ToDtoUserLoginInput(input graphmodel.LoginInput) dto.LoginInput {
	return dto.LoginInput{
		Email:    input.Email,
		Password: input.Password,
	}
}

func ToGqlLoginResponse(output dto.LoginOutput) *graphmodel.LoginResponse {
	if output.User == nil {
		return &graphmodel.LoginResponse{}
	}
	user := ToGqlUser(*output.User)
	return &graphmodel.LoginResponse{
		Token: output.Token,
		User:  &user,
	}
}
