package client

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sesv2"
)

type MockSESGetAccountAPI struct {
	Output *sesv2.GetAccountOutput
	Error  error
}

func (m *MockSESGetAccountAPI) GetAccount(ctx context.Context, params *sesv2.GetAccountInput, optFns ...func(*sesv2.Options)) (*sesv2.GetAccountOutput, error) {
	return m.Output, m.Error
}
