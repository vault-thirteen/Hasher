package a

type ActionName string

const (
	Name_Calculate = ActionName("CALCULATE")
	Name_Check     = ActionName("CHECK")
	Name_Empty     = ActionName("")

	Name_Default = Name_Calculate
)

func (a ActionName) ToString() string {
	return string(a)
}
