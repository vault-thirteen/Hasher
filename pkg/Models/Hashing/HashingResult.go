package h

import (
	"bytes"
	"encoding/hex"
	"strconv"
	"strings"

	c "github.com/vault-thirteen/Hasher/pkg/Models/common"
)

type HashingResult struct {
	ResultType HashingResultType

	Binary  []byte
	Integer int
	Boolean bool
}

func NewHashingResult(data any, resultType HashingResultType) (hr *HashingResult, err error) {
	hr = &HashingResult{
		ResultType: resultType,
	}

	var ok bool
	switch resultType {
	case HashingResultType_Binary:
		hr.Binary, ok = data.([]byte)

	case HashingResultType_Integer:
		hr.Integer, ok = data.(int)

	case HashingResultType_Boolean:
		hr.Boolean, ok = data.(bool)

	default:
		return nil, c.Error(ErrUnknownHashingResultType)
	}
	if !ok {
		return nil, c.Error(ErrTypeAssertion)
	}

	return hr, nil
}

func (hr *HashingResult) Compare(value any) (isEqual bool, err error) {
	switch hr.ResultType {
	case HashingResultType_Binary:
		x, ok := value.([]byte)
		if !ok {
			return false, c.Error(ErrTypeAssertion)
		}

		isEqual = bytes.Equal(hr.Binary, x)
		return isEqual, nil

	case HashingResultType_Integer:
		x, ok := value.(int)
		if !ok {
			return false, c.Error(ErrTypeAssertion)
		}

		isEqual = (hr.Integer == x)
		return isEqual, nil

	case HashingResultType_Boolean:
		x, ok := value.(bool)
		if !ok {
			return false, c.Error(ErrTypeAssertion)
		}

		isEqual = (hr.Boolean == x)
		return isEqual, nil

	default:
		return false, c.Error(ErrUnknownHashingResultType)
	}
}

func (hr *HashingResult) ToString() (s string) {
	switch hr.ResultType {
	case HashingResultType_Binary:
		return strings.ToUpper(hex.EncodeToString(hr.Binary))

	case HashingResultType_Integer:
		return strconv.FormatInt(int64(hr.Integer), 10)

	case HashingResultType_Boolean:
		return c.FormatBooleanAsNumber(hr.Boolean)

	default:
		panic(ErrUnknownHashingResultType)
	}
}
