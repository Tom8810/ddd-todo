package repository

import (
	"context"
	"fmt"

	"github.com/ddd-todo/project-backend/domain/dmodel/agg"
	"github.com/ddd-todo/project-backend/infrastructure/dbmapper"
)

func (r *userRepositoryImpl) Save(ctx context.Context, user *agg.User) error {
	dbUser := dbmapper.ToDBUser(user)

	if err := r.db.WithContext(ctx).Save(dbUser).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepositoryMock) Save(ctx context.Context, user *agg.User) error {
	fmt.Println("Mock UserRepository.Save called")

	notFoundId := "bb04f183-fb98-6d72-eae2-8633ed7a5c2d"
	errCaseId := "c9f374b7-7c97-2529-de08-f091c9c9921c"

	if user.ID.Value() == notFoundId {
		return fmt.Errorf("mocked repository error: user not found")
	}
	if user.ID.Value() == errCaseId {
		return fmt.Errorf("mocked repository error")
	}

	if user.Email.Value() == "existing@example.com" {
		return fmt.Errorf("mocked repository error: email already exists")
	}

	if user.Email.Value() == "error@example.com" {
		return fmt.Errorf("mocked repository error")
	}

	return nil
}
