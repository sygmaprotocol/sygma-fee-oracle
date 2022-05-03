// Copyright 2021 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "chainbridge-fee-oracle",
	Short: "Chainbridge fee oracle CLI",
	Long:  "Chainbridge fee oracle CLI",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmd.AddCommand(oracleServerCmd)
	RootCmd.AddCommand(oracleKeyGenerationCmd)
}
