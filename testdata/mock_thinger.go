package testdata

type MockThinger struct {
	TestMyFunc func(param string) (string, error)

	TestNamedReturn func(multiple bool, types bool) (err error)
}

func (m *MockThinger) MyFunc(param string) (string, error) {
	return m.TestMyFunc(param)
}

func (m *MockThinger) NamedReturn(multiple bool, types bool) (err error) {
	return m.TestNamedReturn(multiple, types)
}
