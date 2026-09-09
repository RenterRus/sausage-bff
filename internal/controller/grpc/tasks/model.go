package tasks

import (
	"context"

	"github.com/RenterRus/sausage-bff/internal/entity"
)

type Tasks interface {
	UpsertTask(ctx context.Context, input *entity.Task, user *entity.User) (string, error)
	SetCompleteStatus(ctx context.Context, input *entity.Task, user *entity.User) (string, error)
	SelectTasksByUser(ctx context.Context, user *entity.User) ([]entity.Task, error)
	DeleteTaskByUser(ctx context.Context, input *entity.Task, user *entity.User) (string, error)
}
