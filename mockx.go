package mockx

import (
	"fmt"
	"reflect"
	"slices"
)

type Mockx struct {
	methods map[string]reflect.Value
	args    map[string][]any
}

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

func (mockxInstance *Mockx) Impl(method string, fn any) {
	if mockxInstance.methods == nil {
		mockxInstance.methods = make(map[string]reflect.Value)
	}
	mockxInstance.methods[method] = reflect.ValueOf(fn)
}

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
	for i, value := range values {
		returnValues[i] = reflect.ValueOf(value)
	}

	funcValue := reflect.MakeFunc(funcType, func(args []reflect.Value) []reflect.Value {
		return returnValues
	})

	mockxInstance.methods[method] = funcValue
}

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
