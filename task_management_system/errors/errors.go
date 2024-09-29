package errors

var (
	ErrParamInvalid = New(InvalidParamID, InvalidParamCode, ParamInvalidMsg)
	ErrHttpFail     = New(HttpFailID, HttpFailCode, HttpFailMsg)
	ErrDecode       = New(FailedToDecodeRequestBodyID, FailedToDecodeRequestBodyCode, FailedToDecodeRequestBodyMsg)
	ErrMarshal      = New(UnableToMarshalID, UnableToMarshalCode, UnableToMarshalMsg)
	ErrUnMarshal    = New(UnableToUnmarshalID, UnableToUnmarshalCode, UnableToUnMarshalMsg)
	ErrRmqRetry     = New(RmqRetryFailedID, RmqRetryReqFailedCode, RmqRetryFailedMsg)
)

const (
	msgSeparator = " "
)

func NewErrParamInvalid(msg ...string) IError {
	return New(InvalidParamID, InvalidParamCode, ParamInvalidMsg, msg...)
}

func NewErrRateLimitBreached(msg ...string) IError {
	return New(RateLimitBreachedID, RateLimitBreachedCode, RateLimitBreachedFailMsg, msg...)
}

func NewRequstValidationFailed(msg ...string) IError {
	return New(RequestValidationFailedID, RequestValidationFailedCode, RequestValidationFailMsg, msg...)
}

//MYSQL Error constructors
func NewErrMySQLReadFail(msg ...string) IError {
	return New(MySqlReadFailID, MysqlReadFailCode, MysqlReadFailMsg, msg...)
}

func NewErrMySQLWriteFail(msg ...string) IError {
	return New(MySqlWriteFailID, MysqlWriteFailCode, MysqlWriteFailMsg, msg...)
}

func NewErrDecode(msg ...string) IError {
	return New(FailedToDecodeRequestBodyID, FailedToDecodeRequestBodyCode, FailedToDecodeRequestBodyMsg, msg...)
}

func NewStripeRespErr(msg ...string) IError {
	return New(StripeHttpFailedID, StripeHttpFailCode, HttpFailMsg, msg...)
}

func NewErrHttpFail(msg ...string) IError {
	return New(HttpFailID, HttpFailCode, HttpFailMsg, msg...)
}

func NewFailedToDecodeStripeResponseBody(msg ...string) IError {
	return New(FailedToDecodeStripeResponseBodyID, FailedToDecodeStripeResponseBodyCode, FailedToDecodeResponseBodyMsg, msg...)
}

func NewResourceNotFound(msg ...string) IError {
	return New(ResourceNotFoundGenericID, ResourceNotFoundCode, ResourceNotFoundMsg, msg...)
}

func NewStripeResponseErr(code, msg string) IError {
	return New(StripeHttpFailedID, code, msg, "Response by stripe:"+code+"|"+msg)
}

func NewStripeServiceNotAvailableErr(msg ...string) IError {
	return New(StripeHttpFailedID, StripeServiceNotAvailableCode, StripeServiceNotFoundMsg, msg...)
}

func NewNoCallbackReceivedErr(msg ...string) IError {
	return New(NoCallBackReceivedFailedID, NoCallBackReceivedFailedCode, NoCallbackReceivedMsg, msg...)
}
