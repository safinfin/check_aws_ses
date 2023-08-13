package cmd

import (
	"os"

	"github.com/mackerelio/checkers"
	"github.com/safinfin/check_aws_ses/client"
	"github.com/safinfin/check_aws_ses/version"
	"github.com/spf13/cobra"
)

var (
	ClientOpts client.ClientOptions

	rootCmd = &cobra.Command{
		Use:     "check_aws_ses",
		Short:   "check_aws_ses is a nagios plugins to check reputation status or send quota usage of your Amazon SES account",
		Long:    `check_aws_ses is a nagios plugins to check reputation status or send quota usage of your Amazon SES account.`,
		Version: version.Version,
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(int(checkers.UNKNOWN))
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.PersistentFlags().SortFlags = false
	rootCmd.Flags().SortFlags = false
	rootCmd.PersistentFlags().StringVarP(&ClientOpts.Region, "region", "R", "", "AWS region (required)")
	rootCmd.PersistentFlags().StringVarP(&ClientOpts.Accesskey, "accesskey", "A", "", "AWS ACCESS KEY ID (required if secretkey is set)")
	rootCmd.PersistentFlags().StringVarP(&ClientOpts.Secretkey, "secretkey", "S", "", "AWS SECRET ACCESS KEY (required if accesskey is set)")
	rootCmd.MarkFlagsRequiredTogether("accesskey", "secretkey")
}
