package repository

import (
	"context"
	"fmt"

	"github.com/ddd-todo/project-backend/domain/dmodel/agg"
	"github.com/ddd-todo/project-backend/domain/dmodel/vo"
	"github.com/ddd-todo/project-backend/infrastructure/dbmapper"
	"github.com/ddd-todo/project-backend/infrastructure/query"
	"github.com/ddd-todo/project-backend/internal/lib"
	"golang.org/x/crypto/bcrypt"
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

	validEmail := "demo@example.com"
	notfoundEmail := "notfound@example.com"

	validPassword := "password"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(validPassword), bcrypt.DefaultCost)

	if email.Value() == notfoundEmail {
		return nil, nil
	}

	if email.Value() == validEmail {
		return &agg.User{
			ID:       vo.UserID(lib.GenerateID()),
			Name:     vo.Name("Mock User"),
			Email:    vo.Email(validEmail),
			Password: vo.Password(hashedPassword),
		}, nil
	}

	return &agg.User{}, nil
}
