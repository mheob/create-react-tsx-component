package utils

import "testing"

func TestToPascalCase(t *testing.T) {
	toPascalCase(t)
}

func BenchmarkToPascalCase(b *testing.B) {
	BenchmarkTest(b, toPascalCase)
}

func TestToCamelCase(t *testing.T) {
	toLowerCamelCase(t)
}

func BenchmarkToCamelCase(b *testing.B) {
	BenchmarkTest(b, toLowerCamelCase)
}

func toPascalCase(tb testing.TB) {
	cases := [][]string{
		{"test_case", "TestCase"},
		{"test.case", "TestCase"},
		{"test", "Test"},
		{"TestCase", "TestCase"},
		{" test  case ", "TestCase"},
		{"", ""},
		{"many_many_words", "ManyManyWords"},
		{"AnyKind of_string", "AnyKindOfString"},
		{"odd-fix", "OddFix"},
		{"numbers2And55with000", "Numbers2And55With000"},
	}
	RunChecks(cases, tb, ToPascalCase)
}

func toLowerCamelCase(tb testing.TB) {
	cases := [][]string{
		{"foo-bar", "fooBar"},
		{"TestCase", "testCase"},
		{"", ""},
		{"AnyKind of_string", "anyKindOfString"},
		{"AnyKind.of-string", "anyKindOfString"},
		{"some string", "someString"},
		{" some string", "someString"},
	}
	RunChecks(cases, tb, ToCamelCase)
}
