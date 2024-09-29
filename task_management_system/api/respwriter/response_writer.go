package respwriter

import (
	"encoding/json"
	"net/http"
	"strconv"

	"task_management_system/errors"
	"task_management_system/util"
)

// ErrorResponse struct
type ErrorResponse struct {
	Success bool         `json:"success"`
	Failure ErrorDetails `json:"failure"`
}

// ErrorDetails struct
type ErrorDetails struct {
	ErrorID   int           `json:"errorId" example:"400001"`
	ErrorCode string        `json:"errorCode" example:"ERROR_VALIDATION_FAILURE"`
	Debug     *DebugDetails `json:"debug,omitempty"`
}

type DebugDetails struct {
	Description  string `json:"description,omitempty" example:"Password should be 8 to 13 character long"`
	Message      string `json:"message,omitempty" example:"Invalid password"`
	ResponseTime string `json:"responseTime,omitempty" example:"2021-07-10T13:03:17"`
}

// SuccessResponse struct
type SuccessResponse struct {
	Success bool        `json:"success"`
	Result  interface{} `json:"result"`
}

// IHttpJSONResponse defines an interface
//
//go:generate go run github.com/golang/mock/mockgen -destination=./response_writer_mock.go -package=respwriter . IHttpJSONResponse
type IHttpJSONResponse interface {
	JSON(r *http.Request, w http.ResponseWriter, result interface{}, v interface{}, statusCode ...int) error
}

type httpJSONResponseWriter struct {
}

var serviceInstance IHttpJSONResponse

// NewHTTPResponseService initialization
func NewHTTPResponseService() IHttpJSONResponse {
	serviceInstance = &httpJSONResponseWriter{}
	return serviceInstance
}

// GetHTTPResponseService is self-explanatory
func GetHTTPResponseService() IHttpJSONResponse {
	return serviceInstance
}

func (c *httpJSONResponseWriter) JSON(r *http.Request, w http.ResponseWriter, result interface{}, v interface{}, statusCodeOpt ...int) error {
	statusCode := http.StatusOK

	if len(statusCodeOpt) != 0 {
		statusCode = statusCodeOpt[0]
	}

	customError, ok := v.(errors.IError)
	if ok {
		statusCode := getHTTPStatus(customError.ErrorID())
		resp := &ErrorResponse{
			Success: false,
			Failure: ErrorDetails{
				ErrorID:   customError.ErrorID(),
				ErrorCode: customError.ErrorCode(),
				Debug: &DebugDetails{
					Message:      customError.ErrorDescription(),
					Description:  customError.ErrorMsg(),
					ResponseTime: util.CurrentTime().String(),
				},
			},
		}
		w.WriteHeader(statusCode)
		return c.marshalToJSON(r, w, resp)

	}

	err, ok := v.(error)
	if ok {
		return c.writeErrorWithCode(r, w, result, err, statusCode)
	}

	w.WriteHeader(statusCode)
	return c.writeJSON(r, w, v)
}

func (c *httpJSONResponseWriter) writeJSON(r *http.Request, w http.ResponseWriter, v interface{}) error {
	// var result interface{}
	// resp := SuccessResponse{
	// 	Success: true,
	// 	Result:  v,
	// }
	// result = resp

	return c.marshalToJSON(r, w, v)
}

// Write interface value as json
func (c *httpJSONResponseWriter) marshalToJSON(r *http.Request, w http.ResponseWriter, v interface{}) error {

	data, err := json.Marshal(v)
	if err != nil {
		return err
	}

	_, err = w.Write(data)
	return err
}

func (c *httpJSONResponseWriter) writeErrorWithCode(r *http.Request, w http.ResponseWriter, result interface{}, errorDetails error, errorID int, messages ...string) error {
	var response interface{}

	title, desc := getTitleAndDescription(messages)
	statusCode, response := c.getErrorResponse(errorID, title, desc, errorDetails.Error(), result)

	w.WriteHeader(statusCode)

	return c.marshalToJSON(r, w, response)
}

func (c *httpJSONResponseWriter) getErrorResponse(errorID int, errorCode string, message string, description string, result interface{}) (int, interface{}) {
	statusCode := getHTTPStatus(errorID)
	resp := &ErrorResponse{
		Success: false,
		Failure: ErrorDetails{
			ErrorID:   errorID,
			ErrorCode: errorCode,
			Debug: &DebugDetails{
				Message:      message,
				Description:  description,
				ResponseTime: util.CurrentTime().String(),
			},
		},
	}

	return statusCode, resp
}

func getHTTPStatus(code int) (status int) {
	firstThreeDigits := "500"
	codeArr := strconv.Itoa(code)
	if len(codeArr) >= 3 {
		firstThreeDigits = codeArr[:3]
	}
	status, _ = strconv.Atoi(firstThreeDigits)
	return
}

func getTitleAndDescription(messages []string) (string, string) {
	var ttl, desc string
	if len(messages) > 0 {
		ttl = messages[0]
	}
	if len(messages) > 1 {
		desc = messages[1]
	}
	return ttl, desc
}
