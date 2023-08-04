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
	if len(args) == 0 {
		directory = args[0]
	}

	if directory == "./..." {
		directory, _ = os.Getwd()
	}

	fmt.Println(fmt.Sprintf("Navigating %s", directory))

	err := filepath.WalkDir(directory, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !strings.HasSuffix(path, ".go") {
			return nil
		}

		fmt.Println(fmt.Sprintf("Navigating %s", path))
		currentDir := filepath.Dir(path)

		f, err := os.Open(path)
		if err != nil {
			return err
		}

		mocks, err := imocker.ParseMock(f)
		if err != nil {
			return nil
		}

		fmt.Println("Mocks parsed")

		for _, mock := range mocks {
			fmt.Println("Iterating mocks")
			output, err := imocker.GenerateTemplate(mock)
			if err != nil {
				return err
			}

			fileName := strings.ToLower(fmt.Sprintf("%s\\mock_%s.go", currentDir, mock.Name))
			fmt.Println(fmt.Sprintf("Creating file %s", fileName))
			mockFile, err := os.Create(fileName)
			if err != nil {
				return err
			}

			formattedOutput, err := format.Source([]byte(output))
			if err != nil {
				return err
			}

			_, err = mockFile.Write(formattedOutput)
			mockFile.Close()
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return
	}
}
