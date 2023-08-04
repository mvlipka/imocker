package testdata

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestExampleFunction(t *testing.T) {
	thinger := &MockThinger{
		TestMyFunc: func(param string) (string, error) {
			assert.Equal(t, "testing", param)
			return "example", nil
		},
		TestNamedReturn: func(multiple, types bool) (err error) {
			assert.True(t, multiple)
			assert.False(t, types)
			return nil
		},
	}

	output, err := ExampleFunction(thinger, true, false, "testing")
	require.NoError(t, err)
	assert.Equal(t, "testing true false: example", output)
}

func TestExampleFunctionNamedReturnError(t *testing.T) {
	thinger := &MockThinger{
		TestMyFunc: func(param string) (string, error) {
			// MyFunc is called after NamedReturn, if then NamedReturn did not return an error
			t.Failed()
			return "", nil
		},
		TestNamedReturn: func(multiple, types bool) (err error) {
			assert.True(t, multiple)
			assert.False(t, types)
			return errors.New("an error")
		},
	}

	_, err := ExampleFunction(thinger, true, false, "testing")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "an error")
}

func TestExampleFunctionMyFuncError(t *testing.T) {
	thinger := &MockThinger{
		TestMyFunc: func(param string) (string, error) {
			assert.Equal(t, "testing", param)
			return "", errors.New("an error")
		},
		TestNamedReturn: func(multiple, types bool) (err error) {
			assert.True(t, multiple)
			assert.False(t, types)
			return nil
		},
	}

	_, err := ExampleFunction(thinger, true, false, "testing")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "an error")
}
