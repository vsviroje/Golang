package request

import (
	"context"
	"time"

	"task_management_system/errors"
)

type GetTaskRequest struct {
	TaskId string `validate:"nonzero" `
	UserId string
	Status string
}

func (r *GetTaskRequest) Validate(ctx context.Context) errors.IError {
	if r.Status != "" && r.UserId == "" {
		return errors.NewErrParamInvalid("userId can not be empty if status is provided")
	}

	return ValidateFields(r)
}

type AddTaskRequest struct {
	Title       string    `json:"title" validate:"nonzero"`
	Description string    `json:"description" validate:"nonzero" `
	DueDate     time.Time `json:"dueDate" validate:"nonzero" `
	UserId      *string   `json:"userId" `
}

func (r *AddTaskRequest) Validate(ctx context.Context) errors.IError {
	return ValidateFields(r)
}

type UpdateTaskRequest struct {
	TaskId      string     `json:"taskId" validate:"nonzero"`
	Title       *string    `json:"title" `
	Description *string    `json:"description"`
	DueDate     *time.Time `json:"dueDate"  `
	UserId      *string    `json:"userId" `
}

func (r *UpdateTaskRequest) Validate(ctx context.Context) errors.IError {
	return ValidateFields(r)
}

type DeleteTaskRequest struct {
	TaskId string `validate:"nonzero" `
}

func (r *DeleteTaskRequest) Validate(ctx context.Context) errors.IError {
	return ValidateFields(r)
}
