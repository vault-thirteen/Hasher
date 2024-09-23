package a

import (
	"strings"

	c "github.com/vault-thirteen/Hasher/pkg/Models/common"
)

const (
	ErrUnknownAction = "action is unknown"
)

type Action struct {
	id   ActionId
	name ActionName
}

func New(actionName string) (a *Action, err error) {
	x := ActionName(strings.ToUpper(actionName))
	switch x {
	case Name_Calculate:
		a = &Action{id: Id_Calculate, name: Name_Calculate}
	case Name_Check:
		a = &Action{id: Id_Check, name: Name_Check}
	case Name_Empty:
		a = &Action{id: Id_Default, name: Name_Default}
	default:
		return nil, c.ErrorA1(ErrUnknownAction, actionName)
	}
	return a, nil
}

func (a *Action) ID() (id ActionId) {
	return a.id
}

func (a *Action) Name() (name ActionName) {
	return a.name
}
