package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use: "version", Aliases: []string{"v"},
	Short: "Print the version number of " + CmdName,
	Long:  `All software has versions. This is ` + CmdName + `'s`,
	Run: func(cmd *cobra.Command, _ []string) {
		fmt.Printf("%v's version is:\n%v", CmdName, cmd.Version)
	},
}
