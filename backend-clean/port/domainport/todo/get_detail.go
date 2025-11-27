package domainport

import (
	"github.com/ddd-todo/project-backend/domain/dmodel/vo"
)

func ToDomainTodoDetailInput(todoID string) (vo.TodoID, error) {
	return vo.NewTodoID(todoID)
}
