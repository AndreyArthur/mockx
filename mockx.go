// Package mockx provides a lightweight mocking utility for Go interfaces.
// It allows you to create mock implementations, register method behaviors and
// retrieve method arguments.
package mockx

import (
	"fmt"
	"reflect"
	"slices"
)

// Mockx is the main struct that manages mock method implementations and args
// tracking.
type Mockx struct {
	methods map[string]reflect.Value
	args    map[string][]any
}

// Init populates the mock instance with zero-value implementations for all
// methods of the given interface.
//
// It's not necessary to call Init into a Mockx instance, but if you don't, you
// need to manually call Impl or Return for all used methods, otherwise, your
// program will panic.
func (mockxInstance *Mockx) Init(nilInterface any) {
	interfaceType := reflect.TypeOf(nilInterface).Elem()

	for i := range interfaceType.NumMethod() {
		methodValue := interfaceType.Method(i)
		methodType := methodValue.Type

		returnValues := make([]reflect.Value, methodType.NumOut())
		for i := range methodType.NumOut() {
			returnValues[i] = reflect.Zero(methodType.Out(i))
		}

		funcValue := reflect.MakeFunc(methodType, func(args []reflect.Value) []reflect.Value {
			return returnValues
		})

		funcAsAny := funcValue.Interface()

		mockxInstance.Impl(methodValue.Name, funcAsAny)
	}
}

// Call invokes a registered mock method with the given arguments. Panics if the
// method is not registered.
func (mockxInstance *Mockx) Call(method string, args ...any) []any {
	funcValue, ok := mockxInstance.methods[method]
	if !ok {
		panic(fmt.Sprintf("Could not call method %q, not registered in mockx instance.", method))
	}

	if mockxInstance.args == nil {
		mockxInstance.args = make(map[string][]any)
	}
	mockxInstance.args[method] = slices.Clone(args)

	reflectionArgs := make([]reflect.Value, len(args))
	for i, value := range args {
		reflectionArgs[i] = reflect.ValueOf(value)
	}

	reflectionReturnValues := funcValue.Call(reflectionArgs)

	returnValues := make([]any, len(reflectionReturnValues))
	for i, value := range reflectionReturnValues {
		returnValues[i] = value.Interface()
	}

	return returnValues
}

// Impl manually registers a function as an implementation for a method. The
// provided function must match the signature of the method.
func (mockxInstance *Mockx) Impl(method string, fn any) {
	if mockxInstance.methods == nil {
		mockxInstance.methods = make(map[string]reflect.Value)
	}
	mockxInstance.methods[method] = reflect.ValueOf(fn)
}

// Return registers a mock method to return the given values.
func (mockxInstance *Mockx) Return(method string, values ...any) {
	if mockxInstance.methods == nil {
		mockxInstance.methods = make(map[string]reflect.Value)
	}

	registeredFuncValue, ok := mockxInstance.methods[method]
	if !ok {
		panic(fmt.Sprintf("Could not call method %q, not registered in mockx instance.", method))
	}
	funcType := registeredFuncValue.Type()

	returnValues := make([]reflect.Value, len(values))
	for i := range funcType.NumOut() {
		valueType := funcType.Out(i)
		if values[i] == nil || (reflect.ValueOf(values[i]).Kind() == reflect.Ptr && reflect.ValueOf(values[i]).IsNil()) {
			returnValues[i] = reflect.Zero(valueType)
		} else {
			returnValues[i] = reflect.ValueOf(values[i])
		}
	}

	funcValue := reflect.MakeFunc(funcType, func(args []reflect.Value) []reflect.Value {
		return returnValues
	})

	mockxInstance.methods[method] = funcValue
}

// Args retrieves the arguments used in the most recent call to the specified
// method. Panics if the method was never called.
func (mockInstance *Mockx) Args(method string) []any {
	if mockInstance.args == nil {
		panic(fmt.Sprintf("Cannot get args for method %q, method was not called.", method))
	}

	args, ok := mockInstance.args[method]
	if !ok {
		panic(fmt.Sprintf("Cannot get args for method %q, method was not called.", method))
	}

	return args
}
