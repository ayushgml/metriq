/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var onlyUsedPercent bool
var inMB bool

var ramCmd = &cobra.Command{
	Use:   "ram",
	Short: "Get the RAM usage of the system",
	Long:  `Get the RAM usage of the system`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get current RAM usage
		vmStat, err := mem.VirtualMemory()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Define unit conversion and unit label
		unitConversion := 1073741824.0
		unitLabel := "GB"
		if inMB {
			unitConversion = 1048576.0
			unitLabel = "MB"
		}

		if onlyUsedPercent {
			fmt.Printf("Used RAM Percentage: %.2f%%\n", vmStat.UsedPercent)
		} else {
			blue := color.New(color.FgBlue).Add(color.Bold)
			red := color.New(color.FgRed)
			blue.Println("RAM Usage:")
			red.Printf("Used RAM Percentage: %.2f%%\n", vmStat.UsedPercent)
			fmt.Printf("Total RAM: %.2f %s\n", float64(vmStat.Total)/unitConversion, unitLabel)
			fmt.Printf("Available RAM: %.2f %s\n", float64(vmStat.Available)/unitConversion, unitLabel)
			fmt.Printf("Used RAM: %.2f %s\n", float64(vmStat.Used)/unitConversion, unitLabel)
			fmt.Printf("Free RAM: %.2f %s\n", float64(vmStat.Free)/unitConversion, unitLabel)
			fmt.Printf("Active RAM: %.2f %s\n", float64(vmStat.Active)/unitConversion, unitLabel)
			fmt.Printf("Inactive RAM: %.2f %s\n", float64(vmStat.Inactive)/unitConversion, unitLabel)
			fmt.Printf("Wired RAM: %.2f %s\n", float64(vmStat.Wired)/unitConversion, unitLabel)
			fmt.Printf("Cached RAM: %.2f %s\n", float64(vmStat.Cached)/unitConversion, unitLabel)
		}
	},
}

func init() {
	rootCmd.AddCommand(ramCmd)

	// Add flags
	ramCmd.Flags().BoolVarP(&onlyUsedPercent, "used-only", "u", false, "Display only the used RAM percentage")
	ramCmd.Flags().BoolVarP(&inMB, "in-mb", "m", false, "Display RAM values in MB instead of GB")
}
