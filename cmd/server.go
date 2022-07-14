package cmd

import "github.com/spf13/cobra"

var (
	serverCmd = &cobra.Command{Use: "server", Short: "Start the engine", Long: CmdName + " it's main process starts by checking strategies and executing orders based on the buy/sell signals these strategies generate", Aliases: []string{"serve", "boot", "start"}, PreRun: func(cmd *cobra.Command, args []string) {}}
)

func init() {
	rootCmd.AddCommand(serverCmd)

	// Declare daemon flag '-d' to run as background process
	Daemon := false
	serverCmd.PersistentFlags().BoolVarP(&Daemon, "daemon", "d", false, "Use the flag '-d' to run main process in background (daemonize)")
}
