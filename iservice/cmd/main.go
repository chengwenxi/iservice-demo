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
			}
			node.Start(config, baseTx)
			return nil
		},
	}
	return cmd
}

const (
	Addr    = "faa1d3mf696gvtwq2dfx03ghe64akf6t5vyz6pe3le"
	ValAddr = "iva1x3f572u057lv88mva2q3z40ls8pup9hsg0lxcp"
	PrivKey = "927be78a5f5b63bb95ff34ed9c6e4b39b6af6d2f9f59731452de659cac9b19db"
	NodeURI = "localhost:26657"
	ChainID = "test"
	Online  = true
	Network = sdk.Testnet
	Mode    = sdk.Commit
	Fee     = "600000000000000000iris-atto"
	Gas     = 20000
)
