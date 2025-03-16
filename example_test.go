package caseconv_test

import (
	"fmt"

	"github.com/Kcrong/caseconv"
)

func ExampleNewConverter_reuse() {
	s := "hello_world"

	conv := caseconv.NewConverter(s)
	fmt.Println(conv.ToPascal())
	fmt.Println(conv.ToKebab())
	// Output:
	// HelloWorld
	// hello-world
}

func ExampleNewConverter_methodChaining() {
	s := "hello_world"
	fmt.Println(caseconv.NewConverter(s).ToPascal())
	// Output:
	// HelloWorld
}

func ExampleNewConverter_packageLevelFunctions() {
	s := "hello_world"
	snake := caseconv.ToSnakeCase(s)
	kebab := caseconv.ToKebabCase(s)
	camel := caseconv.ToCamelCase(s)
	pascal := caseconv.ToPascalCase(s)

	fmt.Println(snake)
	fmt.Println(kebab)
	fmt.Println(camel)
	fmt.Println(pascal)
	// Output:
	// hello_world
	// hello-world
	// helloWorld
	// HelloWorld
}

func ExampleNewConverter_usingCaseType() {
	s := "hello_world"
	snake := caseconv.Format(s, caseconv.Snake)
	kebab := caseconv.Format(s, caseconv.Kebab)
	camel := caseconv.Format(s, caseconv.Camel)
	pascal := caseconv.Format(s, caseconv.Pascal)

	fmt.Println(snake)
	fmt.Println(kebab)
	fmt.Println(camel)
	fmt.Println(pascal)
	// Output:
	// hello_world
	// hello-world
	// helloWorld
	// HelloWorld
}

func ExampleNewConverter_variousInputFormats() {
	inputs := []string{
		"hello_world",
		"hello-world",
		"helloWorld",
		"HelloWorld",
		"HELLO_WORLD",
	}

	for _, s := range inputs {
		fmt.Printf("Input: %s\n", s)
		fmt.Printf("Snake: %s\n", caseconv.ToSnakeCase(s))
		fmt.Printf("Kebab: %s\n", caseconv.ToKebabCase(s))
		fmt.Printf("Camel: %s\n", caseconv.ToCamelCase(s))
		fmt.Printf("Pascal: %s\n", caseconv.ToPascalCase(s))
		fmt.Println()
	}
	// Output:
	// Input: hello_world
	// Snake: hello_world
	// Kebab: hello-world
	// Camel: helloWorld
	// Pascal: HelloWorld
	//
	// Input: hello-world
	// Snake: hello_world
	// Kebab: hello-world
	// Camel: helloWorld
	// Pascal: HelloWorld
	//
	// Input: helloWorld
	// Snake: hello_world
	// Kebab: hello-world
	// Camel: helloWorld
	// Pascal: HelloWorld
	//
	// Input: HelloWorld
	// Snake: hello_world
	// Kebab: hello-world
	// Camel: helloWorld
	// Pascal: HelloWorld
	//
	// Input: HELLO_WORLD
	// Snake: hello_world
	// Kebab: hello-world
	// Camel: helloWorld
	// Pascal: HelloWorld
}

func ExampleNewConverter_consecutiveUppercase() {
	s := "HTTPRequest"
	camel := caseconv.ToCamelCase(s)
	snake := caseconv.ToSnakeCase(s)

	fmt.Println(camel)
	fmt.Println(snake)
	// Output:
	// httpRequest
	// http_request
}

func ExampleNewConverter_consecutiveSeparators() {
	s := "hello__world--test"
	snake := caseconv.ToSnakeCase(s)
	kebab := caseconv.ToKebabCase(s)

	fmt.Println(snake)
	fmt.Println(kebab)
	// Output:
	// hello_world_test
	// hello-world-test
}

func ExampleNewConverter_whitespaceHandling() {
	s := "  hello  world  "
	snake := caseconv.ToSnakeCase(s)
	kebab := caseconv.ToKebabCase(s)

	fmt.Println(snake)
	fmt.Println(kebab)
	// Output:
	// hello_world
	// hello-world
}
