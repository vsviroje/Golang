package constant

const (
	StripePaymentGateway = "stripe"

	//values for Type or TransactionType
	TransactionTypeUpfront    = "UPFRONT"
	TransactionTypeRecurrent  = "RECURRENT"
	TransactionTypeRefund     = "REFUND"
	TransactionTypeUpdateCard = "UPDATE_CARD"

	// Payment Transaction State
	PaymentCreated = "CREATED"

	//Payment Transaction Status
	Success  = "SUCCESS"
	Failure  = "FAILURE"
	Rejected = "REJECTED"
	Pending  = "PENDING"

	// Transaction Type
	TEST_CARD_TRANSACTION = "TEST"
	LIVE_TRANSACTION      = "LIVE"

	STRIPE_SETUP_INTENT     = "confirmCardSetup"
	STRIPE_ONSESSION_INTENT = "confirmCardPayment"

	//intent type
	SetupIntent   = "setup_intent"
	PaymentIntent = "payment_intent"

	//RMQ
	IsRetryReq string = "isRetryReq"
)
