package keys

import (
	"os"
	"os/user"

	"github.com/spf13/cobra"
)

var keysPath = ".iservice"

func init() {
	u, err := user.Current()
	if err != nil && u != nil {
		keysPath = u.HomeDir + string(os.PathSeparator) + ".iservice"
	}
}

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
		Use:   "add <name>",
		Short: "Create a new key",
		RunE: func(cmd *cobra.Command, args []string) error {

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
