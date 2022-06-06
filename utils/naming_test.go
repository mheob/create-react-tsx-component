package utils

import "testing"

func TestConvertFileNameToPascalCase(t *testing.T) {
	convertFileNameToPascalCase(t)
}

func BenchmarkConvertFileNameToPascalCase(b *testing.B) {
	BenchmarkTest(b, convertFileNameToPascalCase)
}

func TestConvertFileNameToSnakeCase(t *testing.T) {
	convertFileNameToSnakeCase(t)
}

func BenchmarkConvertFileNameToSnakeCase(b *testing.B) {
	BenchmarkTest(b, convertFileNameToSnakeCase)
}

func convertFileNameToPascalCase(tb testing.TB) {
	cases := [][]string{
		{"test_case", "TestCase"},
		{"test-case", "TestCase"},
		{"test.case", "TestCase"},
		{"test", "Test"},
		{"TestCase", "TestCase"},
		{" test  case ", "TestCase"},
		{"", ""},
		{"many_many_words", "ManyManyWords"},
		{"AnyKind of_string", "AnyKindOfString"},
		{"numbers2And55with000", "Numbers2And55With000"},
		{"JSONData", "JSONData"},
	}
	fileNameChecks(cases, tb, false)
}

func convertFileNameToSnakeCase(tb testing.TB) {
	cases := [][]string{
		{"", ""},
		{"test_case", "test-case"},
		{"test-case", "test-case"},
		{"test.case", "test-case"},
		{"test", "test"},
		{"TestCase", "test-case"},
		{" test  case ", "test-case"},
		{"many_many_words", "many-many-words"},
		{"AnyKind of_string", "any-kind-of-string"},
		{"numbers2And55with000", "numbers-2-and-55-with-000"},
		{"JSONData", "json-data"},
	}
	fileNameChecks(cases, tb, true)
}

func fileNameChecks(cases [][]string, tb testing.TB, usesKebabCase bool) {
	for _, index := range cases {
		in := index[0]
		out := index[1]
		result := ConvertFileName(in, usesKebabCase)
		if result != out {
			tb.Errorf("%q (%q != %q)", in, result, out)
		}
	}
}
