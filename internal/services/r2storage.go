package services

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

type R2Service struct {
	R2ConfigInput *R2Config
	Config        *aws.Config
}

type R2Config struct {
	R2AccountId       string
	R2AccessKeyId     string
	R2AccessKeySecret string
	R2BucketName      string
}

type R2ServiceInterface interface {
	UploadFile() error
}

func NewR2Service(cfg *R2Config) *R2Service {
	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: fmt.Sprintf("https://%s.r2.cloudflarestorage.com", cfg.R2AccountId),
		}, nil
	})

	r2Config, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(r2Resolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(cfg.R2AccessKeyId, cfg.R2AccessKeySecret, "")),
		config.WithRegion("auto"),
	)
	if err != nil {
		log.Fatal(err)
	}
	return &R2Service{
		R2ConfigInput: cfg,
		Config:        &r2Config,
	}
}
