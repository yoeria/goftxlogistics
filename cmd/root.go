package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// Main command name
	CmdName = "goftxlogistics"
	// This var is to be used only for rootCmd. Use *rootCmd.Version instead
	CmdVersion string = getVersion()

	// Var name says all
	rootCmd = &cobra.Command{
		Use:   CmdName,
		Short: "Use the program it's features by using it's subcommands",
		Long:  "To start a daemon, use the server subcommand",
		Run: func(cmd *cobra.Command, args []string) {
		},
		Version: CmdVersion,

		// Optionals
		FParseErrWhitelist:         cobra.FParseErrWhitelist{},
		CompletionOptions:          cobra.CompletionOptions{},
		TraverseChildren:           false,
		Hidden:                     false,
		SilenceErrors:              false,
		SilenceUsage:               false,
		DisableFlagParsing:         false,
		DisableAutoGenTag:          false,
		DisableFlagsInUseLine:      false,
		DisableSuggestions:         false,
		SuggestionsMinimumDistance: 0,
	}
)



// Declare persistent flags etc...
func init() {
	Verbose := false
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Verbose output")
}

func Exec() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getVersion() string {
		getVersion, err := os.StartProcess("git", []string{"rev-parse", "--short", "HEAD"}, nil)
		if err != nil {
			fmt.Printf("Cannot parse version\nError: %v", err)
		}
		res, err := getVersion.Wait()
		if err != nil {
			fmt.Printf("Cannot parse version\nError: %v", err)
		}
		return res.String()
	}
