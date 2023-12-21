package strconv

func FormatBooleanAsNumber(b bool) string {
	if b == true {
		return "1"
	} else {
		return "0"
	}
}
