package commands

import (
	"github.com/spf13/cobra"
	"log"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of commands",
	Run:   versionRun,
}

func versionRun(cmd *cobra.Command, args []string) {
	log.Println("0.0.0a")
}
