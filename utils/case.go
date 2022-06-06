package utils

import "strings"

func toCamelCase(s string, initCase bool) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}

	builder := strings.Builder{}
	builder.Grow(len(s))
	capNext := initCase

	for i, char := range []byte(s) {
		charIsCap := char >= 'A' && char <= 'Z'
		charIsLow := char >= 'a' && char <= 'z'
		charIsNum := char >= '0' && char <= '9'

		if capNext {
			if charIsLow {
				char += 'A'
				char -= 'a'
			}
		} else if i == 0 {
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
