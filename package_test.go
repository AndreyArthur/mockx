package mockx_test

import (
	"fmt"

	"github.com/AndreyArthur/mockx"
)

type Greeter interface {
	Greet(name string) string
}

type GreeterMock struct {
	mockx.Mockx
}

func (greeter *GreeterMock) Greet(name string) string {
	values := greeter.Call("Greet", name)
	return mockx.Value[string](values[0])
}

func Example_mockx() {
	greeter := &GreeterMock{}
	greeter.Init((*Greeter)(nil))

	greeter.Impl("Greet", func(name string) string {
		return "Hello, " + name + "!"
	})

	result := greeter.Greet("Mockx")

	fmt.Printf("result = %q\n", result)
	// Output:
	// result = "Hello, Mockx!"
}
