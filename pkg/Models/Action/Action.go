package a

import (
	"fmt"
	"strings"
)

const (
	IdCalculate = 1
	IdCheck     = 2
	IdDefault   = IdCalculate
)

const (
	NameCalculate = "CALCULATE"
	NameCheck     = "CHECK"
	NameEmpty     = ""
	NameDefault   = NameCalculate
)

const (
	ErrUnknown = "action is unknown: %v"
)

type Action struct {
	id   byte
	name string
}

func New(actionName string) (action *Action, err error) {
	switch strings.ToUpper(actionName) {
	case NameCalculate:
		return &Action{
			id:   IdCalculate,
			name: NameCalculate,
		}, nil

	case NameCheck:
		return &Action{
			id:   IdCheck,
			name: NameCheck,
		}, nil

	case NameEmpty:
		return &Action{
			id:   IdDefault,
			name: NameDefault,
		}, nil

	default:
		return nil, fmt.Errorf(ErrUnknown, actionName)
	}
}

func (a *Action) ID() (id byte) {
	return a.id
}

func (a *Action) Name() (name string) {
	return a.name
}
