package filter

//Param struct
type Param struct {
	Field    string
	Operator string
	Value    interface{}
}

//const Operation
const (
	StartsWith  = "STARTS_WITH"
	Includes    = "INCLUDES"
	GreaterThan = "GREATER_THAN"
	LessThan    = "LESS_THAN"
	Equals      = "EQUALS"
)

//OrderBy struct
//TODO: Need to remove this
type OrderBy struct {
	Field     string
	Direction int64
}

//Options struct
type Options struct {
	Operation string
	Value     interface{}
}

//const Order
const (
	ASCENDING  = 1
	DESCENDING = -1
)

//order constant
const (
	SearchOrderDesc = "desc"
	SearchOrderAsc  = "asc"
)
