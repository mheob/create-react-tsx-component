package utils

func ConvertFileName(name string, usesKebabCase bool) string {
	if usesKebabCase {
		return ToKebabCase(name)
	}

	return ToPascalCase(name)
}
