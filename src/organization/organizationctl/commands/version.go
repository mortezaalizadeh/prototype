package commands

import (
	"fmt"
	"time"

	"github.com/micro-business/go-core/pkg/util"
	"github.com/spf13/cobra"
)

func versionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Return organizationctl version",
		Long:  "Return organizationctl version",
		Run: func(cmd *cobra.Command, args []string) {
			util.PrintInfo("organizationctl\n")
			util.PrintInfo(fmt.Sprintf("Copyright (C) %d, Connectilly Ltd.\n", time.Now().Year()))
			util.PrintYAML(util.GetVersion())
		},
	}
}
