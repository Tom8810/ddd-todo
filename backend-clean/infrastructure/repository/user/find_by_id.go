package repository

import (
	"context"
	"fmt"

	"github.com/ddd-todo/project-backend/domain/dmodel/agg"
	"github.com/ddd-todo/project-backend/domain/dmodel/vo"
	"github.com/ddd-todo/project-backend/infrastructure/query"
	"github.com/ddd-todo/project-backend/port/dbport"
)

func (r *userRepositoryImpl) FindByID(ctx context.Context, id vo.UserID) (*agg.User, error) {
	q := query.Use(r.db)

	dbUser, err := q.User.WithContext(ctx).Where(q.User.ID.Eq(id.Value())).First()
	if err != nil {
		return nil, err
	}

	return dbport.ToDomainUser(dbUser)
}

func (r *userRepositoryMock) FindByID(ctx context.Context, id vo.UserID) (*agg.User, error) {
	fmt.Println("Mock UserRepository.FindByID called")
	return &agg.User{}, nil
}
