package cmd

import (
	"github.com/ChainSafe/chainbridge-fee-oracle/app"
	"github.com/ChainSafe/chainbridge-fee-oracle/base"
	"github.com/spf13/cobra"
)

type oracleServer struct {
	ConfigPath string
}

var oracleServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Start ChainBridge fee oracle main service",
	RunE:  startServer,
}

var oserver oracleServer

func init() {
	oracleServerCmd.Flags().StringVarP(&oserver.ConfigPath, "config_path", "c", "", "default: ./config/config.template.yaml")
}

func startServer(_ *cobra.Command, _ []string) error {
	appBase := base.NewFeeOracleAppBase(oserver.ConfigPath)

	app.NewFeeOracleApp(appBase).Start()

	return nil
}
