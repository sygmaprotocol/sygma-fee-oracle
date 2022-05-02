// Copyright 2021 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package e2e

import (
	"github.com/ChainSafe/chainbridge-fee-oracle/e2e/setup"
	"testing"
)

func TestDataPrepare(t *testing.T) {
	// setup testdata to database
	err := setup.DataPrepare("./testdb")
	if err != nil {
		panic(err)
	}
}
