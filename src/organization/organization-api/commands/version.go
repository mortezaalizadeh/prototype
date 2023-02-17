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
		Short: "Return order-api version",
		Long:  "Return order-api version",
		Run: func(cmd *cobra.Command, args []string) {
			util.PrintInfo("Order API\n")
			util.PrintInfo(fmt.Sprintf("Copyright (C) %d, Connectilly Ltd.\n", time.Now().Year()))
			util.PrintYAML(util.GetVersion())
		},
	}
}
