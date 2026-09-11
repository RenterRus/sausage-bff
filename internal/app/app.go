package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	authClient "github.com/RenterRus/sausage-auth/docs/proto/v1"
	"github.com/RenterRus/sausage-bff/internal/controller/grpc/auth"
	"github.com/RenterRus/sausage-bff/internal/controller/grpc/tasks"
	general "github.com/RenterRus/sausage-bff/internal/graph/general"
	"github.com/RenterRus/sausage-bff/internal/graph/general/generated"
	tasksClient "github.com/RenterRus/sausage-tasks/docs/proto/v1"
	"github.com/vektah/gqlparser/v2/ast"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type App struct {
	conf      *Config
	tasksConn *grpc.ClientConn
	authConn  *grpc.ClientConn
}

func NewApp(configPath string) (*App, error) {
	lastSlash := 0
	for i, v := range configPath {
		if v == '/' {
			lastSlash = i
		}
	}

	conf, err := ReadConfig(configPath[:lastSlash], configPath[lastSlash+1:])
	if err != nil {
		return nil, fmt.Errorf("ReadConfig: %w", err)
	}

	return &App{
		conf: conf,
	}, nil
}

const defaultPort = "8080"

func (a *App) Run() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	var err error

	if a.tasksConn, err = grpc.NewClient(fmt.Sprintf("%s:%d", a.conf.Services.Tasks.Host, a.conf.Services.Tasks.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials())); err != nil {
		return fmt.Errorf("tasksConn: %v", err)
	}

	if a.authConn, err = grpc.NewClient(fmt.Sprintf("%s:%d", a.conf.Services.Auth.Host, a.conf.Services.Auth.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials())); err != nil {
		return fmt.Errorf("authConn: %v", err)
	}

	tasksClient := tasks.NewTasksController(tasksClient.NewTaskClient(a.tasksConn))
	_ = tasksClient

	authClient := auth.NewAuthController(authClient.NewAuthServiceClient(a.authConn))
	_ = authClient

	srv := handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: &general.Resolver{}}))

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

	return nil
}

func (a *App) Close() {
	a.tasksConn.Close()
	a.authConn.Close()
}
