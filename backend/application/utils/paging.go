package application_utils

import "github.com/ddd-todo/project-backend/graph/graphmodel"

func GetPagingInfo(totalCount int, pageInput graphmodel.PageInput) *graphmodel.PagingInfo {
	hasNextPage := pageInput.Offset+pageInput.Limit < totalCount
	hasPreviousPage := pageInput.Offset > 0

	return &graphmodel.PagingInfo{
		TotalCount:      totalCount,
		HasNextPage:     hasNextPage,
		HasPreviousPage: hasPreviousPage,
		Offset:          pageInput.Offset,
		Limit:           pageInput.Limit,
	}
}
