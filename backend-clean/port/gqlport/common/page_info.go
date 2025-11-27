package gqlport

import (
	dto "github.com/ddd-todo/project-backend/application/usecase/common_dto"
	"github.com/ddd-todo/project-backend/graph/graphmodel"
)

func ToGqlPagingInfo(pagingInfo dto.PagingInfoOutput) *graphmodel.PagingInfo {
	return &graphmodel.PagingInfo{
		TotalCount:      pagingInfo.TotalCount,
		HasNextPage:     pagingInfo.HasNextPage,
		HasPreviousPage: pagingInfo.HasPreviousPage,
		Offset:          pagingInfo.Offset,
		Limit:           pagingInfo.Limit,
	}
}
