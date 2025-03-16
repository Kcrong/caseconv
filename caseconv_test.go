package caseconv

import "testing"

func TestToSnakeCase(t *testing.T) {
	t.Parallel()

	cases := map[string]string{
		"helloWorld":  "hello_world",
		"HelloWorld":  "hello_world",
		"hello":       "hello",
		"HELLO":       "hello",
		"hello_world": "hello_world",
		"hello-world": "hello_world",
		"HELLO_WORLD": "hello_world",
	}

	for input, expected := range cases {
		input, expected := input, expected
		t.Run(input, func(t *testing.T) {
			t.Parallel()
			result := ToSnakeCase(input)
			if result != expected {
				t.Errorf("ToSnakeCase(%q) = %q, expected %q", input, result, expected)
			}
		})
	}
}

func TestToKebabCase(t *testing.T) {
	t.Parallel()

	cases := map[string]string{
		"helloWorld":  "hello-world",
		"HelloWorld":  "hello-world",
		"hello":       "hello",
		"HELLO":       "hello",
		"hello_world": "hello-world",
		"hello-world": "hello-world",
		"HELLO_WORLD": "hello-world",
	}

	for input, expected := range cases {
		input, expected := input, expected
		t.Run(input, func(t *testing.T) {
			t.Parallel()
			result := ToKebabCase(input)
			if result != expected {
				t.Errorf("ToKebabCase(%q) = %q, expected %q", input, result, expected)
			}
		})
	}
}

func TestToCamelCase(t *testing.T) {
	t.Parallel()

	cases := map[string]string{
		"helloWorld":  "helloWorld",
		"HelloWorld":  "helloWorld",
		"hello":       "hello",
		"HELLO":       "hello",
		"hello_world": "helloWorld",
		"hello-world": "helloWorld",
		"HELLO_WORLD": "helloWorld",
	}

	for input, expected := range cases {
		input, expected := input, expected
		t.Run(input, func(t *testing.T) {
			t.Parallel()
			result := ToCamelCase(input)
			if result != expected {
				t.Errorf("ToCamelCase(%q) = %q, expected %q", input, result, expected)
			}
		})
	}
}

func TestToPascalCase(t *testing.T) {
	t.Parallel()

	cases := map[string]string{
		"helloWorld":  "HelloWorld",
		"HelloWorld":  "HelloWorld",
		"hello":       "Hello",
		"HELLO":       "Hello",
		"hello_world": "HelloWorld",
		"hello-world": "HelloWorld",
		"HELLO_WORLD": "HelloWorld",
	}

	for input, expected := range cases {
		input, expected := input, expected
		t.Run(input, func(t *testing.T) {
			t.Parallel()
			result := ToPascalCase(input)
			if result != expected {
				t.Errorf("ToPascalCase(%q) = %q, expected %q", input, result, expected)
			}
		})
	}
}
