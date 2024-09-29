package errors

/*
    Prefixes:
	1. Service Internal	     = 11
	2. Queue 	  		     = 12 (rmq)
	3. Databases  		     = 30
	4. Alerting platforms    = 40 (opsgenie/slack/newrelic)

*/

// Unique errorId mapping to combination of service id + error source + error reason. Convention to be followed shall be https://libertywireless.atlassian.net/wiki/spaces/EQL/pages/2517434629/Development+Guidelines+HTTP#HTTP-Status-Code
const (
	/****************************** Status Code 400 ******************************/
	FailedToDecodeRequestBodyID int = 4000010001
	MissingArgumentGenericID    int = 4000010002
	InvalidParamID              int = 4000010003
	RequestValidationFailedID   int = 4000010004
	StripeHttpFailedID          int = 4000010005
	UnkownStripeCallbackTypeID  int = 4000010006
	StripeGenericFailID         int = 4000010007 //stripe generic error ie. card_decline
	RmqRetryFailedID            int = 4000010008
	NoCallBackReceivedFailedID  int = 4000010009
	ResourceNotFoundGenericID   int = 4000010010

	/****************************** Status Code 404 ******************************/

	/****************************** Status Code 409 ******************************/

	/****************************** Status Code 422 ******************************/
	MissingHeaderID        int = 4220010001
	JSONExtractionFailedID int = 4220010002

	/****************************** Status Code 423 ******************************/

	/****************************** Status Code 424 ******************************/

	/****************************** Status Code 429 ******************************/
	RateLimitBreachedID int = 4290010001

	/****************************** Status Code 500 ******************************/
	PublisherNotFoundID                int = 5000110001
	HttpFailID                         int = 5000110002
	FailedToDecodeStripeResponseBodyID int = 5000110003
	UUIDConversionFailedID             int = 5000110004
	UnableToUnmarshalID                int = 5000110005
	UnableToMarshalID                  int = 5000110006
	//Mysql errors
	MySqlWriteFailID            int = 5000300001
	MySqlReadFailID             int = 5000300002
	PPStmtCreationFailedErrorID int = 5000300003
	MySqlFailedFetchingDataID   int = 5000300004
	MysqlRecordNotFoundErrorID  int = 5000300005
	MysqlFailedToUpdateErrorID  int = 5000300006

	// Specification errors
	FailedToGetCoreSpecificationsID int = 5000300007
	FailedCoreCallbackRequestID     int = 5000300008
)
