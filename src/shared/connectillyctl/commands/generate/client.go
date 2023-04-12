package generate

import (
	"github.com/Connectilly/connectilly/src/shared/connectillyctl/commands/generate/client"
	"github.com/spf13/cobra"
)

func ClientCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "client",
	}

	cmd.AddCommand(
		client.EventCommand(),
	)

	return cmd
}
