package numberCheck

func CheckNumber(n int) string {
	if n <= 0 {
		return "small number"
	}
	if n > 0 && n <= 100 {
		return "medium number"
	} else {
		return "big number"
	}
}
