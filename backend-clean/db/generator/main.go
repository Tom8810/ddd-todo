package main

import (
	"github.com/ddd-todo/project-backend/infrastructure/model"

	"gorm.io/gen"
)

type Querier interface {
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./infrastructure/query",
		Mode:    gen.WithoutContext | gen.WithQueryInterface,
	})

	g.ApplyBasic(model.Models...)

	g.Execute()
}
