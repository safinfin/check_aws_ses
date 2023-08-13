package reputation

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/mackerelio/checkers"
)

type ReputationChecker struct {
	Output *sesv2.GetAccountOutput
}

func NewReputationChecker(o *sesv2.GetAccountOutput) ReputationChecker {
	var r ReputationChecker = ReputationChecker{
		Output: o,
	}

	return r
}

func (r *ReputationChecker) Check() *checkers.Checker {
	reputationStatus := r.Output.EnforcementStatus
	checkStatus := checkers.UNKNOWN

	if *reputationStatus == "PROBATION" {
		checkStatus = checkers.WARNING
	}
	if *reputationStatus == "SHUTDOWN" {
		checkStatus = checkers.CRITICAL
	}
	if *reputationStatus == "HEALTHY" {
		checkStatus = checkers.OK
	}

	msg := fmt.Sprintf("reputation status is %s", *reputationStatus)

	return checkers.NewChecker(checkStatus, msg)
}
