// Copyright 2021 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

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
