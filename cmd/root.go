// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "sygma-fee-oracle",
	Short: "sygma fee oracle CLI",
	Long:  "sygma fee oracle CLI",
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
