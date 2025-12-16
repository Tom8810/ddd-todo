package dbmapper

import (
	"github.com/ddd-todo/project-backend/domain/dmodel/agg"
	"github.com/ddd-todo/project-backend/domain/dmodel/vo"
	"github.com/ddd-todo/project-backend/infrastructure/model"
)

func ToDBUser(a *agg.User) *model.User {
	return &model.User{
		ID:       agg.GetID(a),
		Name:     agg.GetName(a),
		Email:    agg.GetEmail(a),
		Password: agg.GetPassword(a),
	}
}

func ToAggUser(m *model.User) (*agg.User, error) {
	if m == nil {
		return nil, nil
	}
	return agg.ReconstructUser(
		vo.UserID(m.ID),
		vo.Name(m.Name),
		vo.Email(m.Email),
		vo.Password(m.Password),
		m.CreatedAt,
		m.UpdatedAt,
	), nil
}
