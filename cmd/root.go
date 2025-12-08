package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "tglint",
	Short:   "Formatter and linter for terragrunt.hcl",
	Long:    `tglint is a fast, CI-friendly formatter and linter for Terragrunt HCL files.`,
	Version: Version,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(3)
	}
}

func init() {
	rootCmd.AddCommand(fmtCmd)
	rootCmd.AddCommand(lintCmd)
}
