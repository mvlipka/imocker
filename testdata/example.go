package testdata

import "fmt"

type Thinger interface {
	MyFunc(param string) (string, error)
	NamedReturn(multiple bool, types bool) (err error)
	MethodWithNoReturns()
}

type Thing struct {
}

func (t *Thing) NamedReturn(multiple bool, types bool) (err error) {
	return nil
}

func (t *Thing) MyFunc(param string) (error, string) {
	return nil, ""
}

func (t *Thing) MethodWithNoReturns() {

}

func ExampleFunction(thinger Thinger, val1, val2 bool, str string) (string, error) {
	err := thinger.NamedReturn(val1, val2)
	if err != nil {
		return "", fmt.Errorf("error running NamedReturn: %w", err)
	}

	output, err := thinger.MyFunc(str)
	if err != nil {
		return "", fmt.Errorf("error running MyFunc: %w", err)
	}

	thinger.MethodWithNoReturns()

	return fmt.Sprintf("%s %t %t: %s", str, val1, val2, output), nil
}
