// Copyright 2021 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package remoteParam

type RemoteParamResult struct {
	Name  string
	Value string
}

type RemoteParamOperator interface {
	LoadParameter(paramName string) (*RemoteParamResult, error)
}
