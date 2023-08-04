package commands

import (
	"fmt"
	"github.com/mvlipka/imocker/imocker"
	"github.com/spf13/cobra"
	"go/format"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

var GenerateCmd = &cobra.Command{
	Use:     "generate",
	Short:   "Generates mock structures for interfaces in a supplied directory",
	Example: "commands generate ./...",
	Run:     generateRun,
}

func generateRun(cmd *cobra.Command, args []string) {
	directory := "./..."
	if len(args) == 1 {
		directory = args[0]
	}

	// ./... would indicate current working directory
	if directory == "./..." {
		directory, _ = os.Getwd()
	}

	fmt.Println(fmt.Sprintf("Generating mocks for %s and subdirectories", directory))

	// Iterate files in every child directory compiling Go interfaces to Mocks
	err := filepath.WalkDir(directory, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Ignore non-Go source files and mock files
		if !strings.HasSuffix(path, ".go") || strings.HasPrefix(d.Name(), "mock_") {
			return nil
		}

		fmt.Println(fmt.Sprintf("Parsing %s", path))
		currentDir := filepath.Dir(path)

		f, err := os.Open(path)
		if err != nil {
			return err
		}

		mocks, err := imocker.ParseMock(f)
		if err != nil {
			return err
		}

		// Generate each mock
		for _, mock := range mocks {
			fmt.Println("Iterating mocks")
			output, err := imocker.GenerateTemplate(mock)
			if err != nil {
				return err
			}

			// Create the mock file
			fileName := strings.ToLower(fmt.Sprintf("%s\\mock_%s.go", currentDir, mock.Name))
			fmt.Println(fmt.Sprintf("Creating file %s", fileName))
			mockFile, err := os.Create(fileName)
			if err != nil {
				return err
			}

			//// gofmt the mock
			formattedOutput, err := format.Source([]byte(output))
			if err != nil {
				return err
			}

			// Write the mock
			_, err = mockFile.Write(formattedOutput)
			if err != nil {
				return err
			}
			mockFile.Close()
		}

		return nil
	})
	if err != nil {
		panic(err)
	}
}
