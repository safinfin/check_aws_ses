package cmd

import (
	"context"
	"errors"

	"github.com/mackerelio/checkers"
	"github.com/safinfin/check_aws_ses/client"
	"github.com/safinfin/check_aws_ses/reputation"
	"github.com/spf13/cobra"
)

var (
	reputationCmd = &cobra.Command{
		Use:   "reputation",
		Short: "Check the reputation status of your Amazon SES account",
		Long: `Check the reputation status of your Amazon SES account.

The reputation status can be one of the following:
  - HEALTHY (check status is OK)
    There are no reputation-related issues that currently impact your account.
  - PROBATION (check status is WARNING)
    We've identified potential issues with your Amazon SES account.
    We're placing your account under review while you work on correcting these issues.
  - SHUTDOWN (check status is CRITICAL)
    Your account's ability to send email is currently paused because
    of an issue with the email sent from your account. When you correct the issue,
    you can contact us and request that your account's ability to send email is resumed.`,
		Args: func(cmd *cobra.Command, args []string) error {
			if ClientOpts.Region == "" {
				return errors.New("required flag \"--region (-R)\" not set")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.TODO()
			api, clErr := client.NewSESGetAccountAPI(ctx, ClientOpts)
			if clErr != nil {
				checkers.Unknown(clErr.Error()).Exit()
			}
			output, gaErr := client.GetAccountInfo(ctx, api)
			if gaErr != nil {
				checkers.Unknown(gaErr.Error()).Exit()
			}
			r := reputation.NewReputationChecker(output)
			r.Check().Exit()
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(reputationCmd)
	reputationCmd.PersistentFlags().SortFlags = false
	reputationCmd.Flags().SortFlags = false
}
