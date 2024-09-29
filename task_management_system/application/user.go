package application

import (
	"context"
	"log"
	"task_management_system/domain/services"
	"task_management_system/dto/request"
	"task_management_system/dto/response"
	"task_management_system/errors"
)

type UserApplication struct {
	usersService services.IUsersService
}

func NewUserApp(us services.IUsersService) IUsersApplication {
	return &UserApplication{
		usersService: us,
	}
}

func (s *UserApplication) AssignUserTask(ctx context.Context, req *request.AssignUserTaskRequest) errors.IError {
	const funcName = "UserApplication.AssignUserTask "
	err := s.usersService.AssignUserTask(ctx, req)
	if err != nil {
		log.Printf(funcName+"AssignUserTask Failed err: %s", err.Error())
		return err
	}
	return nil
}

func (s *UserApplication) AddUser(ctx context.Context, req *request.AddUserRequest) (*response.AddUsersResponse, errors.IError) {
	const funcName = "UserApplication.AddUser "
	id, err := s.usersService.AddUser(ctx, req)
	if err != nil {
		log.Printf(funcName+"AddUser Failed err: %s", err.Error())
		return nil, err
	}
	return &response.AddUsersResponse{UserId: *id}, nil
}
