package h

type HashingResultType byte

const (
	HashingResultType_Binary  = HashingResultType(1)
	HashingResultType_Integer = HashingResultType(2)
	HashingResultType_Boolean = HashingResultType(3)
)

const (
	Err_UnknownHashingResultType = "unknown hashing result type"
	Err_TypeAssertion            = "type assertion has failed"
)
