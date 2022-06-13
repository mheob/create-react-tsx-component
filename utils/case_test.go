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
		"":                     "",
		"test_case":            "TestCase",
		"test-case":            "TestCase",
		"test.case":            "TestCase",
		"test":                 "Test",
		"TestCase":             "TestCase",
		" test  case ":         "TestCase",
		"many_many_words":      "ManyManyWords",
		"AnyKind of_string":    "AnyKindOfString",
		"numbers2And55with000": "Numbers2And55With000",
		"JSONData":             "JSONData",
	}
	runFileNameChecks(cases, tb, false)
}

func convertFileNameToSnakeCase(tb testing.TB) {
	cases := tests.Cases{
		"":                     "",
		"test_case":            "test-case",
		"test-case":            "test-case",
		"test.case":            "test-case",
		"test":                 "test",
		"TestCase":             "test-case",
		" test  case ":         "test-case",
		"many_many_words":      "many-many-words",
		"AnyKind of_string":    "any-kind-of-string",
		"numbers2And55with000": "numbers-2-and-55-with-000",
		"JSONData":             "json-data",
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
		"test_case":            "TestCase",
		"test.case":            "TestCase",
		"test":                 "Test",
		"TestCase":             "TestCase",
		" test  case ":         "TestCase",
		"":                     "",
		"many_many_words":      "ManyManyWords",
		"AnyKind of_string":    "AnyKindOfString",
		"odd-fix":              "OddFix",
		"numbers2And55with000": "Numbers2And55With000",
	}
	runCaseChecks(cases, tb, ToPascalCase)
}

func toLowerCamelCase(tb testing.TB) {
	cases := tests.Cases{
		"foo-bar":           "fooBar",
		"TestCase":          "testCase",
		"":                  "",
		"AnyKind of_string": "anyKindOfString",
		"AnyKind.of-string": "anyKindOfString",
		"some string":       "someString",
		" some string":      "someString",
	}
	runCaseChecks(cases, tb, ToCamelCase)
}

func toSnakeCase(tb testing.TB) {
	cases := tests.Cases{
		"testCase":             "test_case",
		"TestCase":             "test_case",
		"Test Case":            "test_case",
		" Test Case":           "test_case",
		"Test Case ":           "test_case",
		" Test  Case ":         "test_case",
		"test":                 "test",
		"test_case":            "test_case",
		"Test":                 "test",
		"":                     "",
		"ManyManyWords":        "many_many_words",
		"manyManyWords":        "many_many_words",
		"AnyKind of_string":    "any_kind_of_string",
		"numbers2and55with000": "numbers_2_and_55_with_000",
		"JSONData":             "json_data",
		"userID":               "user_id",
		"AAAbbb":               "aa_abbb",
		"1A2":                  "1_a_2",
		"A1B":                  "a_1_b",
		"A1A2A3":               "a_1_a_2_a_3",
		"A1 A2 A3":             "a_1_a_2_a_3",
		"AB1AB2AB3":            "ab_1_ab_2_ab_3",
		"AB1 AB2 AB3":          "ab_1_ab_2_ab_3",
		"some string":          "some_string",
		" some string":         "some_string",
	}
	runCaseChecks(cases, tb, ToSnakeCase)
}

func toKebabCase(tb testing.TB) {
	cases := tests.Cases{
		"testCase":             "test-case",
		"TestCase":             "test-case",
		"Test Case":            "test-case",
		" Test Case":           "test-case",
		"Test Case ":           "test-case",
		" Test Case ":          "test-case",
		"test":                 "test",
		"test_case":            "test-case",
		"Test":                 "test",
		"":                     "",
		"ManyManyWords":        "many-many-words",
		"manyManyWords":        "many-many-words",
		"AnyKind of_string":    "any-kind-of-string",
		"numbers2and55with000": "numbers-2-and-55-with-000",
		"JSONData":             "json-data",
		"userID":               "user-id",
		"AAAbbb":               "aa-abbb",
		"1A2":                  "1-a-2",
		"A1B":                  "a-1-b",
		"A1A2A3":               "a-1-a-2-a-3",
		"A1 A2 A3":             "a-1-a-2-a-3",
		"AB1AB2AB3":            "ab-1-ab-2-ab-3",
		"AB1 AB2 AB3":          "ab-1-ab-2-ab-3",
		"some string":          "some-string",
		" some string":         "some-string",
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
