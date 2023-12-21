package main

import (
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/furyanprotocol/v4-chain/protocol/app"
	"github.com/furyanprotocol/v4-chain/protocol/app/config"
	"github.com/furyanprotocol/v4-chain/protocol/cmd/furyaprotocold/cmd"
)

func main() {
	config.SetupConfig()

	option := cmd.GetOptionWithCustomStartCmd()
	rootCmd := cmd.NewRootCmd(option)

	cmd.AddTendermintSubcommands(rootCmd)
	cmd.AddInitCmdPostRunE(rootCmd)

	if err := svrcmd.Execute(rootCmd, app.AppDaemonName, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
