package mockx_test

import (
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
