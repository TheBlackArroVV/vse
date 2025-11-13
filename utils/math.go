package utils

func AbsInt(number int) int {
	if number > 0 {
		return number
	}

	return number * -1
}
