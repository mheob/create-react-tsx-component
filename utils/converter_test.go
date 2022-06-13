package utils

import (
	"testing"

	"github.com/mheob/create-react-tsx-component/tests"
	"github.com/stretchr/testify/assert"
)

func TestToPascalCase(t *testing.T) {
	toPascalCase(t)
}

func BenchmarkToPascalCase(b *testing.B) {
	tests.BenchmarkTest(b, toPascalCase)
}

func TestToCamelCase(t *testing.T) {
	toLowerCamelCase(t)
}

func BenchmarkToCamelCase(b *testing.B) {
	tests.BenchmarkTest(b, toLowerCamelCase)
}

func TestToSnakeCase(t *testing.T) {
	toSnakeCase(t)
}

func BenchmarkToSnakeCase(b *testing.B) {
	tests.BenchmarkTest(b, toSnakeCase)
}

func TestToKebabCase(t *testing.T) {
	toKebabCase(t)
}

func BenchmarkToKebabCase(b *testing.B) {
	tests.BenchmarkTest(b, toKebabCase)
}

func TestCharIsDelimiter(t *testing.T) {
	checkCharIsDelimiter(t)
}

func BenchmarkCharIsDelimiter(b *testing.B) {
	tests.BenchmarkTest(b, convertFileNameToPascalCase)
}

func TestConvertFileNameToPascalCase(t *testing.T) {
	convertFileNameToPascalCase(t)
}

func BenchmarkConvertFileNameToPascalCase(b *testing.B) {
	tests.BenchmarkTest(b, convertFileNameToPascalCase)
}

func TestConvertFileNameToSnakeCase(t *testing.T) {
	convertFileNameToSnakeCase(t)
}

func BenchmarkConvertFileNameToSnakeCase(b *testing.B) {
	tests.BenchmarkTest(b, convertFileNameToSnakeCase)
}

func convertFileNameToPascalCase(tb testing.TB) {
	cases := tests.Cases{
		"":                       "",
		"test_case":              "TestCase",
		"test-case":              "TestCase",
		"test.case":              "TestCase",
		"test":                   "Test",
		"TestCase":               "TestCase",
		" test  case ":           "TestCase",
		"many_many_words":        "ManyManyWords",
		"many____delimiters":     "ManyDelimiters",
		"--outside-delimiters--": "OutsideDelimiters",
		"AnyKind of_string":      "AnyKindOfString",
		"numbers2And55with000":   "Numbers2And55With000",
		"JSONData":               "JSONData",
	}
	runFileNameChecks(cases, tb, false)
}

func convertFileNameToSnakeCase(tb testing.TB) {
	cases := tests.Cases{
		"":                       "",
		"test_case":              "test-case",
		"test-case":              "test-case",
		"test.case":              "test-case",
		"test":                   "test",
		"TestCase":               "test-case",
		" test  case ":           "test-case",
		"many_many_words":        "many-many-words",
		"many____delimiters":     "many-delimiters",
		"--outside-delimiters--": "outside-delimiters",
		"AnyKind of_string":      "any-kind-of-string",
		"numbers2And55with000":   "numbers-2-and-55-with-000",
		"JSONData":               "json-data",
	}
	runFileNameChecks(cases, tb, true)
}

func runFileNameChecks(cases tests.Cases, tb testing.TB, usesKebabCase bool) {
	for in, expected := range cases {
		result := ConvertFileName(in, usesKebabCase)
		assert.EqualValues(tb, expected, result)
	}
}

func toPascalCase(tb testing.TB) {
	cases := tests.Cases{
		"":                       "",
		"test_case":              "TestCase",
		"test-case":              "TestCase",
		"test.case":              "TestCase",
		"test":                   "Test",
		"TestCase":               "TestCase",
		" test  case ":           "TestCase",
		"many_many_words":        "ManyManyWords",
		"many____delimiters":     "ManyDelimiters",
		"--outside-delimiters--": "OutsideDelimiters",
		"AnyKind of_string":      "AnyKindOfString",
		"numbers2And55with000":   "Numbers2And55With000",
		"JSONData":               "JSONData",
	}
	runCaseChecks(cases, tb, ToPascalCase)
}

func toLowerCamelCase(tb testing.TB) {
	cases := tests.Cases{
		"":                       "",
		"test_case":              "testCase",
		"test-case":              "testCase",
		"test.case":              "testCase",
		"test":                   "test",
		"TestCase":               "testCase",
		" test  case ":           "testCase",
		"many_many_words":        "manyManyWords",
		"many____delimiters":     "manyDelimiters",
		"--outside-delimiters--": "outsideDelimiters",
		"AnyKind of_string":      "anyKindOfString",
		"numbers2And55with000":   "numbers2And55With000",
		"JSONData":               "jSONData", // FIXME: `jsonData` should be the expected case
	}
	runCaseChecks(cases, tb, ToCamelCase)
}

func toSnakeCase(tb testing.TB) {
	cases := tests.Cases{
		"":                       "",
		"test_case":              "test_case",
		"test-case":              "test_case",
		"test.case":              "test_case",
		"test":                   "test",
		"TestCase":               "test_case",
		" test  case ":           "test_case",
		"many_many_words":        "many_many_words",
		"many____delimiters":     "many_delimiters",
		"--outside-delimiters--": "outside_delimiters",
		"AnyKind of_string":      "any_kind_of_string",
		"numbers2And55with000":   "numbers_2_and_55_with_000",
		"JSONData":               "json_data",
	}
	runCaseChecks(cases, tb, ToSnakeCase)
}

func toKebabCase(tb testing.TB) {
	cases := tests.Cases{
		"":                       "",
		"test_case":              "test-case",
		"test-case":              "test-case",
		"test.case":              "test-case",
		"test":                   "test",
		"TestCase":               "test-case",
		" test  case ":           "test-case",
		"many_many_words":        "many-many-words",
		"many____delimiters":     "many-delimiters",
		"--outside-delimiters--": "outside-delimiters",
		"AnyKind of_string":      "any-kind-of-string",
		"numbers2And55with000":   "numbers-2-and-55-with-000",
		"JSONData":               "json-data",
	}
	runCaseChecks(cases, tb, ToKebabCase)
}

func runCaseChecks(cases tests.Cases, tb testing.TB, cb func(string) string) {
	for in, expected := range cases {
		result := cb(in)
		assert.EqualValues(tb, expected, result)
	}
}

func checkCharIsDelimiter(tb testing.TB) {
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
		result := charIsDelimiter(in)
		if result != out {
			tb.Errorf("%q (%t != %t)", in, result, out)
		}
	}
}
