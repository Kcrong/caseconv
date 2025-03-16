// Package caseconv provides string case conversion utilities
package caseconv

import (
	"strings"
	"unicode"
)

// ToSnakeCase converts a string to snake_case
func ToSnakeCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if i > 0 && unicode.IsUpper(r) {
			result.WriteRune('_')
		}
		result.WriteRune(unicode.ToLower(r))
	}
	return result.String()
}

// ToKebabCase converts a string to kebab-case
func ToKebabCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if i > 0 && unicode.IsUpper(r) {
			result.WriteRune('-')
		}
		result.WriteRune(unicode.ToLower(r))
	}
	return result.String()
}

// ToCamelCase converts a string to camelCase
func ToCamelCase(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}

	var result strings.Builder
	capitalize := false
	for i, r := range s {
		if r == '_' || r == '-' || unicode.IsSpace(r) {
			capitalize = true
			continue
		}

		if i == 0 {
			result.WriteRune(unicode.ToLower(r))
		} else if capitalize {
			result.WriteRune(unicode.ToUpper(r))
			capitalize = false
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}

// ToPascalCase converts a string to PascalCase
func ToPascalCase(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}

	var result strings.Builder
	capitalize := true
	for _, r := range s {
		if r == '_' || r == '-' || unicode.IsSpace(r) {
			capitalize = true
			continue
		}

		if capitalize {
			result.WriteRune(unicode.ToUpper(r))
			capitalize = false
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}
