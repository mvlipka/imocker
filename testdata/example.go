// go:test
package testdata

type Thinger interface {
	MyFunc(param string) error
}

type Thing struct {
}

func (t *Thing) MyFunc(param string) error {
	return nil
}
