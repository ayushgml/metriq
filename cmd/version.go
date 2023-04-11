/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get the version of the CLI tool",
	Long: `Get the verison of the CLI tool`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("metriq v1.0.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
