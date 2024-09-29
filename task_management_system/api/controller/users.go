package controller

import (
	"log"
	"net/http"

	"task_management_system/appcontext"
	"task_management_system/application"
	"task_management_system/dto/request"
	"task_management_system/dto/response"
)

type UsersController struct {
	*BaseController
	userAppication application.IUsersApplication
}

func NewUsersController(baseController *BaseController,
	ua application.IUsersApplication) *UsersController {
	return &UsersController{
		baseController,
		ua,
	}
}

// Adduser godoc
// @Summary add user
// @Description This API will process to add user
// @Tags Users
// @Accept json
// @Produce json
// @Param req body request.AddUserRequest true "Payload"
// @Success 200 {object} response.AddUsersResponse
// @Failure 400 {object} respwriter.ErrorResponse
// @Failure 500 {object} respwriter.ErrorResponse
// @router /v1/user [POST]
func (ctrl *UsersController) AddUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestID := appcontext.GetRequestContext(ctx).RequestID()

	var req request.AddUserRequest
	err := request.DecodeAndValidate(ctx, r, &req)

	if err != nil {
		log.Printf("[AddUser|DecodeAndValidate|failed] err:%s", err.Error())
		ctrl.HandleResponse(HandleResponseDto{requestID, r.URL.Query(), w, r, nil, err}, http.StatusBadRequest)
		return
	}

	resp, err := ctrl.userAppication.AddUser(ctx, &req)
	if err != nil {
		log.Printf("[UsersController.AddUser|userAppication.AddSpecification|failed] err:%s", err.Error())
		ctrl.HandleResponse(HandleResponseDto{requestID, r.URL.Query(), w, r, nil, err}, http.StatusBadRequest)
		return
	}

	resp.RequestId = requestID

	ctrl.HandleResponse(HandleResponseDto{requestID, r.URL.Query(), w, r, resp, nil}, http.StatusOK)

}

// AssignUserTask godoc
// @Summary Assign task to user
// @Description This API will process to assign task to user
// @Tags Users
// @Accept json
// @Produce json
// @Param req body request.AssignUserTaskRequest true "Payload"
// @Success 200 {object} response.AssingUsersTaskResponse
// @Failure 400 {object} respwriter.ErrorResponse
// @Failure 500 {object} respwriter.ErrorResponse
// @router /v1/user [POST]
func (ctrl *UsersController) AssignUserTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestID := appcontext.GetRequestContext(ctx).RequestID()

	var req request.AssignUserTaskRequest
	err := request.DecodeAndValidate(ctx, r, &req)
	if err != nil {
		log.Printf("[AssignUserTask|DecodeAndValidate|failed] err:%s", err.Error())
		ctrl.HandleResponse(HandleResponseDto{requestID, r.URL.Query(), w, r, nil, err}, http.StatusBadRequest)
		return
	}

	err = ctrl.userAppication.AssignUserTask(ctx, &req)
	if err != nil {
		log.Printf("[UsersController.AssignUserTask|userAppication.AssignUserTask|failed] err:%s", err.Error())
		ctrl.HandleResponse(HandleResponseDto{requestID, r.URL.Query(), w, r, nil, err}, http.StatusBadRequest)
		return
	}

	resp := response.AssingUsersTaskResponse{
		RequestId: requestID,
	}

	ctrl.HandleResponse(HandleResponseDto{requestID, r.URL.Query(), w, r, resp, nil}, http.StatusOK)

}
