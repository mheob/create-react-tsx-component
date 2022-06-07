package utils

import "testing"

func TestCharIsDelimiter(t *testing.T) {
	charIsDelimiter(t)
}

func BenchmarkCharIsDelimiter(b *testing.B) {
	BenchmarkTest(b, convertFileNameToPascalCase)
}

func charIsDelimiter(tb testing.TB) {
	cases := map[byte]bool{
		'1': false,
		'a': false,
		'A': false,
		'?': false,
		'=': false,
		'+': false,
		' ': true,
		'_': true,
		'-': true,
		'.': true,
	}

	for in, out := range cases {
		result := CharIsDelimiter(in)
		if result != out {
			tb.Errorf("%q (%t != %t)", in, result, out)
		}
	}
}
