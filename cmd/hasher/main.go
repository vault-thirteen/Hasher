package main

import (
	"fmt"
	"log"
	"os"

	a "github.com/vault-thirteen/Hasher/pkg/Models/Action"
	ch "github.com/vault-thirteen/Hasher/pkg/Models/Check"
	cla "github.com/vault-thirteen/Hasher/pkg/Models/CommandLineArguments"
	c "github.com/vault-thirteen/Hasher/pkg/Models/common"
	ver "github.com/vault-thirteen/auxie/Versioneer"
)

const UsageHint = `Usage:
	[Action] [HashType] [ObjectType] [ObjectPath]

Examples:
	hasher.exe Calculate CRC32 Folder "Images\Cats"
	hasher.exe Check MD5 File MD5sums.txt

Notes:
	Possible actions: Calculate, Check.
	Possible hash types: CRC32, MD5, SHA256, Size, Existence.
	Possible object types: File, Folder or Directory.
	If action name is omitted, the default one is used.
	Default action is calculation.
	Hash sum checking is available for a sum file only.
	Letter case is not important.
	Calculator writes lines with standard line end (CR+LF).
	Checker reads lines with standard line end (CR+LF).
	Change directory (CD) to a working directory before usage.`

func main() {
	args, err := cla.New()
	if err != nil {
		showIntro()
		log.Println(err)
		showUsage()
		os.Exit(1)
		return
	}

	err = work(args)
	mustBeNoError(err)
}

func work(args *cla.CommandLineArguments) (err error) {
	switch args.Action().ID() {
	case a.Id_Calculate:
		err = calculateHash(args)
		if err != nil {
			return err
		}

	case a.Id_Check:
		var checkResults *ch.Check
		checkResults, err = checkHash(args)
		if err != nil {
			return err
		}

		checkResults.PrintReport()

	default:
		err = c.ErrorA1(a.ErrUnknownAction, args.Action())
		return err
	}

	return nil
}

func mustBeNoError(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func showIntro() {
	versioneer, err := ver.New()
	mustBeNoError(err)
	versioneer.ShowIntroText("")
	versioneer.ShowComponentsInfoText()
	fmt.Println()
}

func showUsage() {
	fmt.Println(UsageHint)
}
