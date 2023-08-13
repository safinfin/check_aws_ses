package sendquota

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/mackerelio/checkers"
	"github.com/safinfin/check_aws_ses/client"
	"github.com/stretchr/testify/assert"
)

func TestCheckSendQuota(t *testing.T) {
	ctx := context.TODO()
	cases := []struct {
		caseName        string
		checkStatus     checkers.Status
		checkString     string
		max24HourSend   float64
		sentLast24Hours float64
		options         SendQuotaOptions
	}{
		{
			caseName:        "OK case1",
			checkStatus:     checkers.OK,
			checkString:     "send quota usage is 50% (SentLast24Hours: 100, Max24HourSend: 200)",
			max24HourSend:   200,
			sentLast24Hours: 100,
			options: SendQuotaOptions{
				Warning:  60,
				Critical: 80,
			},
		},
		{
			caseName:        "OK case2",
			checkStatus:     checkers.OK,
			checkString:     "send quota usage is 50% (SentLast24Hours: 100, Max24HourSend: 200)",
			max24HourSend:   200,
			sentLast24Hours: 100,
			options: SendQuotaOptions{
				Warning:  51,
				Critical: 0,
			},
		},
		{
			caseName:        "OK case3",
			checkStatus:     checkers.OK,
			checkString:     "send quota usage is 50% (SentLast24Hours: 100, Max24HourSend: 200)",
			max24HourSend:   200,
			sentLast24Hours: 100,
			options: SendQuotaOptions{
				Warning:  0,
				Critical: 51,
			},
		},
		{
			caseName:        "OK case4",
			checkStatus:     checkers.OK,
			checkString:     "send quota usage is 50% (SentLast24Hours: 100, Max24HourSend: 200)",
			max24HourSend:   200,
			sentLast24Hours: 100,
			options: SendQuotaOptions{
				Warning:  51,
				Critical: 51,
			},
		},
		{
			caseName:        "OK case5",
			checkStatus:     checkers.OK,
			checkString:     "send quota is unlimited because Max24HourSend is unlimited (SentLast24Hours: 200, Max24HourSend: -1)",
			max24HourSend:   -1,
			sentLast24Hours: 200,
			options: SendQuotaOptions{
				Warning:  50,
				Critical: 60,
			},
		},
		{
			caseName:        "WARNING case1",
			checkStatus:     checkers.WARNING,
			checkString:     "send quota usage is 40% (SentLast24Hours: 20000, Max24HourSend: 50000)",
			max24HourSend:   50000,
			sentLast24Hours: 20000,
			options: SendQuotaOptions{
				Warning:  30,
				Critical: 50,
			},
		},
		{
			caseName:        "WARNING case2",
			checkStatus:     checkers.WARNING,
			checkString:     "send quota usage is 40% (SentLast24Hours: 20000, Max24HourSend: 50000)",
			max24HourSend:   50000,
			sentLast24Hours: 20000,
			options: SendQuotaOptions{
				Warning:  30,
				Critical: 0,
			},
		},
		{
			caseName:        "WARNING case3",
			checkStatus:     checkers.WARNING,
			checkString:     "send quota usage is 40% (SentLast24Hours: 20000, Max24HourSend: 50000)",
			max24HourSend:   50000,
			sentLast24Hours: 20000,
			options: SendQuotaOptions{
				Warning:  40,
				Critical: 0,
			},
		},
		{
			caseName:        "CRITICAL case1",
			checkStatus:     checkers.CRITICAL,
			checkString:     "send quota usage is 75% (SentLast24Hours: 60000, Max24HourSend: 80000)",
			max24HourSend:   80000,
			sentLast24Hours: 60000,
			options: SendQuotaOptions{
				Warning:  50,
				Critical: 70,
			},
		},
		{
			caseName:        "CRITICAL case2",
			checkStatus:     checkers.CRITICAL,
			checkString:     "send quota usage is 75% (SentLast24Hours: 60000, Max24HourSend: 80000)",
			max24HourSend:   80000,
			sentLast24Hours: 60000,
			options: SendQuotaOptions{
				Warning:  0,
				Critical: 70,
			},
		},
		{
			caseName:        "CRITICAL case3",
			checkStatus:     checkers.CRITICAL,
			checkString:     "send quota usage is 75% (SentLast24Hours: 60000, Max24HourSend: 80000)",
			max24HourSend:   80000,
			sentLast24Hours: 60000,
			options: SendQuotaOptions{
				Warning:  0,
				Critical: 75,
			},
		},
	}
	mocks := make([]*client.MockSESGetAccountAPI, len(cases))

	for i := range mocks {
		mocks[i] = &client.MockSESGetAccountAPI{
			Output: &sesv2.GetAccountOutput{
				SendQuota: &types.SendQuota{
					Max24HourSend:   cases[i].max24HourSend,
					SentLast24Hours: cases[i].sentLast24Hours,
				},
			},
		}
		output, _ := client.GetAccountInfo(ctx, mocks[i])
		checker := NewSendQuotaChecker(output, cases[i].options)
		assert.Equal(t, cases[i].checkStatus, checker.Check().Status, cases[i].caseName)
		assert.Equal(t, cases[i].checkString, checker.Check().Message, cases[i].caseName)
	}
}
