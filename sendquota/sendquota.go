package sendquota

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/mackerelio/checkers"
)

type SendQuotaChecker struct {
	Output        *sesv2.GetAccountOutput
	SendQuotaOpts SendQuotaOptions
}

type SendQuotaOptions struct {
	Warning  float64
	Critical float64
}

func NewSendQuotaChecker(o *sesv2.GetAccountOutput, so SendQuotaOptions) SendQuotaChecker {
	var s SendQuotaChecker = SendQuotaChecker{
		Output:        o,
		SendQuotaOpts: so,
	}

	return s
}

func (s *SendQuotaChecker) Check() *checkers.Checker {
	max24HourSend := s.Output.SendQuota.Max24HourSend
	sentLast24Hours := s.Output.SendQuota.SentLast24Hours
	sendQuotaUsage := sentLast24Hours / max24HourSend * 100
	checkStatus := checkers.OK

	if max24HourSend == -1 {
		msg := fmt.Sprintf("send quota is unlimited because Max24HourSend is unlimited (SentLast24Hours: %g, Max24HourSend: %g)", sentLast24Hours, max24HourSend)
		return checkers.NewChecker(checkStatus, msg)
	}
	if s.SendQuotaOpts.Warning != 0 && s.SendQuotaOpts.Warning <= sendQuotaUsage {
		checkStatus = checkers.WARNING
	}
	if s.SendQuotaOpts.Critical != 0 && s.SendQuotaOpts.Critical <= sendQuotaUsage {
		checkStatus = checkers.CRITICAL
	}

	msg := fmt.Sprintf("send quota usage is %g%% (SentLast24Hours: %g, Max24HourSend: %g)", sendQuotaUsage, sentLast24Hours, max24HourSend)

	return checkers.NewChecker(checkStatus, msg)
}
