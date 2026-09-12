package tasks

import (
	"context"
	"fmt"

	controllertasks "github.com/RenterRus/sausage-bff/internal/controller/grpc/tasks"
	"github.com/RenterRus/sausage-bff/internal/entity"
)

type tasksCase struct {
	api controllertasks.Tasks
}

func NewTaskCase(api controllertasks.Tasks) TasksCase {
	return &tasksCase{api: api}
}

// DeleteTaskByUser implements TasksCase.
func (t *tasksCase) DeleteTaskByUser(ctx context.Context, input *entity.Task, user *entity.User) (string, error) {
	if input == nil || user == nil {
		return "", fmt.Errorf("DeleteTaskByUser: %w", entity.ErrNoRequiredParams)
	}

	req, err := t.api.DeleteTaskByUser(ctx, input, user)
	if err != nil {
		return "", fmt.Errorf("DeleteTaskByUser.DeleteTaskByUser: %w", err)
	}

	return req, nil
}

// SelectTasksByUser implements TasksCase.
func (t *tasksCase) SelectTasksByUser(ctx context.Context, user *entity.User) ([]entity.Task, error) {
	if user == nil {
		return nil, fmt.Errorf("SelectTasksByUser: %w", entity.ErrNoRequiredParams)
	}

	resp, err := t.api.SelectTasksByUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("SelectTasksByUser.SelectTasksByUser: %w", err)
	}

	return resp, nil
}

// SetCompleteStatus implements TasksCase.
func (t *tasksCase) SetCompleteStatus(ctx context.Context, input *entity.Task, user *entity.User) (string, error) {
	if input == nil || user == nil {
		return "", fmt.Errorf("SetCompleteStatus: %w", entity.ErrNoRequiredParams)
	}

	req, err := t.api.SetCompleteStatus(ctx, input, user)
	if err != nil {
		return "", fmt.Errorf("SetCompleteStatus.SetCompleteStatus: %w", err)
	}

	return req, nil
}

// UpsertTask implements TasksCase.
func (t *tasksCase) UpsertTask(ctx context.Context, input *entity.Task, user *entity.User) (string, error) {
	if input == nil || user == nil {
		return "", fmt.Errorf("UpsertTask: %w", entity.ErrNoRequiredParams)
	}

	req, err := t.api.UpsertTask(ctx, input, user)
	if err != nil {
		return "", fmt.Errorf("UpsertTask.UpsertTask: %w", err)
	}

	return req, nil
}
