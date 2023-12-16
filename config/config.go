package config

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

func (app *Application) SetR2Config() {
	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: fmt.Sprintf("https://%s.r2.cloudflarestorage.com", app.R2AccountId),
		}, nil
	})

	r2Config, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(r2Resolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(app.R2AccessKeyId, app.R2AccessKeySecret, "")),
		config.WithRegion("auto"),
	)
	if err != nil {
		log.Fatal(err)
	}
	app.R2Config = &r2Config
}
