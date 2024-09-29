package errors

// Error glossary mapping to errorCode https://libertywireless.atlassian.net/wiki/spaces/I2/pages/2543288367/System+Error+Glossary
const (
	FailedToDecodeRequestBodyCode string = "ERROR_DECODE_REQUEST_BODY"
	MissingHeaderCode             string = "ERROR_HEADER_MISSING"
	PublisherNotFoundCode         string = "ERROR_PUBLISHER_NOT_FOUND"
	JSONExtractionFailedCode      string = "ERROR_JSON_EXTRACTION"
	//Mysql
	MysqlReadFailCode             string = "ERROR_MYSQL_READ_FAILURE"
	MysqlWriteFailCode            string = "ERROR_MYSQL_WRITE_FAILURE"
	PPStmtCreationFailedErrorCode string = "ERROR_PREPARING_MYSQL_STATEMENT"
	MySqlFailedFetchingDataCode   string = "ERROR_FETCHING_MYSQL_DATA"
	MysqlFailedToUpdateErrorCode  string = "ERROR_UPDATING_MYSQL_DATA"
	ResourceNotFoundCode          string = "ERROR_RESOURCE_NOT_FOUND"

	//MysqlRecordNotFoundErrorCode string="ERROR_"
	MissingArgumentGenericCode string = "ERROR_ARGUMENT_MISSING"
	InvalidParamCode           string = "ERROR_PARAM_INVALID"
	DataNotFoundCode           string = "ERROR_DATA_NOT_FOUND"

	RateLimitBreachedCode                string = "ERROR_RATE_LIMIT"
	RequestValidationFailedCode          string = "ERROR_REQUEST_VALIDATION"
	HttpFailCode                         string = "ERROR_HTTP_FAILURE"
	StripeHttpFailCode                   string = "ERROR_STRIPE_HTTP_FAILURE"
	FailedToDecodeStripeResponseBodyCode string = "ERROR_DECODE_STRIPE_RESPONSE"

	UnkownStripeCallbackTypeCode string = "ERROR_UNKNOWN_STRIPE_CALLBACK_TYPE"

	// Specification errors
	FailedToGetCoreSpecificationsCode string = "ERROR_CORE_SPECIFICATION_GET_FAILED"
	FailedCoreCallbackRequestCode     string = "ERROR_CORE_CALLBACK_REQUEST_FAILED"

	UUIDConversionFailedCode string = "ERROR_UUID_CONVERSION_TO_STRING_FAILED"

	//marshal,unmarshal errors
	UnableToMarshalCode   string = "ERROR_PARAM_INVALID"
	UnableToUnmarshalCode string = "ERROR_PARAM_INVALID"

	//Specification Errors
	DuplicateSpecificationFailedCode string = "ERROR_SPECIFICATION_ALREADY_EXISTS"

	//Stripe setup or payment intent failure
	StripeGenericFailCode         string = "ERROR_STRIPE_FAILURE"
	StripeServiceNotAvailableCode string = "ERROR_STRIPE_SERVICE_NOT_AVAILABLE"

	RmqRetryReqFailedCode        string = "ERROR_RMQ_RETRY_FAILED"
	NoCallBackReceivedFailedCode string = "ERROR_NO_CALL_BACK_RECEIVED"
)
