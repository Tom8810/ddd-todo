package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ddd-todo/project-backend/application/usecase"
	"github.com/ddd-todo/project-backend/controller"
	"github.com/ddd-todo/project-backend/graph"
	"github.com/ddd-todo/project-backend/infrastructure/auth"
	"github.com/ddd-todo/project-backend/infrastructure/database"
	"github.com/ddd-todo/project-backend/infrastructure/repository"
	"github.com/vektah/gqlparser/v2/ast"
)

const defaultPort = "8080"

func main() {
	db := database.NewConnection()

	authService := auth.NewAuthService()

	repos := repository.NewRepositoriesImpl(db)
	usecases := usecase.NewUsecases(repos, authService)
	resolvers := controller.NewResolver(usecases)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.New(graph.NewExecutableSchema(
		graph.Config{Resolvers: resolvers}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
