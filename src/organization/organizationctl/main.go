package main

import (
	"github.com/Connectilly/connectilly/src/organization/organizationctl/commands"
	"github.com/micro-business/go-core/pkg/util"
)

func main() {
	rootCmd := commands.Root()
	util.PrintIfError(rootCmd.Execute())
}
