package client

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
)

type ClientOptions struct {
	Region    string
	Accesskey string
	Secretkey string
}

type SESGetAccountAPI interface {
	GetAccount(ctx context.Context, params *sesv2.GetAccountInput, optFns ...func(*sesv2.Options)) (*sesv2.GetAccountOutput, error)
}

func NewSESGetAccountAPI(ctx context.Context, co ClientOptions) (SESGetAccountAPI, error) {
	var cfg aws.Config
	var err error

	if co.Accesskey != "" && co.Secretkey != "" {
		cfg, err = config.LoadDefaultConfig(ctx, config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     co.Accesskey,
				SecretAccessKey: co.Secretkey,
				SessionToken:    "",
			},
		}), config.WithRegion(co.Region))
	} else {
		cfg, err = config.LoadDefaultConfig(ctx, config.WithRegion(co.Region))
	}

	if err != nil {
		return nil, err
	}

	return sesv2.NewFromConfig(cfg), nil
}

func GetAccountInfo(ctx context.Context, api SESGetAccountAPI) (*sesv2.GetAccountOutput, error) {
	input := sesv2.GetAccountInput{}
	output, err := api.GetAccount(ctx, &input)

	if err != nil {
		return nil, err
	}

	return output, nil
}
