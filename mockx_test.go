package mockx_test

import (
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
	return values[0].(int)
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

	if args[0].(int) != 1 || args[1].(int) != 2 {
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
	result := values[0].(string)

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
	arg := args[0].(string)

	fmt.Printf("arg = %q\n", arg)
	// Output:
	// arg = "Mockx"
}
