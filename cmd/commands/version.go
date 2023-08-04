package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of commands",
	Run:   versionRun,
}

func versionRun(cmd *cobra.Command, args []string) {
	fmt.Println("0.0.0a")
}
