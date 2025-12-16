package controller

import "github.com/ddd-todo/project-backend/application/usecase"

type Resolver struct {
	sevice *usecase.Usecases
}

func NewResolver(usecases *usecase.Usecases) *Resolver {
	return &Resolver{
		sevice: usecases,
	}
}
