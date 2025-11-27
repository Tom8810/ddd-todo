package domainport

import (
	"github.com/ddd-todo/project-backend/application/usecase/user/dto"
	"github.com/ddd-todo/project-backend/domain/dmodel/agg"
)

func ToDtoUserOutput(a *agg.User) dto.UserOutput {
	return dto.UserOutput{
		ID:    a.ID.Value(),
		Name:  a.Name.Value(),
		Email: a.Email.Value(),
	}
}
