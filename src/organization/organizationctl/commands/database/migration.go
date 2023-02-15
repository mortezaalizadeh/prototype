package database

import (
	"github.com/Connectilly/connectilly/src/organization/organizationctl/commands/database/migration"
	"github.com/spf13/cobra"
)

func MigrateCommand(connectionString *string) *cobra.Command {
	cmd := &cobra.Command{
		Use: "migration",
	}

	cmd.AddCommand(
		migration.AddCommand(connectionString),
	)

	return cmd
}
