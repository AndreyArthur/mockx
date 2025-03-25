# Mockx

Mockx is a lightweight and intuitive mocking library for Go interfaces. It simplifies testing by allowing you to create mock implementations, define method behaviors, and capture method arguments with minimal boilerplate.

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
    return values[0].(int)
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

## Advanced Example

Test a `Searcher` that uses a `Sorting` interface dependency:

```go
// Define the interface
type Sorting interface {
    IsSorted(slice []int) bool
}

// Create a mock
type SortingMock struct {
    mockx.Mockx
}

func (m *SortingMock) IsSorted(slice []int) bool {
    values := m.Call("IsSorted", slice)
    return values[0].(bool)
}

// Usage in tests
func TestSearcher(t *testing.T) {
    sorting := &SortingMock{}
    sorting.Init((*Sorting)(nil))
    searcher := NewSearcher(sorting)

    // Force IsSorted to return false (use linear search)
    sorting.Return("IsSorted", false)
    index := searcher.Search([]int{3,1,2}, 3) // Returns 0

    // Override IsSorted with custom logic
    sorting.Impl("IsSorted", func(slice []int) bool {
        // Check if sorted
        for i := 0; i < len(slice)-1; i++ {
            if slice[i] > slice[i+1] {
                return false
            }
        }
        return true
    })
    index = searcher.Search([]int{1,2,3,4}, 4) // Returns 3 (binary search)
}
```

## License

MIT License. See [LICENSE](LICENSE) for details.
