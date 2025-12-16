package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/ddd-todo/project-backend/domain/dmodel/agg"
	"github.com/ddd-todo/project-backend/domain/dmodel/vo"
	"github.com/ddd-todo/project-backend/infrastructure/dbmapper"
	"github.com/ddd-todo/project-backend/infrastructure/query"
	"github.com/ddd-todo/project-backend/internal/lib"
)

func (r *userRepositoryImpl) FindByID(ctx context.Context, id vo.UserID) (*agg.User, error) {
	q := query.Use(r.db)

	dbUser, err := q.User.WithContext(ctx).Where(q.User.ID.Eq(id.Value())).First()
	if err != nil {
		return nil, err
	}

	return dbmapper.ToAggUser(dbUser)
}

func (r *userRepositoryMock) FindByID(ctx context.Context, id vo.UserID) (*agg.User, error) {
	fmt.Println("Mock UserRepository.FindByID called")

	notFoundId := "bb04f183-fb98-6d72-eae2-8633ed7a5c2d"
	errCaseId := "c9f374b7-7c97-2529-de08-f091c9c9921c"

	if id.Value() == notFoundId {
		return nil, nil
	}

	if id.Value() == errCaseId {
		return nil, fmt.Errorf("simulated repository error")
	}

	return &agg.User{
		ID:        vo.UserID(lib.GenerateID()),
		Name:      vo.Name("Mock User"),
		Email:     vo.Email("demo@example.com"),
		Password:  vo.Password("hashedpassword"),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
