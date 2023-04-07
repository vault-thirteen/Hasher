package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	cla "github.com/vault-thirteen/Hasher/pkg/Models/CommandLineArguments"
	ht "github.com/vault-thirteen/Hasher/pkg/Models/HashType"
	ot "github.com/vault-thirteen/Hasher/pkg/Models/ObjectType"
	"github.com/vault-thirteen/Hasher/pkg/hash"
)

func calculateHash(args *cla.CommandLineArguments) (err error) {
	switch args.ObjectType().ID() {
	case ot.IdFile:
		return calculateHashOfFile(args)
	case ot.IdFolder:
		return calculateHashOfFolder(args)
	default:
		return fmt.Errorf(ot.ErrUnknown, args.ObjectType())
	}
}

func calculateHashOfFile(args *cla.CommandLineArguments) (err error) {
	switch args.HashType().ID() {
	case ht.IdCRC32:
		return calculateFileHashCRC32(args)
	case ht.IdMD5:
		return calculateFileHashMD5(args)
	case ht.IdSHA256:
		return calculateFileHashSHA256(args)
	case ht.IdFileSize:
		return calculateFileHashFileSize(args)
	default:
		return fmt.Errorf(ht.ErrUnknown, args.HashType())
	}
}

func calculateFileHashCRC32(args *cla.CommandLineArguments) (err error) {
	var sum uint32
	sum, err = hash.GetFileHashCRC32(args.ObjectPath())
	if err != nil {
		return err
	}

	printHashLine(strings.ToUpper(strconv.FormatUint(uint64(sum), 16)), filepath.Base(args.ObjectPath()))
	return nil
}

func calculateFileHashMD5(args *cla.CommandLineArguments) (err error) {
	var sum []byte
	sum, err = hash.GetFileHashMD5(args.ObjectPath())
	if err != nil {
		return err
	}

	printHashLine(strings.ToUpper(hex.EncodeToString(sum[:])), filepath.Base(args.ObjectPath()))
	return nil
}

func calculateFileHashSHA256(args *cla.CommandLineArguments) (err error) {
	var sum []byte
	sum, err = hash.GetFileHashSHA256(args.ObjectPath())
	if err != nil {
		return err
	}

	printHashLine(strings.ToUpper(hex.EncodeToString(sum[:])), filepath.Base(args.ObjectPath()))
	return nil
}

func calculateFileHashFileSize(args *cla.CommandLineArguments) (err error) {
	var sum int64
	sum, err = hash.GetFileHashFileSize(args.ObjectPath())
	if err != nil {
		return err
	}

	printHashLine(strconv.FormatInt(sum, 10), filepath.Base(args.ObjectPath()))
	return nil
}

func calculateHashOfFolder(args *cla.CommandLineArguments) (err error) {
	switch args.HashType().ID() {
	case ht.IdCRC32:
		return calculateFolderHashCRC32(args)
	case ht.IdMD5:
		return calculateFolderHashMD5(args)
	case ht.IdSHA256:
		return calculateFolderHashSHA256(args)
	case ht.IdFileSize:
		return calculateFolderHashFileSize(args)
	default:
		return fmt.Errorf(ht.ErrUnknown, args.HashType())
	}
}

func calculateFolderHashCRC32(args *cla.CommandLineArguments) (err error) {
	basePath := args.ObjectPath()

	err = filepath.Walk(basePath, crc32DirWalker)
	if err != nil {
		return err
	}

	return nil
}

func crc32DirWalker(path string, fi os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if fi.IsDir() {
		return nil
	}

	var sum uint32
	sum, err = hash.GetFileHashCRC32(path)
	if err != nil {
		return err
	}

	printHashLine(strings.ToUpper(strconv.FormatUint(uint64(sum), 16)), path)

	return nil
}

func calculateFolderHashMD5(args *cla.CommandLineArguments) (err error) {
	basePath := args.ObjectPath()

	err = filepath.Walk(basePath, md5DirWalker)
	if err != nil {
		return err
	}

	return nil
}

func md5DirWalker(path string, fi os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if fi.IsDir() {
		return nil
	}

	var sum []byte
	sum, err = hash.GetFileHashMD5(path)
	if err != nil {
		return err
	}

	printHashLine(strings.ToUpper(hex.EncodeToString(sum[:])), path)

	return nil
}

func calculateFolderHashSHA256(args *cla.CommandLineArguments) (err error) {
	basePath := args.ObjectPath()

	err = filepath.Walk(basePath, sha256DirWalker)
	if err != nil {
		return err
	}

	return nil
}

func calculateFolderHashFileSize(args *cla.CommandLineArguments) (err error) {
	basePath := args.ObjectPath()

	err = filepath.Walk(basePath, fileSizeDirWalker)
	if err != nil {
		return err
	}

	return nil
}

func sha256DirWalker(path string, fi os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if fi.IsDir() {
		return nil
	}

	var sum []byte
	sum, err = hash.GetFileHashSHA256(path)
	if err != nil {
		return err
	}

	printHashLine(strings.ToUpper(hex.EncodeToString(sum[:])), path)

	return nil
}

func fileSizeDirWalker(path string, fi os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if fi.IsDir() {
		return nil
	}

	var sum int64
	sum, err = hash.GetFileHashFileSize(path)
	if err != nil {
		return err
	}

	printHashLine(strconv.FormatInt(sum, 10), path)

	return nil
}
