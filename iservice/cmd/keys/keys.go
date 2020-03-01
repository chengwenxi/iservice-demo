package keys

import (
	"fmt"

	"github.com/irisnet/irishub-sdk-go/crypto"
	"github.com/irisnet/irishub-sdk-go/types"
	"github.com/spf13/cobra"
)

func Commands() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "keys",
		Short: "Add or view local private keys",
	}
	cmd.AddCommand(
		addKeyCommand(),
		listKeysCmd(),
		deleteKeyCommand(),
	)
	return cmd
}

func addKeyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add [name]",
		Short: "Create a new key",
		RunE: func(cmd *cobra.Command, args []string) error {
			keyDao := NewKeyDAO()
			keyManager, err := crypto.NewKeyManager()
			if err != nil {
				return err
			}
			keyManager.GetPrivKey()

			var keyInfo types.KeyInfo
			privateKey, err := keyManager.ExportAsPrivateKey()
			if err != nil {
				return err
			}
			keyInfo.PrivKey = privateKey
			keyInfo.Address = types.AccAddress(keyManager.GetPrivKey().PubKey().Address()).String()
			err = keyDao.Write(args[0], keyInfo)
			if err != nil {
				return err
			}
			mnemonic, _ := keyManager.ExportAsMnemonic()
			fmt.Printf(`add key %s successful.
address: %s
mnemonic: %s
`, args[0], keyInfo.Address, mnemonic)
			return nil
		},
	}
	return cmd
}

func listKeysCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all keys",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	return cmd
}

func deleteKeyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete <name>",
		Short: "Delete the given key",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	return cmd
}
