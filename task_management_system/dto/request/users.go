package request

import (
	"context"

	"task_management_system/errors"
)

type AddUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email" validate:"nonzero" `
}

type AssignUserTaskRequest struct {
	UserId string `json:"userId" validate:"nonzero" `
	TaskId string `json:"taskId" validate:"nonzero" `
}

func (r *AddUserRequest) Validate(ctx context.Context) errors.IError {
	return ValidateFields(r)
}

func (r *AssignUserTaskRequest) Validate(ctx context.Context) errors.IError {
	return ValidateFields(r)
}
