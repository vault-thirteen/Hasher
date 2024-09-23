package h

type HashingResultType byte

const (
	HashingResultType_Binary  = HashingResultType(1)
	HashingResultType_Integer = HashingResultType(2)
	HashingResultType_Boolean = HashingResultType(3)
)

const (
	ErrUnknownHashingResultType = "unknown hashing result type"
	ErrTypeAssertion            = "type assertion has failed"
)
