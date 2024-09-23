package h

import (
	"encoding/hex"
	"os"
	"strconv"
	"strings"

	ht "github.com/vault-thirteen/Hasher/pkg/Models/HashType"
	c "github.com/vault-thirteen/Hasher/pkg/Models/common"
	"github.com/vault-thirteen/auxie/number"
)

const (
	ErrDataIsDamaged = "data is damaged"
)

type Hashing struct {
	// typ is type. The word 'type' can not be used in Go language.
	typ *ht.HashType
}

func NewHashing(hti ht.HashTypeId) (h *Hashing, err error) {
	h = &Hashing{}

	h.typ, err = ht.NewById(hti)
	if err != nil {
		return nil, err
	}

	return h, nil
}

func (h *Hashing) GetType() (ht ht.HashType) {
	if h.typ == nil {
		return
	}

	return *(h.typ)
}

func (h *Hashing) Calculate(file string) (result *HashingResult, err error) {
	var data any

	switch h.typ.ID() {
	case ht.Id_FileExistence:
		data, err = checkFileExistence(file)
		if err != nil {
			return nil, err
		}

		result, err = NewHashingResult(data, HashingResultType_Boolean)
		if err != nil {
			return nil, err
		}

	case ht.Id_FileSize:
		data, err = getFileSize(file)
		if err != nil {
			return nil, err
		}

		result, err = NewHashingResult(data, HashingResultType_Integer)
		if err != nil {
			return nil, err
		}

	case ht.Id_CRC32:
		data, err = calculateBinaryFileHash(file, ht.Id_CRC32)
		if err != nil {
			return nil, err
		}

		result, err = NewHashingResult(data, HashingResultType_Binary)
		if err != nil {
			return nil, err
		}

	case ht.Id_MD5:
		data, err = calculateBinaryFileHash(file, ht.Id_MD5)
		if err != nil {
			return nil, err
		}

		result, err = NewHashingResult(data, HashingResultType_Binary)
		if err != nil {
			return nil, err
		}

	case ht.Id_SHA256:
		data, err = calculateBinaryFileHash(file, ht.Id_SHA256)
		if err != nil {
			return nil, err
		}

		result, err = NewHashingResult(data, HashingResultType_Binary)
		if err != nil {
			return nil, err
		}

	default:
		return nil, c.Error(ht.ErrUnknownHashType)
	}

	return result, nil
}

func (h *Hashing) Verify(file string, value any) (isEqual bool, err error) {
	var result *HashingResult
	result, err = h.Calculate(file)
	if err != nil {
		return false, err
	}

	return result.Compare(value)
}

func (h *Hashing) ParseHash(hashText string) (hashValue any, err error) {
	switch h.typ.ID() {
	case ht.Id_FileExistence:
		var x bool
		x, err = strconv.ParseBool(hashText)
		if err != nil {
			return "", err
		}

		return x, nil

	case ht.Id_FileSize:
		var x int
		x, err = number.ParseInt(hashText)
		if err != nil {
			return "", err
		}

		return x, nil

	case ht.Id_CRC32:
	case ht.Id_MD5:
	case ht.Id_SHA256:

	default:
		return nil, c.Error(ht.ErrUnknownHashType)
	}

	// Binary hash sum.
	var x []byte
	x, err = hex.DecodeString(hashText)
	if err != nil {
		return "", err
	}

	return x, nil
}

// WalkerFn is a method implementing the signature of a 'walk function':
// type WalkFunc func(path string, info fs.FileInfo, err error) error
func (h *Hashing) WalkerFn(path string, fi os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if fi.IsDir() {
		return nil
	}

	var result *HashingResult
	result, err = h.Calculate(path)
	if err != nil {
		return err
	}

	c.PrintHashLine(result.ToString(), path)
	return nil
}

func (h *Hashing) ParseFileLine(line []byte) (hashText string, filePath string, err error) {
	// Hash line starts with a hash text separated with a single space symbol:
	// XXXX <FilePath>
	if h.typ.IsBinary() {
		hashTextLen := int(h.typ.SumSize()) * 2
		if len(line) < hashTextLen+2 {
			return "", "", c.Error(ErrDataIsDamaged)
		}

		hashText = strings.ToUpper(string(line[:hashTextLen]))
		filePath = strings.TrimSpace(string(line[(hashTextLen + 1):]))
		return hashText, filePath, nil
	}

	var p1, p2 string
	p1, p2, err = splitLine(line)
	if err != nil {
		return "", "", err
	}

	filePath = strings.TrimSpace(p2)

	switch h.typ.ID() {
	case ht.Id_FileExistence:
		hashText, err = parseAndFormatBoolean(p1)
	case ht.Id_FileSize:
		hashText, err = parseAndFormatInteger(p1)
	default:
		err = c.Error(ht.ErrUnknownHashType)
	}
	if err != nil {
		return "", "", err
	}

	return hashText, filePath, nil
}

func splitLine(line []byte) (p1, p2 string, err error) {
	parts := strings.Split(string(line), " ")
	if len(parts) != 2 {
		return "", "", c.Error(ErrDataIsDamaged)
	}

	return parts[0], parts[1], nil
}

func parseAndFormatBoolean(in string) (out string, err error) {
	var x bool
	x, err = strconv.ParseBool(in)
	if err != nil {
		return "", err
	}

	return c.FormatBooleanAsNumber(x), nil
}

func parseAndFormatInteger(in string) (out string, err error) {
	_, err = number.ParseInt(in)
	if err != nil {
		return "", err
	}

	return in, nil
}
