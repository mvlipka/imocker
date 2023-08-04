package cmd

import (
	"fmt"
	"github.com/mvlipka/imocker/cmd/commands"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:     "imocker",
	Short:   "imocker is a Golang mock struct generator",
	Long:    "A mock struct generator and implementor for interfaces",
	Run:     rootRun,
	Example: "imocker generate ./...",
}

func init() {
	rootCmd.AddCommand(commands.GenerateCmd)
	rootCmd.AddCommand(commands.VersionCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}

func rootRun(cmd *cobra.Command, args []string) {
	if err := cmd.Help(); err != nil {
		os.Exit(1)
	}
}
