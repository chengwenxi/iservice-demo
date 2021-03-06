package main

import (
	"iservice/iservice/cmd/keys"
	"iservice/iservice/market"
	"iservice/iservice/node"
	"os"

	"github.com/irisnet/irishub-sdk-go/types"
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
		Example: `iservice start [key_name] [market]`,
		Args:    cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			// set market
			if len(args) > 1 {
				market.MarketType = args[1]
			}
			config := types.ClientConfig{
				NodeURI: NodeURI,
				Network: Network,
				ChainID: ChainID,
				Gas:     Gas,
				Fee:     Fee,
				Mode:    Mode,
				Level:   "debug",
				KeyDAO:  keys.NewKeyDAO(),
			}
			baseTx := types.BaseTx{
				From:     args[0],
				Gas:      Gas,
				Fee:      Fee,
				Mode:     Mode,
				Password: "",
			}
			node.Start(config, baseTx)
			return nil
		},
	}
	return cmd
}

var (
	feeAmt, _ = types.NewIntFromString("600000000000000000")
	Fee       = types.NewDecCoins(types.NewDecCoin("iris-atto", feeAmt))
)

const (
	NodeURI = "localhost:26657"
	ChainID = "test"
	Online  = true
	Network = types.Testnet
	Mode    = types.Async
	Gas     = 100000
)
