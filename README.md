# caseconv

[![Tests](https://github.com/Kcrong/caseconv/actions/workflows/test.yml/badge.svg)](https://github.com/Kcrong/caseconv/actions/workflows/test.yml)

A Go package for string case style conversions.

## Features

- `ToSnakeCase`: Convert string to snake_case
- `ToKebabCase`: Convert string to kebab-case
- `ToCamelCase`: Convert string to camelCase
- `ToPascalCase`: Convert string to PascalCase

## Installation

```bash
go get github.com/Kcrong/caseconv
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/Kcrong/caseconv"
)

func main() {
    // Snake Case
    fmt.Println(caseconv.ToSnakeCase("helloWorld"))  // "hello_world"
    
    // Kebab Case
    fmt.Println(caseconv.ToKebabCase("helloWorld"))  // "hello-world"
    
    // Camel Case
    fmt.Println(caseconv.ToCamelCase("hello_world")) // "helloWorld"
    
    // Pascal Case
    fmt.Println(caseconv.ToPascalCase("hello_world")) // "HelloWorld"
}
```

## Testing

```bash
go test -v
```

## License

MIT License 