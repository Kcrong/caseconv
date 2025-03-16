# caseconv

[![Lint and Test](https://github.com/Kcrong/caseconv/actions/workflows/lint.yml/badge.svg)](https://github.com/Kcrong/caseconv/actions/workflows/lint.yml)

`caseconv` is a Go package for string case conversions. It supports conversions between snake_case, kebab-case, camelCase, and PascalCase.

## Installation

```bash
go get github.com/Kcrong/caseconv
```

## Usage

### Method Chaining

```go
import "github.com/Kcrong/caseconv"

str := "hello_world"

// snake_case
result := caseconv.NewConverter(str).ToSnake()  // "hello_world"

// kebab-case
result := caseconv.NewConverter(str).ToKebab()  // "hello-world"

// camelCase
result := caseconv.NewConverter(str).ToCamel()  // "helloWorld"

// PascalCase
result := caseconv.NewConverter(str).ToPascal() // "HelloWorld"
```

### Package-Level Functions

```go
import "github.com/Kcrong/caseconv"

str := "hello_world"

// snake_case
result := caseconv.ToSnakeCase(str)  // "hello_world"

// kebab-case
result := caseconv.ToKebabCase(str)  // "hello-world"

// camelCase
result := caseconv.ToCamelCase(str)  // "helloWorld"

// PascalCase
result := caseconv.ToPascalCase(str) // "HelloWorld"
```

### Using CaseType

```go
import "github.com/Kcrong/caseconv"

str := "hello_world"

// Method chaining with CaseType
result := caseconv.NewConverter(str).To(caseconv.Snake)   // "hello_world"
result := caseconv.NewConverter(str).To(caseconv.Kebab)   // "hello-world"
result := caseconv.NewConverter(str).To(caseconv.Camel)   // "helloWorld"
result := caseconv.NewConverter(str).To(caseconv.Pascal)  // "HelloWorld"

// Using Format function
result := caseconv.Format(str, caseconv.Snake)   // "hello_world"
result := caseconv.Format(str, caseconv.Kebab)   // "hello-world"
result := caseconv.Format(str, caseconv.Camel)   // "helloWorld"
result := caseconv.Format(str, caseconv.Pascal)  // "HelloWorld"
```

## Supported Case Types

- `snake_case`: Words are separated by underscores (_) and all characters are lowercase
- `kebab-case`: Words are separated by hyphens (-) and all characters are lowercase
- `camelCase`: First word starts with lowercase, subsequent words start with uppercase
- `PascalCase`: All words start with uppercase

## Features

- Intuitive API with method chaining
- Support for various input formats (snake_case, kebab-case, camelCase, PascalCase)
- Automatic whitespace handling
- Handling of consecutive separators
- Proper handling of consecutive uppercase letters

## Examples

```go
import "github.com/Kcrong/caseconv"

// Various input format handling
caseconv.NewConverter("hello-world").ToCamel()     // "helloWorld"
caseconv.NewConverter("hello_world").ToPascal()    // "HelloWorld"
caseconv.NewConverter("helloWorld").ToSnake()      // "hello_world"
caseconv.NewConverter("HelloWorld").ToKebab()      // "hello-world"

// Consecutive uppercase handling
caseconv.NewConverter("HTTPRequest").ToCamel()     // "httpRequest"
caseconv.NewConverter("HTTPRequest").ToSnake()     // "http_request"

// Consecutive separator handling
caseconv.NewConverter("hello__world").ToSnake()    // "hello_world"
caseconv.NewConverter("hello--world").ToKebab()    // "hello-world"

// Whitespace handling
caseconv.NewConverter("hello world").ToCamel()     // "helloWorld"
caseconv.NewConverter("hello world").ToSnake()     // "hello_world"
```

## License

MIT License 