// Package caseconv provides string case conversion utilities
package caseconv

import (
	"fmt"
	"strings"
	"unicode"
)

// CaseType represents the type of case conversion
type CaseType int

const (
	// Snake represents snake_case
	Snake CaseType = iota + 1
	// Kebab represents kebab-case
	Kebab
	// Camel represents camelCase
	Camel
	// Pascal represents PascalCase
	Pascal
)

// String returns the string representation of CaseType
func (c CaseType) String() string {
	switch c {
	case Snake:
		return "snake_case"
	case Kebab:
		return "kebab-case"
	case Camel:
		return "camelCase"
	case Pascal:
		return "PascalCase"
	default:
		return "unknown"
	}
}

type Converter interface {
	To(caseType CaseType) string
	ToSnake() string
	ToKebab() string
	ToCamel() string
	ToPascal() string
}

// converter provides a fluent interface for string case conversion.
type converter struct {
	value     string
	separator rune
	firstCap  bool
	wordCap   bool
}

var _ Converter = (*converter)(nil)

// NewConverter creates a new string converter
func NewConverter(s string) Converter {
	return &converter{
		value:     s,
		separator: 0,
		firstCap:  false,
		wordCap:   false,
	}
}

// convert performs the actual case conversion
func (c *converter) convert() string {
	s := strings.TrimSpace(c.value)
	if s == "" {
		return s
	}

	var words []string
	var currentWord strings.Builder
	runes := []rune(s)

	// Split into words
	for i := 0; i < len(runes); i++ {
		r := runes[i]

		// Handle separators
		if r == '_' || r == '-' || unicode.IsSpace(r) {
			if currentWord.Len() > 0 {
				words = append(words, currentWord.String())
				currentWord.Reset()
			}
			continue
		}

		// Check for acronym (sequence of uppercase letters followed by a lowercase letter)
		if unicode.IsUpper(r) {
			// Look ahead to check if this is part of an acronym
			acronymEnd := i
			for j := i + 1; j < len(runes); j++ {
				if !unicode.IsUpper(runes[j]) {
					break
				}
				acronymEnd = j
			}

			// If we found multiple uppercase letters and the next letter is lowercase,
			// treat the uppercase sequence as a word (except the last letter if it's followed by lowercase)
			if acronymEnd > i && acronymEnd+1 < len(runes) && unicode.IsLower(runes[acronymEnd+1]) {
				// Add previous word if exists
				if currentWord.Len() > 0 {
					words = append(words, currentWord.String())
					currentWord.Reset()
				}
				// Add the acronym as a word (excluding the last letter)
				words = append(words, string(runes[i:acronymEnd]))
				// Start new word with the last uppercase letter
				currentWord.WriteRune(runes[acronymEnd])
				i = acronymEnd
				continue
			}
		}

		// Handle regular word boundary
		if i > 0 && unicode.IsUpper(r) && !unicode.IsUpper(runes[i-1]) {
			if currentWord.Len() > 0 {
				words = append(words, currentWord.String())
				currentWord.Reset()
			}
		}

		currentWord.WriteRune(r)
	}

	if currentWord.Len() > 0 {
		words = append(words, currentWord.String())
	}

	// Process words
	for i, word := range words {
		if word == "" {
			continue
		}

		// Convert word to lowercase unless it's a single-letter word
		if len(word) > 1 || !c.wordCap {
			word = strings.ToLower(word)
		}

		// Apply capitalization rules
		if (i == 0 && c.firstCap) || (i > 0 && c.wordCap) {
			word = strings.ToUpper(word[:1]) + word[1:]
		}

		words[i] = word
	}

	// Join words with separator
	if c.separator != 0 {
		return strings.Join(words, string(c.separator))
	}
	return strings.Join(words, "")
}

// ToSnake converts the string to snake_case
func (c *converter) ToSnake() string {
	c.separator = '_'
	c.firstCap = false
	c.wordCap = false
	return c.convert()
}

// ToKebab converts the string to kebab-case
func (c *converter) ToKebab() string {
	c.separator = '-'
	c.firstCap = false
	c.wordCap = false
	return c.convert()
}

// ToCamel converts the string to camelCase
func (c *converter) ToCamel() string {
	c.separator = 0
	c.firstCap = false
	c.wordCap = true
	return c.convert()
}

// ToPascal converts the string to PascalCase
func (c *converter) ToPascal() string {
	c.separator = 0
	c.firstCap = true
	c.wordCap = true
	return c.convert()
}

// To converts the string to the specified case type
func (c *converter) To(caseType CaseType) string {
	switch caseType {
	case Snake:
		return c.ToSnake()
	case Kebab:
		return c.ToKebab()
	case Camel:
		return c.ToCamel()
	case Pascal:
		return c.ToPascal()
	default:
		panic(fmt.Sprintf("unsupported case type: %v", caseType))
	}
}

// String returns the current value of the converter
func (c *converter) String() string {
	return c.value
}

// Format converts a string to the specified case type
func Format(s string, caseType CaseType) string {
	return NewConverter(s).To(caseType)
}

// ToSnakeCase converts a string to snake_case
func ToSnakeCase(s string) string {
	return NewConverter(s).ToSnake()
}

// ToKebabCase converts a string to kebab-case
func ToKebabCase(s string) string {
	return NewConverter(s).ToKebab()
}

// ToCamelCase converts a string to camelCase
func ToCamelCase(s string) string {
	return NewConverter(s).ToCamel()
}

// ToPascalCase converts a string to PascalCase
func ToPascalCase(s string) string {
	return NewConverter(s).ToPascal()
}
