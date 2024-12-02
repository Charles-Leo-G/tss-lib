package cmd

import (
	"github.com/spf13/cobra"

	"github.com/Charles-Leo-G/tss-lib/common"
	"github.com/Charles-Leo-G/tss-lib/server"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:    "server",
	Short:  "bootstrap and relay server helps node (dynamic ip) discovery and NAT traversal",
	Long:   "bootstrap and relay server helps node (dynamic ip) discovery and NAT traversal",
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		server.NewTssBootstrapServer(common.TssCfg.Home, common.TssCfg.P2PConfig)
		select {}
	},
}
