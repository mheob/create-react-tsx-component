package utils

import "testing"

func TestToSnakeCase(t *testing.T) {
	toSnakeCase(t)
}

func BenchmarkToSnakeCase(b *testing.B) {
	BenchmarkTest(b, toSnakeCase)
}

func TestToKebabCase(t *testing.T) {
	toKebabCase(t)
}

func BenchmarkToKebabCase(b *testing.B) {
	BenchmarkTest(b, toKebabCase)
}

func toSnakeCase(tb testing.TB) {
	cases := [][]string{
		{"testCase", "test_case"},
		{"TestCase", "test_case"},
		{"Test Case", "test_case"},
		{" Test Case", "test_case"},
		{"Test Case ", "test_case"},
		{" Test  Case ", "test_case"},
		{"test", "test"},
		{"test_case", "test_case"},
		{"Test", "test"},
		{"", ""},
		{"ManyManyWords", "many_many_words"},
		{"manyManyWords", "many_many_words"},
		{"AnyKind of_string", "any_kind_of_string"},
		{"numbers2and55with000", "numbers_2_and_55_with_000"},
		{"JSONData", "json_data"},
		{"userID", "user_id"},
		{"AAAbbb", "aa_abbb"},
		{"1A2", "1_a_2"},
		{"A1B", "a_1_b"},
		{"A1A2A3", "a_1_a_2_a_3"},
		{"A1 A2 A3", "a_1_a_2_a_3"},
		{"AB1AB2AB3", "ab_1_ab_2_ab_3"},
		{"AB1 AB2 AB3", "ab_1_ab_2_ab_3"},
		{"some string", "some_string"},
		{" some string", "some_string"},
	}
	RunChecks(cases, tb, ToSnakeCase)
}

func toKebabCase(tb testing.TB) {
	cases := [][]string{
		{"testCase", "test-case"},
		{"TestCase", "test-case"},
		{"Test Case", "test-case"},
		{" Test Case", "test-case"},
		{"Test Case ", "test-case"},
		{" Test Case ", "test-case"},
		{"test", "test"},
		{"test_case", "test-case"},
		{"Test", "test"},
		{"", ""},
		{"ManyManyWords", "many-many-words"},
		{"manyManyWords", "many-many-words"},
		{"AnyKind of_string", "any-kind-of-string"},
		{"numbers2and55with000", "numbers-2-and-55-with-000"},
		{"JSONData", "json-data"},
		{"userID", "user-id"},
		{"AAAbbb", "aa-abbb"},
		{"1A2", "1-a-2"},
		{"A1B", "a-1-b"},
		{"A1A2A3", "a-1-a-2-a-3"},
		{"A1 A2 A3", "a-1-a-2-a-3"},
		{"AB1AB2AB3", "ab-1-ab-2-ab-3"},
		{"AB1 AB2 AB3", "ab-1-ab-2-ab-3"},
		{"some string", "some-string"},
		{" some string", "some-string"},
	}
	RunChecks(cases, tb, ToKebabCase)
}
