package main

import (
	"path/filepath"

	cla "github.com/vault-thirteen/Hasher/pkg/Models/CommandLineArguments"
	h "github.com/vault-thirteen/Hasher/pkg/Models/Hashing"
	ot "github.com/vault-thirteen/Hasher/pkg/Models/ObjectType"
	c "github.com/vault-thirteen/Hasher/pkg/Models/common"
)

func calculateHash(args *cla.CommandLineArguments) (err error) {
	switch args.ObjectType().ID() {
	case ot.Id_File:
		return calculateHashOfFile(args)
	case ot.Id_Folder:
		return calculateHashOfFolder(args)
	default:
		return c.ErrorA1(ot.ErrUnknownObjectType, args.ObjectType())
	}
}

func calculateHashOfFile(args *cla.CommandLineArguments) (err error) {
	filePath := args.ObjectPath()

	var hasher *h.Hashing
	hasher, err = h.NewHashing(args.HashType().ID())
	if err != nil {
		return err
	}

	var result *h.HashingResult
	result, err = hasher.Calculate(filePath)
	if err != nil {
		return err
	}

	c.PrintHashLine(result.ToString(), filepath.Base(filePath))
	return nil
}

func calculateHashOfFolder(args *cla.CommandLineArguments) (err error) {
	basePath := args.ObjectPath()

	var hasher *h.Hashing
	hasher, err = h.NewHashing(args.HashType().ID())
	if err != nil {
		return err
	}

	err = filepath.Walk(basePath, hasher.WalkerFn)
	if err != nil {
		return err
	}

	return nil
}
