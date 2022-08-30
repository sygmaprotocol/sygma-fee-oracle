// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func NewAWSClient() *aws.Config {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		panic("aws configuration error: " + err.Error())
	}

	return &cfg
}
