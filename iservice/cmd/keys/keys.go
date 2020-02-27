package keys

import (
	"os"
	"path/filepath"

	"github.com/irisnet/irishub-sdk-go/crypto"
	"github.com/irisnet/irishub-sdk-go/types"
	"github.com/spf13/cobra"
)

var keysPath = os.ExpandEnv(filepath.Join("$HOME", ".iservice"))

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
			keyDao := NewKeyDAO(keysPath)
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
			err = keyDao.Write(args[0], keyInfo)
			if err != nil {
				return err
			}
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
