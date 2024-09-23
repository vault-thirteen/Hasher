package ot

import (
	"strings"

	c "github.com/vault-thirteen/Hasher/pkg/Models/common"
)

const (
	ErrUnknownObjectType = "unknown object type"
)

type ObjectType struct {
	id   ObjectTypeId
	name ObjectTypeName
}

func New(objectTypeName ObjectTypeName) (ot *ObjectType, err error) {
	return NewByName(objectTypeName)
}

func NewByName(objectTypeName ObjectTypeName) (ot *ObjectType, err error) {
	x := ObjectTypeName(strings.ToUpper(string(objectTypeName)))
	switch x {
	case Name_File:
		ot = &ObjectType{id: Id_File, name: Name_File}
	case Name_Folder:
		ot = &ObjectType{id: Id_Folder, name: Name_Folder}
	case Name_Directory:
		ot = &ObjectType{id: Id_Folder, name: Name_Folder}
	default:
		return nil, c.ErrorS1(ErrUnknownObjectType, objectTypeName)
	}
	return ot, nil
}

func (ot *ObjectType) ID() (id ObjectTypeId) {
	return ot.id
}

func (ot *ObjectType) Name() (name ObjectTypeName) {
	return ot.name
}
