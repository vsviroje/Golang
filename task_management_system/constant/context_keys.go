package constant

// CtxRmqConnectionName keys to circumvent go lint error
// `should not use basic type string as key in context.WithValue`
type (
	CtxRmqConnectionName struct{}
)
