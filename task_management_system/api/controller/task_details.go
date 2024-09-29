package controller

import (
	"log"
	"net/http"

	"task_management_system/appcontext"
	"task_management_system/application"
	"task_management_system/dto/request"
	"task_management_system/dto/response"
)

type TaskDetailsController struct {
	*BaseController
	taskApplication application.ITaskApplication
}

func NewTaskDetailsController(baseController *BaseController,
	ta application.ITaskApplication) *TaskDetailsController {
	callbackControllerInstance := &TaskDetailsController{
		baseController,
		ta,
	}
	return callbackControllerInstance
}

// GetTask godoc
// @Summary Get task
// @Description This API will process to add task
// @Tags Task
// @Produce json
// @Param taskId query string true "taskId eg, 1"
// @Param userId query string true "userId eg, 2"
// @Param status query string true "status eg, pending"
// @Success 200 {object} response.GetTaskResponse
// @Failure 400 {object} respwriter.ErrorResponse
// @Failure 500 {object} respwriter.ErrorResponse
// @router /v1/task [GET]
func (ctrl *TaskDetailsController) GetTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestID := appcontext.GetRequestContext(ctx).RequestID()
	var req request.GetTaskRequest

	req.TaskId = r.URL.Query().Get("taskId")
	req.UserId = r.URL.Query().Get("userId")
	req.Status = r.URL.Query().Get("status")

	err := request.DecodeAndValidate(ctx, r, &req)

	if err != nil {
		log.Printf("[TaskDetailsController.GetTask|DecodeAndValidate|failed] err:%s", err.Error())
		ctrl.HandleResponse(HandleResponseDto{requestID, r.URL.Query(), w, r, nil, err}, http.StatusBadRequest)
		return
	}

	resp, err := ctrl.taskApplication.GetTask(ctx, &req)
	if err != nil {
		log.Printf("[TaskDetailsController.GetTask|taskApplication.GetTask|failed] err:%s", err.Error())
		ctrl.HandleResponse(HandleResponseDto{requestID, r.URL.Query(), w, r, nil, err}, http.StatusBadRequest)
		return
	}

	resp.RequestId = requestID

	ctrl.HandleResponse(HandleResponseDto{requestID, r.URL.Query(), w, r, resp, nil}, http.StatusOK)

}

// AddTask godoc
// @Summary add task
// @Description This API will process to add task
// @Tags Users
// @Accept json
// @Produce json
// @Param req body request.AddTaskRequest true "Payload"
// @Success 200 {object} response.AddTaskResponse
// @Failure 400 {object} respwriter.ErrorResponse
// @Failure 500 {object} respwriter.ErrorResponse
// @router /v1/task [POST]
func (ctrl *TaskDetailsController) AddTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestID := appcontext.GetRequestContext(ctx).RequestID()

	var req request.AddTaskRequest
	err := request.DecodeAndValidate(ctx, r, &req)

	if err != nil {
		log.Printf("[AddTask|DecodeAndValidate|failed] err:%s", err.Error())
		ctrl.HandleResponse(HandleResponseDto{requestID, r.URL.Query(), w, r, nil, err}, http.StatusBadRequest)
		return
	}

	resp, err := ctrl.taskApplication.AddTask(ctx, &req)
	if err != nil {
		log.Printf("[TaskDetailsController.AddTask|taskApplication.AddTask|failed] err:%s", err.Error())
		ctrl.HandleResponse(HandleResponseDto{requestID, r.URL.Query(), w, r, nil, err}, http.StatusBadRequest)
		return
	}

	resp.RequestId = requestID

	ctrl.HandleResponse(HandleResponseDto{requestID, r.URL.Query(), w, r, resp, nil}, http.StatusOK)

}

// Update task godoc
// @Summary update task
// @Description This API will process to update task details
// @Tags Users
// @Accept json
// @Produce json
// @Param req body request.UpdateTaskRequest true "Payload"
// @Success 200 {object} response.UpdateTaskResponse
// @Failure 400 {object} respwriter.ErrorResponse
// @Failure 500 {object} respwriter.ErrorResponse
// @router /v1/task [PUT]
func (ctrl *TaskDetailsController) UpdateTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestID := appcontext.GetRequestContext(ctx).RequestID()

	var req request.UpdateTaskRequest
	err := request.DecodeAndValidate(ctx, r, &req)

	if err != nil {
		log.Printf("[UpdateTask|DecodeAndValidate|failed] err:%s", err.Error())
		ctrl.HandleResponse(HandleResponseDto{requestID, r.URL.Query(), w, r, nil, err}, http.StatusBadRequest)
		return
	}

	err = ctrl.taskApplication.UpdateTask(ctx, &req)
	if err != nil {
		log.Printf("[TaskDetailsController.UpdateTask|taskApplication.UpdateTask|failed] err:%s", err.Error())
		ctrl.HandleResponse(HandleResponseDto{requestID, r.URL.Query(), w, r, nil, err}, http.StatusBadRequest)
		return
	}

	resp := response.UpdateTaskResponse{
		RequestId: requestID,
	}

	ctrl.HandleResponse(HandleResponseDto{requestID, r.URL.Query(), w, r, resp, nil}, http.StatusOK)

}

// DeleteUser godoc
// @Summary delete user
// @Description This API will process to delete user
// @Tags Users
// @Produce json
// @Param taskId query string true "taskId eg, 1"
// @Success 200 {object} response.DeleteTaskResponse
// @Failure 400 {object} respwriter.ErrorResponse
// @Failure 500 {object} respwriter.ErrorResponse
// @router /v1/task [DELETE]
func (ctrl *TaskDetailsController) DeleteTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestID := appcontext.GetRequestContext(ctx).RequestID()

	var req request.DeleteTaskRequest
	req.TaskId = r.URL.Query().Get("taskId")
	err := request.DecodeAndValidate(ctx, r, &req)
	if err != nil {
		log.Printf("[DeleteTask|DecodeAndValidate|failed] err:%s", err.Error())
		ctrl.HandleResponse(HandleResponseDto{requestID, r.URL.Query(), w, r, nil, err}, http.StatusBadRequest)
		return
	}

	err = ctrl.taskApplication.DeleteTask(ctx, &req)
	if err != nil {
		log.Printf("[TaskDetailsController.DeleteTask|taskApplication.DeleteTask|failed] err:%s", err.Error())
		ctrl.HandleResponse(HandleResponseDto{requestID, r.URL.Query(), w, r, nil, err}, http.StatusBadRequest)
		return
	}

	resp := response.DeleteTaskResponse{
		RequestId: requestID,
	}

	ctrl.HandleResponse(HandleResponseDto{requestID, r.URL.Query(), w, r, resp, nil}, http.StatusOK)

}
