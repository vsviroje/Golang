package application

import (
	"context"

	"task_management_system/dto/request"
	"task_management_system/dto/response"
	"task_management_system/errors"
)

//go:generate mockery --name IUsersApplication --inpackage --filename=user_application_mock.go
type IUsersApplication interface {
	AssignUserTask(ctx context.Context, req *request.AssignUserTaskRequest) (err errors.IError)
	AddUser(ctx context.Context, req *request.AddUserRequest) (resp *response.AddUsersResponse, err errors.IError)
}
