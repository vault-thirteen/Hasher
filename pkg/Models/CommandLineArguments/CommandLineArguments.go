package cla

import (
	"errors"
	"fmt"
	"os"

	a "github.com/vault-thirteen/Hasher/pkg/Models/Action"
	ht "github.com/vault-thirteen/Hasher/pkg/Models/HashType"
	ot "github.com/vault-thirteen/Hasher/pkg/Models/ObjectType"
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
		computedArgs = append(computedArgs, a.NameDefault)
		computedArgs = append(computedArgs, os.Args[1])
		computedArgs = append(computedArgs, os.Args[2])
		computedArgs = append(computedArgs, os.Args[3])
	default:
		return nil, errors.New(ErrSyntax)
	}

	// Stop those fools who are blind.
	cla = &CommandLineArguments{
		objectPath: computedArgs[3],
	}
	cla.objectType, err = ot.New(computedArgs[2])
	if err != nil {
		return nil, err
	}

	// Fill the rest data.
	cla.hashType, err = ht.New(computedArgs[1])
	if err != nil {
		return nil, err
	}

	cla.action, err = a.New(computedArgs[0])
	if err != nil {
		return nil, err
	}

	// Some of the algorithms work with non-existent objects,
	// so sometimes absent objects are not an error.
	if (cla.action.ID() == a.IdCheck) ||
		(cla.hashType.ID() != ht.IdFileExistence) {
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
	case ot.IdFile:
		objectExists, err = file.FileExists(cla.ObjectPath())
	case ot.IdFolder:
		objectExists, err = file.FolderExists(cla.ObjectPath())
	default:
		return fmt.Errorf(ot.ErrUnknown, cla.objectType)
	}
	if err != nil {
		return err
	}
	if !objectExists {
		return errors.New(ErrObjectIsNotFound)
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
