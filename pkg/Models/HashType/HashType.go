package ht

import (
	"fmt"
	"strings"
)

const (
	IdCRC32    = 1
	IdMD5      = 2
	IdSHA256   = 3
	IdFileSize = 4
)

const (
	NameCRC32    = "CRC32"
	NameMD5      = "MD5"
	NameSHA256   = "SHA256"
	NameFileSize = "SIZE"
)

const (
	ErrUnknown = "hash type is unknown: %v"
)

type HashType struct {
	id   byte
	name string
}

func New(hashTypeName string) (hashType *HashType, err error) {
	switch strings.ToUpper(hashTypeName) {
	case NameCRC32:
		return &HashType{
			id:   IdCRC32,
			name: NameCRC32,
		}, nil

	case NameMD5:
		return &HashType{
			id:   IdMD5,
			name: NameMD5,
		}, nil

	case NameSHA256:
		return &HashType{
			id:   IdSHA256,
			name: NameSHA256,
		}, nil

	case NameFileSize:
		return &HashType{
			id:   IdFileSize,
			name: NameFileSize,
		}, nil

	default:
		return nil, fmt.Errorf(ErrUnknown, hashTypeName)
	}
}

func (ht *HashType) ID() (id byte) {
	return ht.id
}

func (ht *HashType) Name() (name string) {
	return ht.name
}
