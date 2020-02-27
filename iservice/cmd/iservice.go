package main

import (
	"iservice/iservice/cmd/keys"
	"iservice/iservice/node"
	"os"

	sdk "github.com/irisnet/irishub-sdk-go/types"
	"github.com/spf13/cobra"
)

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(keys.Commands(), startCmd())
}

var rootCmd = &cobra.Command{
	Use:          "iservice",
	Short:        "iservice daemon",
	SilenceUsage: true,
}

func startCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "start",
		Short:   "start daemon",
		Example: `iservice start --node <tendermint rpc>`,
		RunE: func(cmd *cobra.Command, args []string) error {
			config := sdk.SDKConfig{
				NodeURI: NodeURI,
				Network: Network,
				ChainID: ChainID,
				Gas:     Gas,
				Fee:     Fee,
				Mode:    Mode,
				Online:  Online,
			}
			baseTx := sdk.BaseTx{
				From:     args[0],
				Password: args[1],
				Gas:      Gas,
				Fee:      Fee,
				Memo:     "service",
				Mode:     Mode,
			}
			node.Start(config, baseTx)
			return nil
		},
	}
	return cmd
}

const (
	NodeURI = "localhost:26657"
	ChainID = "test"
	Online  = true
	Network = sdk.Testnet
	Mode    = sdk.Commit
	Fee     = "600000000000000000iris-atto"
	Gas     = 20000
)
