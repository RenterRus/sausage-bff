package app

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	conf *Config
	pgx  *pgxpool.Pool
	s    *grpc.Server
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

func (a *App) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", a.conf.GRPC.Host, a.conf.GRPC.Port))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	a.s = grpc.NewServer()
	reflection.Register(a.s)

	defer func() {
		a.s.Stop()
	}()

	a.pgx, err = pgxpool.New(context.Background(), fmt.Sprintf("%s://%s:%s@%s:%d/%s",
		a.conf.PSQL.Provider, a.conf.PSQL.Username, a.conf.PSQL.Password,
		a.conf.PSQL.Host, a.conf.PSQL.Port, a.conf.PSQL.DBName))
	if err != nil {
		return fmt.Errorf("Run.NewDBManager: %w", err)
	}
	defer func() {
		a.pgx.Close()
	}()

	log.Printf("gRPC server listening on %s", fmt.Sprintf("%s:%d", a.conf.GRPC.Host, a.conf.GRPC.Port))

	if err := a.s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}

func (a *App) Close() {
	a.pgx.Close()

	a.s.Stop()
}
