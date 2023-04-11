/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// diskCmd represents the disk command
var diskCmd = &cobra.Command{
	Use:   "disk",
	Short: "Get the disk usage of the system",
	Long: `Get the disk usage of the system`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get current disk usage
		blue := color.New(color.FgBlue).Add(color.Bold)
		red := color.New(color.FgRed)
		diskStat, err := disk.Usage("/")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		
		blue.Println("Disk Usage:")
		red.Printf("Used Disk Percentage: %.2f%%\n", diskStat.UsedPercent)
		fmt.Printf("Total Disk: %.2f GB\n", float64(diskStat.Total)/1073741824)
		fmt.Printf("Available Disk: %.2f GB\n", float64(diskStat.Free)/1073741824)
		fmt.Printf("Used Disk: %.2f GB\n", float64(diskStat.Used)/1073741824)

		partition, err := disk.Partitions(false)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println()
		blue.Println("Disk Partitions:")
		for _, part := range partition {
			fmt.Printf("Device: %s\tMountpoint: %s\tFstype: %s\n", part.Device, part.Mountpoint, part.Fstype)
		}

		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(diskCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// diskCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// diskCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
