package common

import (
	"errors"
	"fmt"
)

// Go language is very lame. Methods for errors are placed in different
// packages making the process of error creation very boring. These functions
// try to make this process a little bit easier.

// Placeholders.
const (
	FmtPlaceHolder_String = "%s"
	FmtPlaceHolder_Any    = "%v"
)

// Separators.
const (
	FmtErrorSeparatorA = ": "
	FmtErrorSeparatorB = "; "
)

// FmtErrorExtender_S1 is an addition to an error message for showing a single
// textual value after the message.
const FmtErrorExtender_S1 = FmtErrorSeparatorA + FmtPlaceHolder_String

// FmtErrorExtender_S2 is an addition to an error message for showing two
// textual values after the message.
const FmtErrorExtender_S2 = FmtErrorSeparatorA + FmtPlaceHolder_String + FmtErrorSeparatorB + FmtPlaceHolder_String

// FmtErrorExtender_A1 is an addition to an error message for showing a single
// value (of any type) after the message.
const FmtErrorExtender_A1 = FmtErrorSeparatorA + FmtPlaceHolder_Any

func Error(errMsg string) error {
	return errors.New(errMsg)
}

func ErrorS1(errMsg string, strValue any) error {
	return fmt.Errorf(errMsg+FmtErrorExtender_S1, strValue)
}

func ErrorS2(errMsg string, strValue1, strValue2 any) error {
	return fmt.Errorf(errMsg+FmtErrorExtender_S2, strValue1, strValue2)
}

func ErrorA1(errMsg string, anyValue any) error {
	return fmt.Errorf(errMsg+FmtErrorExtender_A1, anyValue)
}
