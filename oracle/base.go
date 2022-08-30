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
	Name() string
	IsEnabled() bool
}
