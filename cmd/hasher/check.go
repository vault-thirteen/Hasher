package main

import (
	"fmt"
	"io"
	"log"
	"os"

	ch "github.com/vault-thirteen/Hasher/pkg/Models/Check"
	cla "github.com/vault-thirteen/Hasher/pkg/Models/CommandLineArguments"
	h "github.com/vault-thirteen/Hasher/pkg/Models/Hashing"
	ot "github.com/vault-thirteen/Hasher/pkg/Models/ObjectType"
	c "github.com/vault-thirteen/Hasher/pkg/Models/common"
	ae "github.com/vault-thirteen/auxie/errors"
	"github.com/vault-thirteen/auxie/reader"
)

const (
	ErrFolderCheckIsNotPossible = "folder check is not possible, use file check"
	ErrFErrorOnLine             = "Error on line %v."
)

func checkHash(args *cla.CommandLineArguments) (results *ch.Check, err error) {
	switch args.ObjectType().ID() {
	case ot.Id_File:
		return checkHashesInFile(args)
	case ot.Id_Folder:
		return nil, c.Error(ErrFolderCheckIsNotPossible)
	default:
		return nil, c.ErrorA1(ot.ErrUnknownObjectType, args.ObjectType())
	}
}

func checkHashesInFile(args *cla.CommandLineArguments) (results *ch.Check, err error) {
	var hasher *h.Hashing
	hasher, err = h.NewHashing(args.HashType().ID())
	if err != nil {
		return nil, err
	}

	mainFilePath := args.ObjectPath()

	var f *os.File
	f, err = os.Open(mainFilePath)
	if err != nil {
		return nil, err
	}
	defer func() {
		derr := f.Close()
		if derr != nil {
			err = ae.Combine(err, derr)
		}
	}()

	var rdr = reader.New(f)
	var line []byte
	var lineN int = 1
	results = ch.NewCheck()
	var hashText string
	var hash any
	var result *ch.CheckedFile

	defer func() {
		if err != nil {
			log.Println(fmt.Sprintf(ErrFErrorOnLine, lineN))
		}
	}()

	for {
		line, err = rdr.ReadLineEndingWithCRLF()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}

		result = &ch.CheckedFile{}

		hashText, result.Path, err = hasher.ParseFileLine(line)
		if err != nil {
			return nil, err
		}

		hash, err = hasher.ParseHash(hashText)
		if err != nil {
			return nil, err
		}

		result.Ok, err = hasher.Verify(result.Path, hash)
		if err != nil {
			return nil, err
		}

		results.AddFile(result)
		lineN++
		continue
	}

	return results, nil
}
