package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	cla "github.com/vault-thirteen/Hasher/pkg/Models/CommandLineArguments"
	ht "github.com/vault-thirteen/Hasher/pkg/Models/HashType"
	ot "github.com/vault-thirteen/Hasher/pkg/Models/ObjectType"
	"github.com/vault-thirteen/Hasher/pkg/hash"
	"github.com/vault-thirteen/auxie/reader"
	"github.com/vault-thirteen/errorz"
)

const (
	ErrFolderCheckIsNotPossible        = "folder check is not possible, use file check"
	ErrChecksumMismatch                = "checksum mismatch: '%s' vs '%s'"
	ErrDataIsDamaged                   = "data is damaged"
	ErrDamagedEntriesCountLimitReached = "too many damaged entries"
)

const (
	TplHr                    = "--------------------------------------------------------------------------------"
	TplError                 = "[ERROR] "
	TplOK                    = "[  OK ] "
	TplSummary               = "Total files checked: %d. Good files: %d. Bad files: %d."
	MsgErrorsWereFound       = "Errors were found !"
	MsgAllClear              = "All clear."
	DamagedEntriesCountLimit = 10
)

func checkHash(args *cla.CommandLineArguments) (err error) {
	switch args.ObjectType().ID() {
	case ot.IdFile:
		return checkHashesInFile(args)
	case ot.IdFolder:
		return errors.New(ErrFolderCheckIsNotPossible)
	default:
		return fmt.Errorf(ot.ErrUnknown, args.ObjectType())
	}
}

func checkHashesInFile(args *cla.CommandLineArguments) (err error) {
	switch args.HashType().ID() {
	case ht.IdCRC32:
		return checkCRC32HashesInFile(args.ObjectPath())
	case ht.IdMD5:
		return checkMD5HashesInFile(args.ObjectPath())
	case ht.IdSHA256:
		return checkSHA256HashesInFile(args.ObjectPath())
	default:
		return fmt.Errorf(ht.ErrUnknown, args.HashType())
	}
}

func checkCRC32HashesInFile(filePath string) (err error) {
	var f *os.File
	f, err = os.Open(filePath)
	if err != nil {
		return err
	}
	defer func() {
		derr := f.Close()
		if derr != nil {
			err = errorz.Combine(err, derr)
		}
	}()

	var (
		r                             = reader.New(f)
		buf                           []byte
		checkErr                      error
		file                          string
		nGood, nBad, nTotal, nDamaged int
	)

	fmt.Println(TplHr)

	for {
		if nDamaged > DamagedEntriesCountLimit {
			return errors.New(ErrDamagedEntriesCountLimitReached)
		}

		// Yes. Each line in the hash sum file must end with CR+LF.
		// If you are a user of Unix, Linux, OS X, Mac OS or any other OS with
		// non-standard line ends, please, check this article:
		// https://en.wikipedia.org/wiki/Newline
		buf, err = r.ReadLineEndingWithCRLF()
		if err == nil {
			file, checkErr = checkCRC32Line(buf)
			if checkErr != nil {
				nBad++
				if len(file) == 0 {
					nDamaged++
					file = "???"
				}
				fmt.Println(TplError, file)
			} else {
				nGood++
				fmt.Println(TplOK, file)
			}
			nTotal++
			continue
		}

		if err == io.EOF {
			break
		}

		return err
	}

	fmt.Println(TplHr)
	if (nTotal == nGood) && (nBad == 0) {
		fmt.Println(MsgAllClear)
	} else {
		fmt.Println(MsgErrorsWereFound)
	}
	fmt.Println(fmt.Sprintf(TplSummary, nTotal, nGood, nBad))

	return nil
}

func checkCRC32Line(buf []byte) (file string, err error) {
	if len(buf) < 10 {
		return file, errors.New(ErrDataIsDamaged)
	}

	var sumA string
	var sumTmp uint32
	sumA = strings.ToUpper(string(buf[:8]))
	file = strings.TrimSpace(string(buf[9:]))
	sumTmp, err = hash.GetFileHashCRC32(file)
	if err != nil {
		return file, err
	}

	if sumA != strings.ToUpper(strconv.FormatUint(uint64(sumTmp), 16)) {
		err = fmt.Errorf(ErrChecksumMismatch, sumA, strings.ToUpper(strconv.FormatUint(uint64(sumTmp), 16)))
		return file, err
	}

	return file, nil
}

