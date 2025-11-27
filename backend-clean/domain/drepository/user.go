package drepository

import (
	"context"

	"github.com/ddd-todo/project-backend/domain/dmodel/agg"
	"github.com/ddd-todo/project-backend/domain/dmodel/vo"
)

type UserRepository interface {
	Save(ctx context.Context, user *agg.User) error
	FindByEmail(ctx context.Context, email vo.Email) (*agg.User, error)
	FindByID(ctx context.Context, id vo.UserID) (*agg.User, error)
	ExistsByEmail(ctx context.Context, email vo.Email) (bool, error)
	Delete(ctx context.Context, id vo.UserID) error
}
