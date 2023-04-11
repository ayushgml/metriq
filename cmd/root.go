/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "metriq",
	Short: "Metriq is a CLI tool for monitoring system metrics.",
	Long: `Metriq is a powerful and easy-to-use CLI tool that allows you to monitor your system's performance metrics,
such as CPU usage, memory usage, disk usage, and network traffic.
With Metriq, you can keep an eye on your system's health and analyze its performance over time.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
	banner:= `
	 _______  _______ _________ _______ _________ _______ 
	(       )(  ____ \\__   __/(  ____ )\__   __/(  ___  )
	| () () || (    \/   ) (   | (    )|   ) (   | (   ) |
	| || || || (__       | |   | (____)|   | |   | |   | |
	| |(_)| ||  __)      | |   |     __)   | |   | |   | |
	| |   | || (         | |   | (\ (      | |   | | /\| |
	| )   ( || (____/\   | |   | ) \ \_____) (___| (_\ \ |
	|/     \|(_______/   )_(   |/   \__/\_______/(____\/_)
`
	fmt.Println(banner)
	fmt.Println(rootCmd.Long)
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


