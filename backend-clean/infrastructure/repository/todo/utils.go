package repository

import (
	"github.com/ddd-todo/project-backend/domain/drepository"
	"github.com/ddd-todo/project-backend/infrastructure/query"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

func buildWhereConditions(filter drepository.TodoFilter, q *query.Query) []gen.Condition {
	var conditions []gen.Condition

	if filter.Priority != nil {
		conditions = append(conditions, q.Todo.Where(q.Todo.Priority.Eq(filter.Priority.Value())))
	}

	if filter.UserID != nil {
		conditions = append(conditions, q.Todo.Where(q.Todo.UserId.Eq(filter.UserID.Value())))
	}

	if filter.Status != nil {
		conditions = append(conditions, q.Todo.Where(q.Todo.Status.Eq(string(*filter.Status))))
	}

	if filter.FromDate != nil {
		conditions = append(conditions, q.Todo.Where(q.Todo.Deadline.Gte(*filter.FromDate)))
	}

	if filter.ToDate != nil {
		conditions = append(conditions, q.Todo.Where(q.Todo.Deadline.Lte(*filter.ToDate)))
	}

	if filter.Keyword != nil && *filter.Keyword != "" {
		keyword := "%" + *filter.Keyword + "%"
		conditions = append(conditions,
			q.Todo.Where(
				q.Todo.Title.Like(keyword),
			).Or(
				q.Todo.Description.Like(keyword),
			),
		)
	}

	return conditions
}

func buildOrderBy(
	p drepository.PageInput,
	q *query.Query,
) field.Expr {
	if p.Sort == nil || p.Sort.Field == nil {
		return q.Todo.CreatedAt.Desc()
	}

	var f field.Field
	switch *p.Sort.Field {
	case drepository.TodoSortByCreatedAt:
		f = field.Field(q.Todo.CreatedAt)

	case drepository.TodoSortByDeadline:
		f = field.Field(q.Todo.Deadline)

	case drepository.TodoSortByPriority:
		f = field.Field(q.Todo.Priority)

	default:
		f = field.Field(q.Todo.CreatedAt)
	}

	if p.Sort.Direction == drepository.SortAsc {
		return f.Asc()
	}
	return f.Desc()
}

func applyPaging(query query.ITodoDo, p drepository.PageInput) query.ITodoDo {
	if p.Paging == nil {
		return query.Limit(20) // Default limit
	}

	query = query.Limit(p.Paging.Limit)
	if p.Paging.Offset > 0 {
		query = query.Offset(p.Paging.Offset)
	}

	return query
}
