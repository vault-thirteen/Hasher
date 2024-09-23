package common

import "fmt"

const (
	OutputFormat = "%s %s\r\n" // [1]=Sum, [2]=ObjectName.
)

func FormatBooleanAsNumber(b bool) string {
	if b == true {
		return "1"
	} else {
		return "0"
	}
}

func PrintHashLine(hashSumText string, fileRelPath string) {
	fmt.Printf(OutputFormat, hashSumText, fileRelPath)
}
