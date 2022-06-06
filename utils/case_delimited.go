package utils

import "strings"

func ToSnakeCase(str string) string {
	return toDelimited(str, '_')
}

func ToKebabCase(str string) string {
	return toDelimited(str, '-')
}

func toDelimited(s string, delimiter uint8) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}

	builder := strings.Builder{}
	builder.Grow(len(s) + 2) // nominal 2 bytes of extra space for inserted delimiters

	for i, char := range []byte(s) {
		// treat acronyms as words, eg for JSONData -> JSON is a whole word
		if i+1 < len(s) {
			next := s[i+1]

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
					if prevIsCap := i > 0 &&
						s[i-1] >= 'A' &&
						s[i-1] <= 'Z'; prevIsCap {
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
