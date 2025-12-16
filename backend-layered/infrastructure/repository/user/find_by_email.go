package repository

import (
	"context"
	"fmt"

	"github.com/ddd-todo/project-backend/domain/dmodel/agg"
	"github.com/ddd-todo/project-backend/domain/dmodel/vo"
	"github.com/ddd-todo/project-backend/infrastructure/dbmapper"
	"github.com/ddd-todo/project-backend/infrastructure/query"
)

func (r *userRepositoryImpl) FindByEmail(ctx context.Context, email vo.Email) (*agg.User, error) {
	q := query.Use(r.db)

	dbUser, err := q.User.WithContext(ctx).Where(q.User.Email.Eq(email.Value())).First()
	if err != nil {
		return nil, err
	}

	return dbmapper.ToAggUser(dbUser)
}

func (r *userRepositoryMock) FindByEmail(ctx context.Context, email vo.Email) (*agg.User, error) {
	fmt.Println("Mock UserRepository.FindByEmail called")
	return &agg.User{}, nil
}
