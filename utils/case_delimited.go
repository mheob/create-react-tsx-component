package utils

import "strings"

func ToSnakeCase(str string) string {
	return toDelimited(str, '_')
}

func ToKebabCase(str string) string {
	return toDelimited(str, '-')
}

func toDelimited(str string, delimiter uint8) string {
	str = strings.TrimSpace(str)
	if str == "" {
		return str
	}

	builder := strings.Builder{}
	builder.Grow(len(str) + 2) // nominal 2 bytes of extra space for inserted delimiters

	for index, char := range []byte(str) {
		// treat acronyms as words, eg for JSONData -> JSON is a whole word
		if index+1 < len(str) {
			next := str[index+1]

			if CharIsDelimiter(char) && CharIsDelimiter(next) {
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
					if prevIsCap := index > 0 &&
						str[index-1] >= 'A' &&
						str[index-1] <= 'Z'; prevIsCap {
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

		if CharIsDelimiter(char) {
			builder.WriteByte(delimiter)
		} else {
			builder.WriteByte(char)
		}
	}

	output := builder.String()

	return strings.ToLower(output)
}
