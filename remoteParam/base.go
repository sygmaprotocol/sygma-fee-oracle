// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package remoteParam

type RemoteParamResult struct {
	Name  string
	Value string
}

type RemoteParamOperator interface {
	LoadParameter(paramName string) (*RemoteParamResult, error)
}
