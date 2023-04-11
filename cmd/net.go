/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// netCmd represents the net command
var netCmd = &cobra.Command{
	Use:   "net",
	Short: "Get the network usage of the system",
	Long: `Get the network usage of the system`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get current network usage
		netStat, err := net.IOCounters(true)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		arr := make([]uint64, len(netStat))
		for i, v := range netStat {
			arr[i] = v.BytesSent
		}
		for i := 0; i < len(arr); i++ {
			for j := i + 1; j < len(arr); j++ {
				if arr[i] < arr[j] {
					arr[i], arr[j] = arr[j], arr[i]
				}
			}
		}
		blue := color.New(color.FgBlue).Add(color.Bold)
		red := color.New(color.FgRed)
		blue.Println("Network Usage (In Descending Order of bytes sent):")
		for i := 0; i < 5; i++ {
			for _, v1 := range netStat {
				if arr[i] == v1.BytesSent {
					red.Printf("Interface: %s\n", v1.Name)
					fmt.Printf("Bytes Sent: %d\n", v1.BytesSent)
					fmt.Printf("Bytes Received: %d\n", v1.BytesRecv)
					fmt.Printf("Packets Sent: %d\n", v1.PacketsSent)
					fmt.Printf("Packets Received: %d\n", v1.PacketsRecv)
					fmt.Printf("Error In: %d\n", v1.Errin)
					fmt.Printf("Error Out: %d\n", v1.Errout)
					fmt.Printf("Drop In: %d\n", v1.Dropin)
					fmt.Printf("Drop Out: %d\n", v1.Dropout)
					fmt.Println()
				}
			}
		}

		
	},
}

func init() {
	rootCmd.AddCommand(netCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// netCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// netCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
