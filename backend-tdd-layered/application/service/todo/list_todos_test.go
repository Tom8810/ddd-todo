package service_test

import (
	"context"
	"testing"

	service "github.com/ddd-todo/project-backend/application/service/todo"
	"github.com/ddd-todo/project-backend/graph/graphmodel"
	"github.com/ddd-todo/project-backend/infrastructure/repository"
	"github.com/ddd-todo/project-backend/internal/lib"
)

func TestListTodos(t *testing.T) {
	repo := repository.NewRepositoriesMock()
	service := service.NewTodoService(repo)
	ctx := context.Background()

	filter := graphmodel.TodoFilter{
		Keyword: lib.StringPtr("First"),
	}
	nonExistentFilter := graphmodel.TodoFilter{
		Keyword: lib.StringPtr("NonExistent"),
	}
	listRepoErrFilter := graphmodel.TodoFilter{
		Keyword: lib.StringPtr("ListError"),
	}
	countRepoErrFilter := graphmodel.TodoFilter{
		Keyword: lib.StringPtr("CountError"),
	}
	invalidFilter := graphmodel.TodoFilter{
		Status: lib.StringPtr("INVALID_STATUS"),
	}

	pageInput := graphmodel.PageInput{
		Limit:         1,
		Offset:        0,
		SortDirection: graphmodel.SortDirectionAsc,
	}
	errSortDirectionPageInput := graphmodel.PageInput{
		Limit:         1,
		Offset:        0,
		SortDirection: "INVALID_DIRECTION",
	}
	errSortKeyPageInput := graphmodel.PageInput{
		Limit:         1,
		Offset:        0,
		SortKey:       lib.StringPtr("INVALID_KEY"),
		SortDirection: graphmodel.SortDirectionAsc,
	}

	tests := []struct {
		name  string
		input struct {
			filter    graphmodel.TodoFilter
			pageInput graphmodel.PageInput
		}
		wantErr bool
		wantNum int
	}{
		{
			name: "正常系: フィルターなしでTodoリストを取得できること",
			input: struct {
				filter    graphmodel.TodoFilter
				pageInput graphmodel.PageInput
			}{
				filter:    graphmodel.TodoFilter{},
				pageInput: pageInput,
			},
			wantErr: false,
			wantNum: 2,
		},
		{
			name: "正常系: フィルターありでTodoリストを取得できること",
			input: struct {
				filter    graphmodel.TodoFilter
				pageInput graphmodel.PageInput
			}{
				filter:    filter,
				pageInput: pageInput,
			},
			wantErr: false,
			wantNum: 1,
		},
		{
			name: "正常系: 空のTodoリストを取得できること",
			input: struct {
				filter    graphmodel.TodoFilter
				pageInput graphmodel.PageInput
			}{
				filter:    nonExistentFilter,
				pageInput: pageInput,
			},
			wantErr: false,
			wantNum: 0,
		},
		{
			name: "異常系: listでエラーが発生した場合、エラーが返されること",
			input: struct {
				filter    graphmodel.TodoFilter
				pageInput graphmodel.PageInput
			}{
				filter:    listRepoErrFilter,
				pageInput: pageInput,
			},
			wantErr: true,
			wantNum: 0,
		},
		{
			name: "異常系: 無効なソート方向が指定された場合、エラーが返されること",
			input: struct {
				filter    graphmodel.TodoFilter
				pageInput graphmodel.PageInput
			}{
				filter:    filter,
				pageInput: errSortDirectionPageInput,
			},
			wantErr: true,
			wantNum: 0,
		},
		{
			name: "異常系: 無効なソートキーが指定された場合、エラーが返されること",
			input: struct {
				filter    graphmodel.TodoFilter
				pageInput graphmodel.PageInput
			}{
				filter:    filter,
				pageInput: errSortKeyPageInput,
			},
			wantErr: true,
			wantNum: 0,
		},
		{
			name: "異常系: 無効なフィルター条件が指定された場合、エラーが返されること",
			input: struct {
				filter    graphmodel.TodoFilter
				pageInput graphmodel.PageInput
			}{
				filter:    invalidFilter,
				pageInput: pageInput,
			},
			wantErr: true,
			wantNum: 0,
		},
		{
			name: "異常系: countでエラーが発生した場合、エラーが返されること",
			input: struct {
				filter    graphmodel.TodoFilter
				pageInput graphmodel.PageInput
			}{
				filter:    countRepoErrFilter,
				pageInput: pageInput,
			},
			wantErr: true,
			wantNum: 0,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := service.ListTodos(ctx, tt.input.filter, tt.input.pageInput)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListTodos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && len(got.Items) != tt.wantNum {
				t.Errorf("ListTodos() got num = %v, want %v", len(got.Items), tt.wantNum)
			}
		})
	}
}
