package constant

// logger keys
const (
	Identifier            = "identifier"
	AsyncRequest          = "async"
	ServiceInitialisation = "init"
	RequestID             = "request_id"
	Args                  = "args"
	HttpStatusCode        = "status_code"
	HttpUrl               = "request_url"
	UserAgent             = "user_agent"
	RequestMethod         = "method"
	Application           = "application"
	LogType               = "label"
	Response              = "response"
	Headers               = "headers"
	Query                 = "query"
)

// logger labels
const (
	Internal = "INTERNAL_SYSTEMS"
	External = "EXTERNAL_REQUESTS"
	Events   = "EVENT_PIPELINE"
)

const (
	MongoTransactionalDatabaseId = "mongo"
	MysqlTransactionalDatabaseId = "mysql"
)
