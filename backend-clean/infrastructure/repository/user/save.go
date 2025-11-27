package repository

import (
	"context"
	"fmt"

	"github.com/ddd-todo/project-backend/domain/dmodel/agg"
	"github.com/ddd-todo/project-backend/port/dbport"
)

func (r *userRepositoryImpl) Save(ctx context.Context, user *agg.User) error {
	dbUser := dbport.ToDBUser(user)

	if err := r.db.WithContext(ctx).Save(dbUser).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepositoryMock) Save(ctx context.Context, user *agg.User) error {
	fmt.Println("Mock UserRepository.Save called")
	return nil
}
