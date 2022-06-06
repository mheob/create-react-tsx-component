package utils

import "strings"

func toCamelCase(str string, initCase bool) string {
	str = strings.TrimSpace(str)
	if str == "" {
		return str
	}

	builder := strings.Builder{}
	builder.Grow(len(str))
	capNext := initCase

	for index, char := range []byte(str) {
		charIsCap := char >= 'A' && char <= 'Z'
		charIsLow := char >= 'a' && char <= 'z'
		charIsNum := char >= '0' && char <= '9'

		if capNext {
			if charIsLow {
				char += 'A'
				char -= 'a'
			}
		} else if index == 0 {
			if charIsCap {
				char += 'a'
				char -= 'A'
			}
		}

		if charIsCap || charIsLow {
			builder.WriteByte(char)
			capNext = false
		} else if charIsNum {
			builder.WriteByte(char)
			capNext = true
		} else {
			capNext = CharIsDelimiter(char)
		}
	}

	return builder.String()
}

func ToPascalCase(str string) string {
	return toCamelCase(str, true)
}

func ToCamelCase(str string) string {
	return toCamelCase(str, false)
}
