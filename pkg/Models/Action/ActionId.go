package a

type ActionId byte

const (
	Id_Calculate = ActionId(1)
	Id_Check     = ActionId(2)

	Id_Default = Id_Calculate
)
