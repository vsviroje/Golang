package errors

// Descriptive error message which shall map to the human understable error description
const (
	//MySQL generic messages
	MysqlReadFailMsg  string = "Error, database read failed"
	MysqlWriteFailMsg string = "Error, database write failed"

	//validator
	ValidatorZeroValueMsg   string = "empty, null or zero value"
	ValidatorMissingArgsMsg string = "invalid due to fields: %s"

	ParamInvalidMsg string = "Error, parameter invalid"

	MissingPaymentMethodID        string = "Error: could not get the stringified paymentMethodId from event"
	MissingPaymentCustomerID      string = "Error: could not get the stringified stripeCustomerId from event"
	RateLimitBreachedFailMsg      string = "Error, rate limit breached"
	RequestValidationFailMsg      string = "Error, request validation failed"
	HttpFailMsg                   string = "Error, downstream http request failed"
	FailedToDecodeRequestBodyMsg  string = "Error, read request body failed"
	FailedToDecodeResponseBodyMsg string = "Error, read downstream response body failed"

	UnkownStripeCallbackTypeMsg string = "Error, Unknown type from stripe callback"

	UnableToMarshalMsg   = "parse error"
	UnableToUnMarshalMsg = "parse error"

	ResourceNotFoundMsg      string = "Error, resource not found"
	StripeServiceNotFoundMsg string = "Stripe Error, Service Unavailable"

	RmqRetryFailedMsg     string = "Rmq Error, Retry failed"
	NoCallbackReceivedMsg string = "No Callback Received"
)
