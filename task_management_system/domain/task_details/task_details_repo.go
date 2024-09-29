package task_details

import (
	"context"

	"task_management_system/errors"
)

//go:generate mockery --name ITaskDetailsRepo --inpackage --filename=task_details_repo_mock.go
type ITaskDetailsRepo interface {
	GetTaskDetailsById(ctx context.Context, id string) (*TaskDetails, errors.IError)
	GetTaskDetailsByUserId(ctx context.Context, userId string) ([]TaskDetails, errors.IError)
	GetTaskDetailsByUserIdAndStatus(ctx context.Context, userId, status string) ([]TaskDetails, errors.IError)
	AddTaskDetails(ctx context.Context, data *TaskDetails) (*string, errors.IError)
	UpdateTaskDetails(ctx context.Context, data *TaskDetails) errors.IError
	DeleteTaskDetailsById(ctx context.Context, id string) errors.IError
}
