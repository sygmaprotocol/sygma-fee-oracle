// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package oracle

import (
	_ "github.com/golang/mock/mockgen/model"
	"github.com/pkg/errors"
)

var (
	ErrNotSupported = errors.New("unsupported source")
)

type Oracle interface {
	// Source of the oracle. Because one oracle implementation can be used for different oracle sources
	Source() string
	IsEnabled() bool
}
