package cmd

import (
	"fmt"
	"os"

	"github.com/pinpt/go-common/log"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "schema",
	Short: "Run the schema tool",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// in case we need to add something to the logs
func newCommandLogger(cmd *cobra.Command, opts ...log.WithLogOptions) log.LoggerCloser {
	return log.NewCommandLogger(cmd, opts...)
}
