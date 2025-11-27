package drepository

import (
	"context"

	"github.com/ddd-todo/project-backend/domain/dmodel/agg"
	"github.com/ddd-todo/project-backend/domain/dmodel/vo"
)

type TodoRepository interface {
	Save(ctx context.Context, todo *agg.Todo) error
	FindByID(ctx context.Context, id vo.TodoID) (*agg.Todo, error)
	Delete(ctx context.Context, id vo.TodoID) error
	List(ctx context.Context, filter TodoFilter, pageInput PageInput) ([]*agg.Todo, error)
	Count(ctx context.Context, filter TodoFilter) (int, error)
}
