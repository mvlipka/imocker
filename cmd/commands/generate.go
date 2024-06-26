package commands

import (
	"fmt"
	"github.com/mvlipka/imocker/imocker"
	"github.com/spf13/cobra"
	"go/format"
	"io/fs"
	"log"
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
	directory, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if len(args) == 1 {
		directory = args[0]
	}

	log.Println(fmt.Sprintf("Generating mocks for %s and subdirectories", directory))

	// Iterate files in every child directory compiling Go interfaces to Mocks
	err = filepath.WalkDir(directory, walkDirectoryFn)
	if err != nil {
		panic(err)
	}
}

func walkDirectoryFn(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return fmt.Errorf("error walking directory: %w", err)
	}

	// Ignore non-Go source files and mock files
	if !strings.HasSuffix(path, ".go") || strings.HasPrefix(d.Name(), "mock_") {
		return nil
	}

	log.Println(fmt.Sprintf("Parsing %s", path))
	currentDir := filepath.Dir(path)

	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}

	mocks, err := imocker.ParseMock(f)
	if err != nil {
		return fmt.Errorf("error parsing mock: %w", err)
	}

	// Generate each mock
	for _, mock := range mocks {
		log.Println("Iterating mocks")
		output, err := imocker.GenerateTemplate(mock)
		if err != nil {
			return fmt.Errorf("error generating template: %w", err)
		}

		// Create the mock file
		fileName := fmt.Sprintf("%s_mock.go", os.ExpandEnv(filepath.Join(currentDir, strings.ToLower(mock.Name))))
		log.Println(fmt.Sprintf("Creating file %s", fileName))
		mockFile, err := os.Create(fileName)
		if err != nil {
			_ = mockFile.Close()
			return fmt.Errorf("error creating file: %w", err)
		}

		// gofmt the mock
		formattedOutput, err := format.Source([]byte(output))
		if err != nil {
			_ = mockFile.Close()
			return fmt.Errorf("error formatting source: %w", err)
		}

		// Write the mock
		_, err = mockFile.Write(formattedOutput)
		if err != nil {
			_ = mockFile.Close()
			return fmt.Errorf("error writing to sthe mock file: %w", err)
		}

		_ = mockFile.Close()
	}

	return nil
}
