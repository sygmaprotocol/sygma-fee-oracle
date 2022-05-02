// Copyright 2021 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package cmd

import (
	"github.com/ChainSafe/chainbridge-fee-oracle/app"
	"github.com/ChainSafe/chainbridge-fee-oracle/base"
	"github.com/spf13/cobra"
)

type oracleServer struct {
	ConfigPath         string
	DomainConfigPath   string
	ResourceConfigPath string
	KeyFilePath        string
	KeyType            string
}

var oracleServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Start ChainBridge fee oracle main service",
	RunE:  startServer,
}

var oserver oracleServer

func init() {
	oracleServerCmd.Flags().StringVarP(&oserver.ConfigPath, "config_path", "c", "", "")
	oracleServerCmd.Flags().StringVarP(&oserver.DomainConfigPath, "domain_config_path", "d", "", "")
	oracleServerCmd.Flags().StringVarP(&oserver.ResourceConfigPath, "resource_config_path", "r", "", "")
	oracleServerCmd.Flags().StringVarP(&oserver.KeyFilePath, "keyfile_path", "k", "", "")
	oracleServerCmd.Flags().StringVarP(&oserver.KeyType, "key_type", "t", "", "Support: secp256k1")
}

func startServer(_ *cobra.Command, _ []string) error {
	appBase := base.NewFeeOracleAppBase(oserver.ConfigPath, oserver.DomainConfigPath, oserver.ResourceConfigPath, oserver.KeyFilePath, oserver.KeyType)

	app.NewFeeOracleApp(appBase).Start()

	return nil
}
