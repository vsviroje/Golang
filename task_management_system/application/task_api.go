package application

import (
	"context"

	"task_management_system/dto/request"
	"task_management_system/dto/response"
	"task_management_system/errors"
)

//go:generate mockery --name ITaskApplication --inpackage --filename=task_application_mock.go
type ITaskApplication interface {
	GetTask(ctx context.Context, req *request.GetTaskRequest) (resp *response.GetTaskResponse, err errors.IError)
	AddTask(ctx context.Context, req *request.AddTaskRequest) (resp *response.AddTaskResponse, err errors.IError)
	UpdateTask(ctx context.Context, req *request.UpdateTaskRequest) (err errors.IError)
	DeleteTask(ctx context.Context, req *request.DeleteTaskRequest) (err errors.IError)
}
