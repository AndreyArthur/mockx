package mockx_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/AndreyArthur/mockx"
)

type Calculator interface {
	Add(a int, b int) int
}

type CalculatorMock struct {
	mockx.Mockx
}

func (calculator *CalculatorMock) Add(a int, b int) int {
	values := calculator.Call("Add", a, b)
	return mockx.Value[int](values[0])
}

func TestReference(t *testing.T) {
	{
		var err error = errors.New("Error has occured.")
		var untyped any = err

		recovered := mockx.Reference[error](untyped)

		if recovered.Error() != "Error has occured." {
			t.Fatal("Unexpected error message.")
		}
	}

	{
		var err error = nil
		var untyped any = err

		recovered := mockx.Reference[error](untyped)

		if recovered != nil {
			t.Fatal("Expected recovered error to be nil.")
		}
	}
}

func TestValue(t *testing.T) {
	var value int = 64
	var untyped any = value

	recovered := mockx.Value[int](untyped)

	if recovered != 64 {
		t.Fatal("Expected recovered int to be 64.")
	}
}

func TestMockxInit(t *testing.T) {
	calculator := &CalculatorMock{}

	calculator.Init((*Calculator)(nil))

	result := calculator.Add(1, 2)

	if result != 0 {
		t.Fatal("Expected default value to be zero.")
	}
}

func TestMockxImpl(t *testing.T) {
	calculator := &CalculatorMock{}
	calculator.Init((*Calculator)(nil))

	calculator.Impl("Add", func(a int, b int) int {
		return a + b
	})

	result := calculator.Add(1, 2)

	if result != 3 {
		t.Fatal("Expected default value to be three.")
	}
}

func TestMockxReturn(t *testing.T) {
	calculator := &CalculatorMock{}
	calculator.Init((*Calculator)(nil))

	calculator.Return("Add", 64)

	result := calculator.Add(1, 2)

	if result != 64 {
		t.Fatal("Expected default value to be sixty four.")
	}
}

func TestMockxArgs(t *testing.T) {
	calculator := &CalculatorMock{}
	calculator.Init((*Calculator)(nil))

	calculator.Add(1, 2)
	args := calculator.Args("Add")

	if mockx.Value[int](args[0]) != 1 || mockx.Value[int](args[1]) != 2 {
		t.Fatal("Expected arguments to be saved.")
	}
}

func ExampleMockx_Init() {
	greeter := &GreeterMock{}
	greeter.Init((*Greeter)(nil))

	result := greeter.Greet("Mockx")

	fmt.Printf("result = %q\n", result)
	// Output:
	// result = ""
}

func ExampleMockx_Call() {
	greeter := &GreeterMock{}
	greeter.Init((*Greeter)(nil))

	values := greeter.Call("Greet", "Mockx")
	result := mockx.Value[string](values[0])

	fmt.Printf("result = %q\n", result)
	// Output:
	// result = ""
}

func ExampleMockx_Impl() {
	greeter := &GreeterMock{}
	greeter.Init((*Greeter)(nil))

	greeter.Impl("Greet", func(name string) string {
		return "Welcome, " + name + "."
	})

	result := greeter.Greet("Mockx")

	fmt.Printf("result = %q\n", result)
	// Output:
	// result = "Welcome, Mockx."
}

func ExampleMockx_Return() {
	greeter := &GreeterMock{}
	greeter.Init((*Greeter)(nil))

	greeter.Return("Greet", "Hello, Golang!")

	result := greeter.Greet("Mockx")

	fmt.Printf("result = %q\n", result)
	// Output:
	// result = "Hello, Golang!"
}

func ExampleMockx_Args() {
	greeter := &GreeterMock{}
	greeter.Init((*Greeter)(nil))

	greeter.Greet("Mockx")
	args := greeter.Args("Greet")
	arg := mockx.Value[string](args[0])

	fmt.Printf("arg = %q\n", arg)
	// Output:
	// arg = "Mockx"
}

func ExampleReference() {
	var untyped any = nil
	recovered := mockx.Reference[error](untyped)

	fmt.Printf("recovered = %v\n", recovered)

	// Output:
	// recovered = <nil>
}

func ExampleValue() {
	var untyped any = 42
	recovered := mockx.Value[int](untyped)

	fmt.Printf("recovered = %v\n", recovered)

	// Output:
	// recovered = 42
}
