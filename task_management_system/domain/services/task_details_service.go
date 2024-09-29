package services

import (
	"context"
	"task_management_system/config"
	"task_management_system/domain/task_details"
	"task_management_system/dto/request"
	"task_management_system/dto/response"
	"task_management_system/errors"
)

//go:generate mockery --name ITaskDetailsService --inpackage --filename=task_details_service_mock.go
type ITaskDetailsService interface {
	GetTaskDetails(ctx context.Context, req *request.GetTaskRequest) (resp interface{}, err errors.IError)
	AddTaskDetails(ctx context.Context, req *request.AddTaskRequest) (id *string, err errors.IError)
	UpdateTaskDetails(ctx context.Context, req *request.UpdateTaskRequest) (err errors.IError)
	DeleteTaskDetails(ctx context.Context, req *request.DeleteTaskRequest) (err errors.IError)
}

type TaskDetailsService struct {
	config          *config.GeneralConfig
	taskDetailsRepo task_details.ITaskDetailsRepo
}

func NewTaskDetailsService(
	cfg *config.GeneralConfig,
	tdr task_details.ITaskDetailsRepo,
) ITaskDetailsService {
	return &TaskDetailsService{
		config:          cfg,
		taskDetailsRepo: tdr,
	}
}

func (s *TaskDetailsService) GetTaskDetails(ctx context.Context, req *request.GetTaskRequest) (interface{}, errors.IError) {
	var data interface{}
	var list []task_details.TaskDetails
	var err errors.IError
	var listResp []response.GetTaskDetails

	if req.TaskId != "" {
		var rec *task_details.TaskDetails
		rec, err = s.taskDetailsRepo.GetTaskDetailsById(ctx, req.TaskId)
		if err != nil {
			data = s.parseTaskDetailsToGetTaskResp(rec)
		}
	} else if req.Status != "" && req.UserId != "" {
		list, err = s.taskDetailsRepo.GetTaskDetailsByUserIdAndStatus(ctx, req.UserId, req.Status)
		if err != nil {
			for _, task := range list {
				listResp = append(listResp, s.parseTaskDetailsToGetTaskResp(&task))
			}
			data = listResp
		}
	} else if req.Status == "" && req.UserId != "" {
		list, err = s.taskDetailsRepo.GetTaskDetailsByUserId(ctx, req.UserId)
		if err != nil {
			for _, task := range list {
				listResp = append(listResp, s.parseTaskDetailsToGetTaskResp(&task))
			}
			data = listResp
		}
	} else {
		err = errors.New(errors.MissingArgumentGenericID, errors.MissingArgumentGenericCode, "No valid data")
	}
	return data, err
}

func (s *TaskDetailsService) AddTaskDetails(ctx context.Context, req *request.AddTaskRequest) (id *string, err errors.IError) {
	return s.taskDetailsRepo.AddTaskDetails(ctx, s.parseAddTaskReqToTaskDetails(req))
}

func (s *TaskDetailsService) UpdateTaskDetails(ctx context.Context, req *request.UpdateTaskRequest) (err errors.IError) {
	return s.taskDetailsRepo.UpdateTaskDetails(ctx, s.parseUpdateTaskReqToTaskDetails(req))
}

func (s *TaskDetailsService) DeleteTaskDetails(ctx context.Context, req *request.DeleteTaskRequest) (err errors.IError) {
	return s.taskDetailsRepo.DeleteTaskDetailsById(ctx, req.TaskId)
}

func (s *TaskDetailsService) parseTaskDetailsToGetTaskResp(rec *task_details.TaskDetails) response.GetTaskDetails {
	return response.GetTaskDetails{
		TaskId:      *rec.Id,
		UserId:      rec.UserId,
		Status:      *rec.Status,
		Title:       rec.Title,
		Description: rec.Description,
		DueDate:     rec.DueDate,
	}
}

func (s *TaskDetailsService) parseAddTaskReqToTaskDetails(req *request.AddTaskRequest) *task_details.TaskDetails {
	record := task_details.TaskDetails{}
	record.Title = &req.Title
	record.Description = &req.Description
	record.DueDate = &req.DueDate
	if req.UserId != nil {
		record.UserId = req.UserId
	}
	return &record
}

func (s *TaskDetailsService) parseUpdateTaskReqToTaskDetails(req *request.UpdateTaskRequest) *task_details.TaskDetails {
	record := task_details.TaskDetails{}
	record.Title = req.Title
	record.Description = req.Description
	record.DueDate = req.DueDate
	record.UserId = req.UserId
	record.Id = &req.TaskId
	return &record
}
