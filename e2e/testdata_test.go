// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package e2e

import (
	"github.com/ChainSafe/sygma-fee-oracle/e2e/setup"
	"testing"
)

func TestDataPrepare(t *testing.T) {
	// setup testdata to database
	err := setup.DataPrepare("./testdb")
	if err != nil {
		panic(err)
	}
}
