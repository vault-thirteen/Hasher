package ht

import (
	"strings"

	hs "github.com/vault-thirteen/Hasher/pkg/Models/HashSize"
	c "github.com/vault-thirteen/Hasher/pkg/Models/common"
)

const (
	ErrUnknownHashType = "unknown hash type"
)

type HashType struct {
	id   HashTypeId
	name HashTypeName

	isBinary bool
	sumSize  hs.HashSize
}

func New(hashTypeName HashTypeName) (ht *HashType, err error) {
	return NewByName(hashTypeName)
}

func NewByName(hashTypeName HashTypeName) (ht *HashType, err error) {
	x := HashTypeName(strings.ToUpper(string(hashTypeName)))
	switch x {
	case Name_FileExistence:
		ht = &HashType{id: Id_FileExistence, name: Name_FileExistence, isBinary: false, sumSize: hs.HashSize_None}
	case Name_FileSize:
		ht = &HashType{id: Id_FileSize, name: Name_FileSize, isBinary: false, sumSize: hs.HashSize_None}
	case Name_CRC32:
		ht = &HashType{id: Id_CRC32, name: Name_CRC32, isBinary: true, sumSize: hs.HashSize_CRC32}
	case Name_MD5:
		ht = &HashType{id: Id_MD5, name: Name_MD5, isBinary: true, sumSize: hs.HashSize_MD5}
	case Name_SHA1:
		ht = &HashType{id: Id_SHA1, name: Name_SHA1, isBinary: true, sumSize: hs.HashSize_SHA1}
	case Name_SHA256:
		ht = &HashType{id: Id_SHA256, name: Name_SHA256, isBinary: true, sumSize: hs.HashSize_SHA256}
	default:
		return nil, c.ErrorS1(ErrUnknownHashType, hashTypeName)
	}
	return ht, nil
}

func NewById(hashTypeId HashTypeId) (ht *HashType, err error) {
	switch hashTypeId {
	case Id_FileExistence:
		ht = &HashType{id: Id_FileExistence, name: Name_FileExistence, isBinary: false, sumSize: hs.HashSize_None}
	case Id_FileSize:
		ht = &HashType{id: Id_FileSize, name: Name_FileSize, isBinary: false, sumSize: hs.HashSize_None}
	case Id_CRC32:
		ht = &HashType{id: Id_CRC32, name: Name_CRC32, isBinary: true, sumSize: hs.HashSize_CRC32}
	case Id_MD5:
		ht = &HashType{id: Id_MD5, name: Name_MD5, isBinary: true, sumSize: hs.HashSize_MD5}
	case Id_SHA1:
		ht = &HashType{id: Id_SHA1, name: Name_SHA1, isBinary: true, sumSize: hs.HashSize_SHA1}
	case Id_SHA256:
		ht = &HashType{id: Id_SHA256, name: Name_SHA256, isBinary: true, sumSize: hs.HashSize_SHA256}
	default:
		return nil, c.ErrorA1(ErrUnknownHashType, hashTypeId)
	}
	return ht, nil
}

func (ht *HashType) ID() (id HashTypeId) {
	return ht.id
}

func (ht *HashType) Name() (name HashTypeName) {
	return ht.name
}

func (ht *HashType) IsBinary() (isBinary bool) {
	return ht.isBinary
}

func (ht *HashType) SumSize() (sumSize hs.HashSize) {
	return ht.sumSize
}
