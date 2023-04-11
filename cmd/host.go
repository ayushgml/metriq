/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// hostCmd represents the host command
var hostCmd = &cobra.Command{
	Use:   "host",
	Short: "Get the host info",
	Long: `Get the host info`,
	Run: func(cmd *cobra.Command, args []string) {
		blue := color.New(color.FgBlue).Add(color.Bold)
		red := color.New(color.FgRed)
		yellow := color.New(color.FgYellow)
		// Get current host info
		hostInfo, err := host.Info()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		blue.Println("Host Info:")
		fmt.Printf("Hostname: %s\n", hostInfo.Hostname)
		fmt.Printf("Uptime: %.2f minutes\n", float64(hostInfo.Uptime)/60)
		fmt.Printf("Boot Time: %.2f minutes\n", float64(hostInfo.BootTime)/60)
		fmt.Println("Procs: ", hostInfo.Procs)
		fmt.Printf("OS: %s\n", hostInfo.OS)
		fmt.Printf("Platform: %s\n", hostInfo.Platform)
		fmt.Printf("Platform Family: %s\n", hostInfo.PlatformFamily)
		fmt.Printf("Platform Version: %s\n", hostInfo.PlatformVersion)
		fmt.Printf("Host ID: %s\n", hostInfo.HostID)
		red.Println("Don't show host ID to anyone!")
		

		fmt.Println()
		// Temperature stats
		tempinfo, err := host.SensorsTemperatures()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		blue.Println("Sensors temperature Info:")
		// Sort based on temperatures descending order
		for _, temp := range tempinfo {
			if(temp.Temperature < 50){
				fmt.Printf("%s: %.2f°C\n", temp.SensorKey, temp.Temperature)
			}else if(temp.Temperature < 70){
				yellow.Printf("%s: %.2f°C\n", temp.SensorKey, temp.Temperature)
			} else {
				red.Printf("%s: %.2f°C\n", temp.SensorKey, temp.Temperature)
			}
		}

		userInfo, err := host.Users()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		blue.Println("Users Info:")
		for _, user := range userInfo {
			fmt.Printf("User: %s\n", user.User)
			fmt.Printf("Terminal: %s\n", user.Terminal)
			fmt.Printf("Started: %.2f minutes\n", float64(user.Started)/60)
			fmt.Println()
		}



	},
}

func init() {
	rootCmd.AddCommand(hostCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hostCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hostCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
