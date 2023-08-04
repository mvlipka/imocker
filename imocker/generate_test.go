package imocker

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestParseFile(t *testing.T) {

	f, err := os.Open("../testdata/example.go")
	require.NoError(t, err)
	defer f.Close()

	mocks, err := ParseMock(f)
	assert.NoError(t, err)
	require.Len(t, mocks, 1)

	require.Len(t, mocks[0].Methods, 1)
	require.Contains(t, mocks[0].Methods, "MyFunc")

	myFunc := mocks[0].Methods["MyFunc"]
	assert.Len(t, myFunc.NamedParams, 1)
	require.Contains(t, myFunc.NamedParams, "param")
	assert.Contains(t, myFunc.NamedParams["param"], "string")

	assert.Len(t, myFunc.UnNamedReturns, 1)
	require.Len(t, myFunc.UnNamedReturns, 1)
	assert.Contains(t, myFunc.UnNamedReturns[0], "error")
}

func TestGenerateTemplate(t *testing.T) {
	output, err := GenerateTemplate(Mock{
		Package: "testmock",
		Name:    "TestStruct",
		Methods: map[string]Method{
			"TestMethod": {
				NamedParams: map[string]string{
					"test": "string",
				},
				NamedReturns:  nil,
				UnNamedParams: nil,
				UnNamedReturns: []string{
					"error",
				},
			},
		},
	})

	require.NoError(t, err)
	assert.Equal(t, "\npackage testmock\n\ntype MockTestStruct struct {\n\t\n\tTestTestMethod func(test string) (error)\n\t\n}\n\n\nfunc (m *MockTestStruct) TestMethod(test string) (error) {\n\treturn m.TestTestMethod(test)\n}\n\n", output)
}
