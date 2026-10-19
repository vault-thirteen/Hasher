package ch

import "github.com/vault-thirteen/auxie/cio"

const (
	TplHr       = "--------------------------------------------------------------------------------"
	TplError    = "[ERROR] "
	TplOK       = "[  OK ] "
	TplSummary  = "Total files checked: %d. Good files: %d. Bad files: %d."
	MsgAllClear = "All clear."
)

type Check struct {
	Files   []*CheckedFile
	Counter CheckCounter
}

func NewCheck() (c *Check) {
	c = new(Check)
	c.Files = make([]*CheckedFile, 0, 256)
	return c
}

func (c *Check) AddFile(file *CheckedFile) {
	c.Counter.Total++

	if file == nil {
		c.Counter.Bad++
		return
	}

	c.Files = append(c.Files, file)

	if file.Ok {
		c.Counter.Good++
	} else {
		c.Counter.Bad++
	}
}

func (c *Check) PrintReport() {
	// If all files are good.
	if c.Counter.Bad == 0 {
		cio.PrintMessage(cio.FormatMessage(TplSummary, c.Counter.Total, c.Counter.Good, c.Counter.Bad))
		cio.PrintMessage(MsgAllClear)
		return
	}

	// Otherwise, show the details.
	lineFormat := "%s %s"

	// 1. Show good files.
	cio.PrintMessage(TplHr)
	for _, file := range c.Files {
		if file.Ok {
			cio.PrintMessage(cio.FormatMessage(lineFormat, TplOK, file.Path))
		}
	}

	// 2. Show other files.
	cio.PrintMessage(TplHr)
	for _, file := range c.Files {
		if !file.Ok {
			cio.PrintMessage(cio.FormatMessage(lineFormat, TplError, file.Path))
		}
	}

	// 3. Summary.
	cio.PrintMessage(TplHr)
	cio.PrintMessage(cio.FormatMessage(TplSummary, c.Counter.Total, c.Counter.Good, c.Counter.Bad))
}
