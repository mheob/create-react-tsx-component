package utils

import "testing"

func RunChecks(cases [][]string, tb testing.TB, cb func(string) string) {
	for _, index := range cases {
		in := index[0]
		out := index[1]
		result := cb(in)
		if result != out {
			tb.Errorf("%q (%q != %q)", in, result, out)
		}
	}
}

func BenchmarkTest(b *testing.B, cb func(testing.TB)) {
	for n := 0; n < b.N; n++ {
		cb(b)
	}
}
