/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/process"
	"github.com/spf13/cobra"
	"github.com/fatih/color"
	"os"
)

// cpuCmd represents the cpu command
var cpuCmd = &cobra.Command{
	Use:   "cpu",
	Short: "Get the CPU usage of the system",
	Long: `Get the CPU usage of the system`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get current CPU usage
		p, err := process.NewProcess(int32(os.Getpid()))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		cpuPercent, err := p.CPUPercent()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		blue := color.New(color.FgBlue).Add(color.Bold)
		red := color.New(color.FgRed)
		blue.Println("CPU Usage:")
		red.Printf("Used CPU Percentage: %.2f%%\n", cpuPercent)

	},
}

func init() {
	rootCmd.AddCommand(cpuCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cpuCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cpuCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
