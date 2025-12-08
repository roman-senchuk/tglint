package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const Version = "1.0.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  `Print the version number of tglint`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("tglint version %s\n", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	
	// Add --version flag to root command
	rootCmd.Version = Version
	rootCmd.SetVersionTemplate("tglint version {{.Version}}\n")
}
