package cmd

import (
	"os"
	"weather/cmd/current"
	"weather/cmd/forecast"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "weather",
	Short: "A cmd line weather app",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// Add custom commands
func addSubcommandPalettes() {
	rootCmd.AddCommand(forecast.ForecastCmd)
	rootCmd.AddCommand(current.CurrentCmd)
}

// User init called
func init() {
	// Add custom commands
	addSubcommandPalettes()
}
