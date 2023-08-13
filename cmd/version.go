package cmd

import (
	"fmt"

	"github.com/safinfin/check_aws_ses/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version information",
	Long:  `Print the version information.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s version %s\n", version.Name, version.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
