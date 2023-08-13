package reputation

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/mackerelio/checkers"
	"github.com/safinfin/check_aws_ses/client"
	"github.com/stretchr/testify/assert"
)

func TestCheckReputation(t *testing.T) {
	ctx := context.TODO()
	cases := []struct {
		caseName         string
		checkStatus      checkers.Status
		checkString      string
		reputationStatus string
	}{
		{
			caseName:         "case1 OK",
			checkStatus:      checkers.OK,
			checkString:      "reputation status is HEALTHY",
			reputationStatus: "HEALTHY",
		},
		{
			caseName:         "case2 WARNING",
			checkStatus:      checkers.WARNING,
			checkString:      "reputation status is PROBATION",
			reputationStatus: "PROBATION",
		},
		{
			caseName:         "case3 CRITICAL",
			checkStatus:      checkers.CRITICAL,
			checkString:      "reputation status is SHUTDOWN",
			reputationStatus: "SHUTDOWN",
		},
		{
			caseName:         "case4 UNKNOWN",
			checkStatus:      checkers.UNKNOWN,
			checkString:      "reputation status is NOTHING",
			reputationStatus: "NOTHING",
		},
	}
	mocks := make([]*client.MockSESGetAccountAPI, len(cases))

	for i := range mocks {
		mocks[i] = &client.MockSESGetAccountAPI{
			Output: &sesv2.GetAccountOutput{
				EnforcementStatus: &cases[i].reputationStatus,
			},
		}
		output, _ := client.GetAccountInfo(ctx, mocks[i])
		checker := NewReputationChecker(output)
		assert.Equal(t, cases[i].checkStatus, checker.Check().Status, cases[i].caseName)
		assert.Equal(t, cases[i].checkString, checker.Check().Message, cases[i].caseName)
	}
}
