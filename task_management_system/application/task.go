package application

import (
	"context"
	"task_management_system/domain/services"
	"task_management_system/dto/request"
	"task_management_system/dto/response"
	"task_management_system/errors"

	"log"
)

type TaskApplicaion struct {
	taskService services.ITaskDetailsService
}

func NewTaskApp(ts services.ITaskDetailsService) ITaskApplication {
	return &TaskApplicaion{
		taskService: ts,
	}
}

func (s *TaskApplicaion) GetTask(ctx context.Context, req *request.GetTaskRequest) (*response.GetTaskResponse, errors.IError) {
	const funcName = "TaskApplicaion.GetTask "
	_, err := s.taskService.GetTaskDetails(ctx, req)
	if err != nil {
		log.Printf(funcName+"GetTaskDetails Failed err: %s", err.Error())
		return nil, err
	}
	return &response.GetTaskResponse{}, nil
}

func (s *TaskApplicaion) AddTask(ctx context.Context, req *request.AddTaskRequest) (*response.AddTaskResponse, errors.IError) {
	const funcName = "TaskApplicaion.AddTask "
	id, err := s.taskService.AddTaskDetails(ctx, req)
	if err != nil {
		log.Printf(funcName+"AddTaskDetails Failed err: %s", err.Error())
		return nil, err
	}
	return &response.AddTaskResponse{TaskId: *id}, nil
}

func (s *TaskApplicaion) UpdateTask(ctx context.Context, req *request.UpdateTaskRequest) errors.IError {
	const funcName = "TaskApplicaion.UpdateTask "
	err := s.taskService.UpdateTaskDetails(ctx, req)
	if err != nil {
		log.Printf(funcName+"UpdateTaskDetails Failed err: %s", err.Error())
		return err
	}

	return nil
}

func (s *TaskApplicaion) DeleteTask(ctx context.Context, req *request.DeleteTaskRequest) errors.IError {
	const funcName = "TaskApplicaion.DeleteTask "
	err := s.taskService.DeleteTaskDetails(ctx, req)
	if err != nil {
		log.Printf(funcName+"DeleteTaskDetails Failed err: %s", err.Error())
		return err
	}

	return nil
}
