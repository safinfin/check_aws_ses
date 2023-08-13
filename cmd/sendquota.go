package cmd

import (
	"context"
	"errors"

	"github.com/mackerelio/checkers"
	"github.com/safinfin/check_aws_ses/client"
	"github.com/safinfin/check_aws_ses/sendquota"
	"github.com/spf13/cobra"
)

var (
	sendQuotaOpts sendquota.SendQuotaOptions

	sendquotaCmd = &cobra.Command{
		Use:   "sendquota",
		Short: "Check the send quota usage of your Amazon SES account in the current region",
		Long: `Check the send quota usage of your Amazon SES account in the current region.

send quota is an object that contains information about the per-day and per-second sending
limits for your Amazon SES account in the current AWS Region.

- SentLast24Hours
  The number of emails sent from your Amazon SES account in the current region over the past 24 hours.
- Max24HourSend
  The maximum number of emails that you can send in the current region over a 24-hour period.
  A value of "-1" signifies an unlimited quota.
- SendQuotaUsage (%)
  SentLast24Hours / Max24HourSend * 100`,
		Args: func(cmd *cobra.Command, args []string) error {
			if ClientOpts.Region == "" {
				return errors.New("required flag \"--region (-R)\" not set")
			}
			if sendQuotaOpts.Warning == 0 && sendQuotaOpts.Critical == 0 {
				return errors.New("at least one of either warning or critical is required")
			}
			if sendQuotaOpts.Critical != 0 && sendQuotaOpts.Warning > sendQuotaOpts.Critical {
				return errors.New("critical must be greater than or equal to warning")
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
			s := sendquota.NewSendQuotaChecker(output, sendQuotaOpts)
			s.Check().Exit()
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(sendquotaCmd)
	sendquotaCmd.PersistentFlags().SortFlags = false
	sendquotaCmd.Flags().SortFlags = false
	sendquotaCmd.Flags().Float64VarP(&sendQuotaOpts.Warning, "warning", "w", 0, "send quota usage result in warning status (%)")
	sendquotaCmd.Flags().Float64VarP(&sendQuotaOpts.Critical, "critical", "c", 0, "send quota usage result in critical status (%)")
}
