package utils

import "strings"

func ToPascalCase(str string) string {
	return toCamelCase(str, true)
}

func ToCamelCase(str string) string {
	return toCamelCase(str, false)
}

func ToSnakeCase(str string) string {
	return toDelimited(str, '_')
}

func ToKebabCase(str string) string {
	return toDelimited(str, '-')
}

func ConvertFileName(name string, usesKebabCase bool) string {
	if usesKebabCase {
		return ToKebabCase(name)
	}

	return ToPascalCase(name)
}

func toCamelCase(str string, isUpperFirstChar bool) string {
	str = strings.TrimSpace(str)
	if str == "" {
		return str
	}

	builder := strings.Builder{}
	builder.Grow(len(str))
	capNext := isUpperFirstChar

	for i, char := range []byte(str) {
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
			capNext = charIsDelimiter(char)
		}
	}

	return builder.String()
}

func toDelimited(str string, delimiter byte) string {
	str = strings.TrimSpace(str)
	if str == "" {
		return str
	}

	builder := strings.Builder{}
	builder.Grow(len(str) + 2) // nominal 2 bytes of extra space for inserted delimiters

	for i, char := range []byte(str) {
		// treat acronyms as words, e.g. for JSONData -> JSON is a whole word
		if i+1 < len(str) {
			next := str[i+1]

			if charIsDelimiter(char) && charIsDelimiter(next) {
				continue
			}

			charIsCap := char >= 'A' && char <= 'Z'
			charIsLow := char >= 'a' && char <= 'z'
			charIsNum := char >= '0' && char <= '9'
			nextIsCap := next >= 'A' && next <= 'Z'
			nextIsLow := next >= 'a' && next <= 'z'
			nextIsNum := next >= '0' && next <= '9'

			// add underscore if next letter case type is changed
			if (charIsCap && (nextIsLow || nextIsNum)) ||
				(charIsLow && (nextIsCap || nextIsNum)) ||
				(charIsNum && (nextIsCap || nextIsLow)) {
				if charIsCap && nextIsLow {
					if prevIsCap := i > 0 &&
						str[i-1] >= 'A' &&
						str[i-1] <= 'Z'; prevIsCap {
						builder.WriteByte(delimiter)
					}
				}

				builder.WriteByte(char)

				if charIsLow || charIsNum || nextIsNum {
					builder.WriteByte(delimiter)
				}

				continue
			}
		}

		if charIsDelimiter(char) {
			builder.WriteByte(delimiter)
		} else {
			builder.WriteByte(char)
		}
	}

	output := builder.String()

	return strings.ToLower(output)
}

func charIsDelimiter(char byte) bool {
	delimiter := []byte{' ', '_', '-', '.'}

	for _, value := range delimiter {
		if value == char {
			return true
		}
	}

	return false
}
