package main

import (
	"fmt"
	"log"
	"os"

	a "github.com/vault-thirteen/Hasher/pkg/Models/Action"
	cla "github.com/vault-thirteen/Hasher/pkg/Models/CommandLineArguments"
	ver "github.com/vault-thirteen/Versioneer"
)

const UsageHint = `Usage:
	[Action] [HashType] [ObjectType] [ObjectPath]

Examples:
	hasher.exe Calculate CRC32 Folder "Images\Cats"
	hasher.exe Check MD5 File MD5sums.txt

Notes:
	Possible actions: Calculate, Check.
	Possible hash types: CRC32, MD5, SHA256, Size.
	Possible object types: File, Folder or Directory.
	If action name is omitted, the default one is used.
	Default action is calculation.
	Hash sum checking is available for a sum file only.
	Letter case is not important.
	Calculator writes lines with standard line end (CR+LF).
	Checker reads lines with standard line end (CR+LF).
	Change directory (CD) to a working directory before usage.`

const (
	OutputFormat = "%v %s\r\n" // [1]=Sum, [2]=ObjectName.
)

func main() {
	args, err := cla.New()
	if err != nil {
		log.Println(err)
		showIntro()
		showUsage()
		os.Exit(1)
		return
	}

	switch args.Action().ID() {
	case a.IdCalculate:
		err = calculateHash(args)
	case a.IdCheck:
		err = checkHash(args)
	default:
		err = fmt.Errorf(a.ErrUnknown, args.Action())
	}
	mustBeNoError(err)
}

func mustBeNoError(err error) {
	if err != nil {
		panic(err)
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

func printHashLine(sum any, fileRelPath string) {
	fmt.Printf(OutputFormat, sum, fileRelPath)
}
