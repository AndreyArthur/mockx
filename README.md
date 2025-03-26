# Mockx

Mockx is a lightweight and intuitive mocking library for Go interfaces. It simplifies testing by allowing you to create mock implementations, define method behaviors, and capture method arguments with minimal boilerplate.

- [Motivation](#motivation)
- [Installation](#installation)
- [Docs](https://pkg.go.dev/github.com/AndreyArthur/mockx)
- [Features](#features)
- [Usage](#usage)
- [Examples](#examples)
- [License](#license)

## Motivation

This section reflects only the author's opinion. Whether you agree or disagree, you are encouraged to continue using the library as you see fit.

- Dissatisfaction with the excessive boilerplate required to create mocks with other libraries, which practically forces the use of code generation.
- Code generation feels like a hack rather than a proper solution in any sense.
- Lack of clarity and directness when defining mock behaviors, especially due to the use of expects.
- Using expects in mocks is conceptually wrong. If you find expects useful in your mocks, you are probably not writing unit tests. Additionally, integration tests should not use mocks.
- Using a bunch of "anything" as expects feels even worse.

## Installation

```bash
go get github.com/AndreyArthur/mockx
```

## Features

- **Automatic Method Initialization**: Generate zero-value implementations for all interface methods.
- **Method Behavior Mocking**: Override methods with custom implementations or predefined return values.
- **Argument Capture**: Easily retrieve arguments passed to mocked methods during tests.
- **Lightweight**: No external dependencies and minimal setup required.
- **Example-Driven**: Comprehensive examples included in tests for quick learning.

## Usage

### 1. Define Your Interface

```go
type Calculator interface {
    Add(a int, b int) int
}
```

### 2. Create a Mock Struct

Embed `mockx.Mockx` and implement the interface methods using `Call`:

```go
type CalculatorMock struct {
    mockx.Mockx
}

func (m *CalculatorMock) Add(a int, b int) int {
    values := m.Call("Add", a, b)
    return mockx.Value[int](values[0])
}
```

### 3. Initialize the Mock

Use `Init` to auto-generate zero-value implementations for all methods:

```go
calculator := &CalculatorMock{}
calculator.Init((*Calculator)(nil)) // Pass a nil interface pointer
```

### 4. Mock Method Behaviors

#### Set a Custom Implementation

```go
calculator.Impl("Add", func(a int, b int) int {
    return a + b
})
```

#### Define Return Values

```go
calculator.Return("Add", 42) // Always returns 42
```

#### Capture Arguments

```go
calculator.Add(1, 2)
args := calculator.Args("Add") // Returns [1, 2]
```

## Examples

### Basic Mock Setup

```go
// Initialize mock and use default zero values
calculator := &CalculatorMock{}
calculator.Init((*Calculator)(nil))

result := calculator.Add(3, 4) // Returns 0 (default)
```

### Override Method Implementation

```go
calculator.Impl("Add", func(a, b int) int {
    return a * b // Change behavior to multiply
})

result := calculator.Add(3, 4) // Returns 12
```

### Force Specific Return Value

```go
calculator.Return("Add", 100)

result := calculator.Add(5, 5) // Returns 100, ignores inputs
```

### Retrieve Method Arguments

```go
calculator.Add(10, 20)
args := calculator.Args("Add") // [10, 20]
```

## License

MIT License. See [LICENSE](LICENSE) for details.
