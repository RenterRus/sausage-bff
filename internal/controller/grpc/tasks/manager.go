package tasks

import (
	"context"
	"fmt"

	"github.com/RenterRus/sausage-bff/internal/entity"
	v1 "github.com/RenterRus/sausage-tasks/docs/proto/v1"
	"github.com/samber/lo"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// !!! Перевести tasks на cqrs

type tasksManager struct {
	client v1.TaskClient
}

func NewTasksController(client v1.TaskClient) Tasks {
	return &tasksManager{
		client: client,
	}
}

func (t *tasksManager) UpsertTask(ctx context.Context, input *entity.Task, user *entity.User) (string, error) {
	if input == nil || user == nil {
		return "", fmt.Errorf("UpsertTask: %w", entity.ErrNoRequiredParams)
	}

	resp, err := t.client.UpsertTask(ctx, &v1.UpsertTaskRequest{
		TransactionUuid: input.TransactionUUID,
		UserId:          user.UUID,
		Title:           input.Title,
		Comment:         input.Comment,
		Priority:        input.Priority,
		StartDate: &timestamppb.Timestamp{
			Seconds: input.StartDate.Unix(),
		},
		Complete: input.Complete,
	})
	if err != nil {
		return "", fmt.Errorf("UpsertTask: %w", err)
	}

	return resp.GetStatus(), nil
}

func (t *tasksManager) SetCompleteStatus(ctx context.Context, input *entity.Task, user *entity.User) (string, error) {
	if input == nil || user == nil {
		return "", fmt.Errorf("SetCompleteStatus: %w", entity.ErrNoRequiredParams)
	}

	resp, err := t.client.SetCompleteStatus(ctx, &v1.SetCompleteStatusRequest{
		Complete:        &input.Complete,
		TransactionUuid: input.TransactionUUID,
		UserId:          user.UUID,
	})
	if err != nil {
		return "", fmt.Errorf("SetCompleteStatus: %w", err)
	}

	return resp.GetStatus(), nil
}

// !!! Добавить курсорный пейджинг
func (t *tasksManager) SelectTasksByUser(ctx context.Context, user *entity.User) ([]entity.Task, error) {
	if user == nil {
		return nil, fmt.Errorf("SelectTasksByUser: %w", entity.ErrNoRequiredParams)
	}

	resp, err := t.client.SelectTasksByUser(ctx, &v1.SelectTasksRequest{
		UserId: user.UUID,
	})
	if err != nil {
		return nil, fmt.Errorf("SelectTasksByUser.SelectTasksByUser: %w", err)
	}

	return lo.Map(resp.Tasks, func(item *v1.TaskRow, _ int) entity.Task {
		return entity.Task{
			TransactionUUID: &item.TransactionUuid,
			Title:           item.GetTitle(),
			Comment:         item.Comment,
			Priority:        item.GetPriority(),
			StartDate:       item.GetStartDate().AsTime(),
			Complete:        item.GetComplete(),
			CreatedAt:       item.GetCreatedAt().AsTime(),
			UpdatedAt:       item.GetUpdatedAt().AsTime(),
		}
	}), nil
}

func (t *tasksManager) DeleteTaskByUser(ctx context.Context, input *entity.Task, user *entity.User) (string, error) {
	if input == nil || user == nil {
		return "", fmt.Errorf("DeleteTaskByUser: %w", entity.ErrNoRequiredParams)
	}

	resp, err := t.client.DeleteTaskByUser(ctx, &v1.DeleteTaskRequest{
		UserId:          user.UUID,
		TransactionUuid: input.TransactionUUID,
	})
	if err != nil {
		return "", fmt.Errorf("DeleteTaskByUser.DeleteTaskByUser: %w", err)
	}

	return resp.GetStatus(), nil
}
