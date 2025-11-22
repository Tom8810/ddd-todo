package gqlmapper

import (
	"github.com/ddd-todo/project-backend/domain/dmodel/agg"
	"github.com/ddd-todo/project-backend/domain/dmodel/vo"
	"github.com/ddd-todo/project-backend/graph/graphmodel"
)

func ToGqlUser(a *agg.User) *graphmodel.User {
	return &graphmodel.User{
		ID:       agg.GetID(a),
		Name:     agg.GetName(a),
		Email:    agg.GetEmail(a),
		Password: agg.GetPassword(a),
	}
}

func ToAggUser(m *graphmodel.User) *agg.User {
	return agg.NewUser(
		vo.UserID(m.ID),
		vo.Name(m.Name),
		vo.Email(m.Email),
		vo.Password(m.Password),
	)
}
