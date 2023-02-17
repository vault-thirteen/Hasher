package ot

import (
	"fmt"
	"strings"
)

const (
	IdFile   = 1
	IdFolder = 2
)

const (
	NameFile      = "FILE"
	NameFolder    = "FOLDER"
	NameDirectory = "DIRECTORY"
)

const (
	ErrUnknown = "object type is unknown: %v"
)

type ObjectType struct {
	id   byte
	name string
}

func New(objectTypeName string) (objectType *ObjectType, err error) {
	switch strings.ToUpper(objectTypeName) {
	case NameFile:
		return &ObjectType{
			id:   IdFile,
			name: NameFile,
		}, nil

	case NameFolder:
		return &ObjectType{
			id:   IdFolder,
			name: NameFolder,
		}, nil

	case NameDirectory:
		return &ObjectType{
			id:   IdFolder,
			name: NameFolder,
		}, nil

	default:
		return nil, fmt.Errorf(ErrUnknown, objectTypeName)
	}
}

func (ot *ObjectType) ID() (id byte) {
	return ot.id
}

func (ot *ObjectType) Name() (name string) {
	return ot.name
}
