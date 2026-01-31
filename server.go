package main

import (
	user_entity "learn_gqlgen/auth/entity"
	user_memory "learn_gqlgen/auth/repository/memory"
	user_service "learn_gqlgen/auth/services"
	"learn_gqlgen/graph"

	//"learn_gqlgen/todo/repository"
	todo_entity "learn_gqlgen/todo/entity"
	todo_memory "learn_gqlgen/todo/repository/memory"
	todo_service "learn_gqlgen/todo/services"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vektah/gqlparser/v2/ast"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	var users []*user_entity.User
	var todos []*todo_entity.Todo
	userRepository := user_memory.NewAuthRepository(users)
	userService := user_service.NewAuthService(userRepository)
	todoRepository := todo_memory.NewTodoRepository(todos)
	todoService := todo_service.NewTodoService(todoRepository)
	resolver := &graph.Resolver{
		TodoService: todoService,
		AuthService: userService,
	}
	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))
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
