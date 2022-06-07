package utils

import "testing"

type Cases = map[string]string

func BenchmarkTest(b *testing.B, cb func(testing.TB)) {
	for n := 0; n < b.N; n++ {
		cb(b)
	}
}
