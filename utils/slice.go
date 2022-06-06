package utils

func CharIsDelimiter(char byte) bool {
	delimiter := [...]byte{' ', '_', '-', '.'}

	for _, value := range delimiter {
		if value == char {
			return true
		}
	}

	return false
}
