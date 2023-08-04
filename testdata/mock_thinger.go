package testdata

type MockThinger struct {
	TestMyFunc func(param string) error
}

func (m *MockThinger) MyFunc(param string) error {
	return m.TestMyFunc(param)
}
