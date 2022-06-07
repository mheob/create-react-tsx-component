package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	cases := Cases{
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
	cases := Cases{
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

func runFileNameChecks(cases Cases, tb testing.TB, usesKebabCase bool) {
	for in, expected := range cases {
		result := ConvertFileName(in, usesKebabCase)
		assert.EqualValues(tb, expected, result)
	}
}
