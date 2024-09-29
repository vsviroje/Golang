package controller

import (
	"net/http"
	"task_management_system/api/respwriter"
	"task_management_system/appcontext"
)

// HealthController struct
type HealthController struct {
	*BaseController
}

// NewHealthController function
func NewHealthController(baseController *BaseController) *HealthController {
	return &HealthController{BaseController: baseController}
}

// HealthCheck godoc
// @Summary Check the health of the service.
// @Description This API will return the server status.
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} respwriter.SuccessResponse
// @Failure 500 {object} respwriter.ErrorResponse
// @router /health [get]
func (c *HealthController) HealthCheck(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestID := appcontext.GetRequestContext(ctx).RequestID()
	c.HandleResponse(HandleResponseDto{requestID, nil, w, r, respwriter.SuccessResponse{Success: true}, nil}, http.StatusOK)
}
