package commands

import (
	"github.com/Connectilly/connectilly/src/shared/connectillyctl/commands/generate"
	"github.com/spf13/cobra"
)

func generateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "generate",
	}

	cmd.AddCommand(
		generate.ClientCommand(),
	)

	return cmd
}
