package repository

import (
	"context"
	"fmt"

	"github.com/ddd-todo/project-backend/domain/dmodel/vo"
	"github.com/ddd-todo/project-backend/infrastructure/query"
)

func (r *userRepositoryImpl) ExistsByEmail(ctx context.Context, email vo.Email) (bool, error) {
	q := query.Use(r.db)

	count, err := q.User.WithContext(ctx).Where(q.User.Email.Eq(email.Value())).Count()
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *userRepositoryMock) ExistsByEmail(ctx context.Context, email vo.Email) (bool, error) {
	fmt.Println("Mock UserRepository.ExistsByEmail called")
	return false, nil
}
