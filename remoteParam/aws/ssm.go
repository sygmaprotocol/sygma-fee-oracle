// Copyright 2021 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package aws

import (
	"context"
	"github.com/ChainSafe/sygma-fee-oracle/remoteParam"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

var _ remoteParam.RemoteParamOperator = (*SSMClient)(nil)

type SSMClient struct {
	client *ssm.Client
}

func NewSSMClient(cfg aws.Config) *SSMClient {
	return &SSMClient{client: ssm.NewFromConfig(cfg)}
}

func (s *SSMClient) LoadParameter(parameterName string) (*remoteParam.RemoteParamResult, error) {
	input := &ssm.GetParameterInput{
		Name: &parameterName,
	}

	out, err := s.client.GetParameter(context.Background(), input)
	if err != nil {
		return nil, err
	}

	return &remoteParam.RemoteParamResult{
		Name:  *out.Parameter.Name,
		Value: *out.Parameter.Value,
	}, nil
}
