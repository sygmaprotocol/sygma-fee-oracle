// Copyright 2021 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

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
