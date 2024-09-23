package cla

import (
	"os"

	a "github.com/vault-thirteen/Hasher/pkg/Models/Action"
	ht "github.com/vault-thirteen/Hasher/pkg/Models/HashType"
	ot "github.com/vault-thirteen/Hasher/pkg/Models/ObjectType"
	c "github.com/vault-thirteen/Hasher/pkg/Models/common"
	"github.com/vault-thirteen/auxie/file"
)

const (
	ErrSyntax           = "syntax error"
	ErrObjectIsNotFound = "object is not found"
)

type CommandLineArguments struct {
	action     *a.Action
	hashType   *ht.HashType
	objectType *ot.ObjectType
	objectPath string
}

func New() (cla *CommandLineArguments, err error) {
	// Is the used syntax full or short ?
	var computedArgs = make([]string, 0, 4)
	switch len(os.Args) {
	case 5:
		computedArgs = append(computedArgs, os.Args[1])
		computedArgs = append(computedArgs, os.Args[2])
		computedArgs = append(computedArgs, os.Args[3])
		computedArgs = append(computedArgs, os.Args[4])

	case 4:
		computedArgs = append(computedArgs, a.Name_Default.ToString())
		computedArgs = append(computedArgs, os.Args[1])
		computedArgs = append(computedArgs, os.Args[2])
		computedArgs = append(computedArgs, os.Args[3])

	default:
		return nil, c.Error(ErrSyntax)
	}

	// Stop those fools who are blind.
	cla = &CommandLineArguments{
		objectPath: computedArgs[3],
	}

	cla.objectType, err = ot.New(ot.ObjectTypeName(computedArgs[2]))
	if err != nil {
		return nil, err
	}

	// Fill the rest data.
	cla.hashType, err = ht.New(ht.HashTypeName(computedArgs[1]))
	if err != nil {
		return nil, err
	}

	cla.action, err = a.New(computedArgs[0])
	if err != nil {
		return nil, err
	}

	// Some of the algorithms work with non-existent objects,
	// so sometimes absent objects are not an error.
	if (cla.action.ID() == a.Id_Check) ||
		(cla.hashType.ID() != ht.Id_FileExistence) {
		err = ensureObjectExists(cla)
		if err != nil {
			return nil, err
		}
	}

	return cla, nil
}

func ensureObjectExists(cla *CommandLineArguments) (err error) {
	var objectExists bool
	switch cla.objectType.ID() {
	case ot.Id_File:
		objectExists, err = file.FileExists(cla.ObjectPath())
	case ot.Id_Folder:
		objectExists, err = file.FolderExists(cla.ObjectPath())
	default:
		return c.ErrorS1(ot.ErrUnknownObjectType, cla.objectType)
	}
	if err != nil {
		return err
	}
	if !objectExists {
		return c.Error(ErrObjectIsNotFound)
	}

	return nil
}

func (cla *CommandLineArguments) Action() (action *a.Action) {
	return cla.action
}

func (cla *CommandLineArguments) HashType() (hashType *ht.HashType) {
	return cla.hashType
}

func (cla *CommandLineArguments) ObjectType() (objectType *ot.ObjectType) {
	return cla.objectType
}

func (cla *CommandLineArguments) ObjectPath() (objectPath string) {
	return cla.objectPath
}
