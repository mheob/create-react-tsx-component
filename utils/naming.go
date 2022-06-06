package utils

func ConvertFileName(name string, useKebabCase bool) string {
	if useKebabCase {
		return ToKebabCase(name)
	}

	return ToPascalCase(name)
}
