package ch

import "fmt"

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
		c.Counter.Damaged++
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
	nBD := c.Counter.Bad + c.Counter.Damaged

	// IF all files are good.
	if nBD == 0 {
		fmt.Println(fmt.Sprintf(TplSummary, c.Counter.Total, c.Counter.Good, nBD))
		fmt.Println(MsgAllClear)
		return
	}

	// Otherwise, show the details.
	lineFormat := "%s %s"

	// 1. Show good files.
	fmt.Println(TplHr)
	for _, file := range c.Files {
		if file.Ok {
			fmt.Println(fmt.Sprintf(lineFormat, TplOK, file.Path))
		}
	}

	// 2. Show other files.
	fmt.Println(TplHr)
	for _, file := range c.Files {
		if !file.Ok {
			fmt.Println(fmt.Sprintf(lineFormat, TplError, file.Path))
		}
	}

	// 3. Summary.
	fmt.Println(TplHr)
	fmt.Println(fmt.Sprintf(TplSummary, c.Counter.Total, c.Counter.Good, nBD))
}
