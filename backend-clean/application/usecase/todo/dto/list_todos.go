package dto

import dto "github.com/ddd-todo/project-backend/application/usecase/common_dto"

type ListTodosInput struct {
	Filter TodoFilterInput
	Page   TodoPageInput
}

type ListTodosOutput struct {
	Todos      []TodoOutput
	PagingInfo dto.PagingInfoOutput
}
