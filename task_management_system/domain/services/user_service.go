package services

import (
	"context"
	"task_management_system/config"
	"task_management_system/domain/task_details"
	"task_management_system/domain/users"
	"task_management_system/dto/request"
	"task_management_system/errors"
)

//go:generate mockery --name IUsersService --inpackage --filename=users_service_mock.go
type IUsersService interface {
	AddUser(ctx context.Context, req *request.AddUserRequest) (id *string, err errors.IError)
	AssignUserTask(ctx context.Context, req *request.AssignUserTaskRequest) (err errors.IError)
}

type UsersService struct {
	config          *config.GeneralConfig
	taskDetailsRepo task_details.ITaskDetailsRepo
	usersRepo       users.IUsersRepo
}

func NewUsersService(
	cfg *config.GeneralConfig,
	tdr task_details.ITaskDetailsRepo,
	ur users.IUsersRepo,
) IUsersService {
	return &UsersService{
		config:          cfg,
		taskDetailsRepo: tdr,
		usersRepo:       ur,
	}
}

func (s *UsersService) AddUser(ctx context.Context, req *request.AddUserRequest) (*string, errors.IError) {
	id, err := s.taskDetailsRepo.AddTaskDetails(ctx, nil)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func (s *UsersService) AssignUserTask(ctx context.Context, req *request.AssignUserTaskRequest) errors.IError {
	taskRec, err := s.taskDetailsRepo.GetTaskDetailsById(ctx, req.TaskId)
	if err != nil {
		return err
	}
	userRec, err := s.usersRepo.GetUserById(ctx, req.UserId)
	if err != nil {
		return err
	}
	if !*userRec.IsDeleted {
		taskRec.UserId = userRec.Id
		err := s.taskDetailsRepo.UpdateTaskDetails(ctx, taskRec)
		if err != nil {
			return err
		}
	}

	return nil
}
