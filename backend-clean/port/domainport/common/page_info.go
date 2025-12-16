package domainport

import (
	common_dto "github.com/ddd-todo/project-backend/application/usecase/common_dto"
	"github.com/ddd-todo/project-backend/application/usecase/todo/dto"
)

func ToDtoPagingInfoOutput(pageInput dto.TodoPageInput, totalCount int) common_dto.PagingInfoOutput {
	hasNextPage := pageInput.Offset+pageInput.Limit < totalCount
	hasPreviousPage := pageInput.Offset > 0

	return common_dto.PagingInfoOutput{
		TotalCount:      totalCount,
		HasNextPage:     hasNextPage,
		HasPreviousPage: hasPreviousPage,
		Offset:          pageInput.Offset,
		Limit:           pageInput.Limit,
	}
}