func checkMD5HashesInFile(filePath string) (err error) {
	var f *os.File
	f, err = os.Open(filePath)
	if err != nil {
		return err
	}
	defer func() {
		derr := f.Close()
		if derr != nil {
			err = errorz.Combine(err, derr)
		}
	}()

	var (
		r                             = reader.New(f)
		buf                           []byte
		checkErr                      error
		file                          string
		nGood, nBad, nTotal, nDamaged int
	)

	fmt.Println(TplHr)

	for {
		if nDamaged > DamagedEntriesCountLimit {
			return errors.New(ErrDamagedEntriesCountLimitReached)
		}

		// Yes. Each line in the hash sum file must end with CR+LF.
		// If you are a user of Unix, Linux, OS X, Mac OS or any other OS with
		// non-standard line ends, please, check this article:
		// https://en.wikipedia.org/wiki/Newline
		buf, err = r.ReadLineEndingWithCRLF()
		if err == nil {
			file, checkErr = checkMD5Line(buf)
			if checkErr != nil {
				nBad++
				if len(file) == 0 {
					nDamaged++
					file = "???"
				}
				fmt.Println(TplError, file)
			} else {
				nGood++
				fmt.Println(TplOK, file)
			}
			nTotal++
			continue
		}

		if err == io.EOF {
			break
		}

		return err
	}

	fmt.Println(TplHr)
	if (nTotal == nGood) && (nBad == 0) {
		fmt.Println(MsgAllClear)
	} else {
		fmt.Println(MsgErrorsWereFound)
	}
	fmt.Println(fmt.Sprintf(TplSummary, nTotal, nGood, nBad))

	return nil
}

func checkMD5Line(buf []byte) (file string, err error) {
	if len(buf) < 34 {
		return file, errors.New(ErrDataIsDamaged)
	}

	var sumA string
	var sumTmp []byte
	sumA = strings.ToUpper(string(buf[:32]))
	file = strings.TrimSpace(string(buf[33:]))
	sumTmp, err = hash.GetFileHashMD5(file)
	if err != nil {
		return file, err
	}

	if sumA != strings.ToUpper(hex.EncodeToString(sumTmp)) {
		err = fmt.Errorf(ErrChecksumMismatch, sumA, strings.ToUpper(hex.EncodeToString(sumTmp)))
		return file, err
	}

	return file, nil
}

func checkSHA256HashesInFile(filePath string) (err error) {
	var f *os.File
	f, err = os.Open(filePath)
	if err != nil {
		return err
	}
	defer func() {
		derr := f.Close()
		if derr != nil {
			err = errorz.Combine(err, derr)
		}
	}()

	var (
		r                             = reader.New(f)
		buf                           []byte
		checkErr                      error
		file                          string
		nGood, nBad, nTotal, nDamaged int
	)

	fmt.Println(TplHr)

	for {
		if nDamaged > DamagedEntriesCountLimit {
			return errors.New(ErrDamagedEntriesCountLimitReached)
		}

		// Yes. Each line in the hash sum file must end with CR+LF.
		// If you are a user of Unix, Linux, OS X, Mac OS or any other OS with
		// non-standard line ends, please, check this article:
		// https://en.wikipedia.org/wiki/Newline
		buf, err = r.ReadLineEndingWithCRLF()
		if err == nil {
			file, checkErr = checkSHA256Line(buf)
			if checkErr != nil {
				nBad++
				if len(file) == 0 {
					nDamaged++
					file = "???"
				}
				fmt.Println(TplError, file)
			} else {
				nGood++
				fmt.Println(TplOK, file)
			}
			nTotal++
			continue
		}

		if err == io.EOF {
			break
		}

		return err
	}

	fmt.Println(TplHr)
	if (nTotal == nGood) && (nBad == 0) {
		fmt.Println(MsgAllClear)
	} else {
		fmt.Println(MsgErrorsWereFound)
	}
	fmt.Println(fmt.Sprintf(TplSummary, nTotal, nGood, nBad))

	return nil
}

func checkSHA256Line(buf []byte) (file string, err error) {
	if len(buf) < 66 {
		return file, errors.New(ErrDataIsDamaged)
	}

	var sumA string
	var sumTmp []byte
	sumA = strings.ToUpper(string(buf[:64]))
	file = strings.TrimSpace(string(buf[65:]))
	sumTmp, err = hash.GetFileHashSHA256(file)
	if err != nil {
		return file, err
	}

	if sumA != strings.ToUpper(hex.EncodeToString(sumTmp)) {
		err = fmt.Errorf(ErrChecksumMismatch, sumA, strings.ToUpper(hex.EncodeToString(sumTmp)))
		return file, err
	}

	return file, nil
}
