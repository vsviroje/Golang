package controller

import (
	"log"
	"net/http"

	"task_management_system/api/respwriter"
	"task_management_system/errors"
)

// BaseController struct
type BaseController struct {
	respwriter.IHttpJSONResponse
}

// NewBaseController for initialisation
func NewBaseController(httpResponseService respwriter.IHttpJSONResponse) *BaseController {
	return &BaseController{
		httpResponseService,
	}
}

type HandleResponseDto struct {
	RequestID    string
	RequestQuery interface{}
	W            http.ResponseWriter
	R            *http.Request
	Result       interface{}
	Err          error
}

// HandleResponse function
func (b *BaseController) HandleResponse(dto HandleResponseDto, statusCodeOpt ...int) {
	if dto.Err != nil {
		errorID := http.StatusInternalServerError
		customError, ok := dto.Err.(errors.IError)
		if ok {
			errorID = customError.ErrorID()
		}
		log.Printf("error HandleResponse err:%v | resp:%+v", dto.Err, dto)

		_ = b.JSON(dto.R, dto.W, dto.Result, dto.Err, errorID)
		return
	}
	log.Printf("success HandleResponse resp:%+v", dto)

	customStatusCode := http.StatusOK
	if len(statusCodeOpt) > 0 {
		customStatusCode = statusCodeOpt[0]
	}
	_ = b.JSON(dto.R, dto.W, nil, dto.Result, customStatusCode)

}
