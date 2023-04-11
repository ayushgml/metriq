/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/spf13/cobra"
)

var continuous bool
var showRecommendations bool

func generateRecommendations(loadScore, memScore, cpuScore, netScore, diskScore float64) {
	// Get the operating system information
	hostStat, err := host.Info()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
	os := strings.ToLower(hostStat.Platform)

	// Display recommendations for low-scoring components
	if loadScore < 0.3 || memScore < 0.3 || cpuScore < 0.3 || netScore < 0.3 || diskScore < 0.3 {
		fmt.Println("\nRecommendations:")

		if loadScore < 0.3 {
			fmt.Println("\n- High Load:")
			fmt.Println("  * Close unnecessary applications")
			fmt.Println("  * Check for resource-intensive processes and consider terminating or optimizing them")
		}

		if memScore < 0.3 {
			fmt.Println("\n- High Memory Usage:")
			fmt.Println("  * Close unnecessary applications")
			fmt.Println("  * Check for memory leaks in running applications")
			fmt.Println("  * Upgrade your system's RAM")
		}

		if cpuScore < 0.3 {
			fmt.Println("\n- High CPU Usage:")
			fmt.Println("  * Close unnecessary applications")
			fmt.Println("  * Check for CPU-intensive processes and consider terminating or optimizing them")
			fmt.Println("  * Upgrade your system's CPU")
		}

		if netScore < 0.3 {
			fmt.Println("\n- High Network Usage:")
			fmt.Println("  * Limit bandwidth usage by closing unnecessary network connections")
			fmt.Println("  * Check for network-intensive processes and consider terminating or optimizing them")
			fmt.Println("  * Upgrade your network hardware")
		}

		if diskScore < 0.3 {
			fmt.Println("\n- High Disk Usage:")
			fmt.Println("  * Clean up unnecessary files and applications")
			fmt.Println("  * Check for disk-intensive processes and consider terminating or optimizing them")
			fmt.Println("  * Upgrade your system's storage device")
		}

		// Provide OS-specific instructions
		fmt.Println("\nOS-specific instructions:")
		switch os {
		case "windows":
			fmt.Println("  * Use Task Manager to monitor and manage processes")
			fmt.Println("  * Use Disk Cleanup to clean up unnecessary files")
		case "darwin": // macOS
			fmt.Println("  * Use Activity Monitor to monitor and manage processes")
			fmt.Println("  * Use built-in or third-party tools to clean up unnecessary files")
		case "linux":
			fmt.Println("  * Use 'top', 'htop', or a system monitor application to monitor and manage processes")
			fmt.Println("  * Use 'rm' or a file manager to clean up unnecessary files")
		default:
			fmt.Println("  * Refer to your operating system's documentation for monitoring and managing processes and disk usage")
		}
	} else {
		fmt.Println("\nYour system health is in good condition. No recommendations needed.")
	}
}



// healthCmd represents the health command
var healthCmd = &cobra.Command{
	Use:   "health",
	Short: "Get the health of the system",
	Long:  `Get the health of the system`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get the overrall health of the system
		// Get the current load
		for {
			loadStat, err := load.Avg()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			// Get the current memory usage
			memStat, err := mem.VirtualMemory()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			// Get the current CPU usage
			cpuStat, err := cpu.Percent(0, false)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			// Get the current network usage
			netStat, err := net.IOCounters(true)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			// Get the current disk usage
			diskStat, err := disk.Usage("/")
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			// Get the current host info
			hostStat, err := host.Info()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			loadWeight := 0.2
			memWeight := 0.2
			cpuWeight := 0.2
			netWeight := 0.2
			diskWeight := 0.2

			// Normalize factors
			loadScore := 1 - loadStat.Load1/float64(hostStat.Procs)
			memScore := 1 - memStat.UsedPercent/100
			cpuScore := 1 - cpuStat[0]/100
			netScore := 1 - (float64(netStat[0].BytesSent)+float64(netStat[0].BytesRecv))/float64(netStat[0].BytesSent+netStat[0].BytesRecv+1)
			diskScore := 1 - diskStat.UsedPercent/100

			// Calculate the health score
			score := loadScore*loadWeight + memScore*memWeight + cpuScore*cpuWeight + netScore*netWeight + diskScore*diskWeight

			blue := color.New(color.FgBlue)
			red := color.New(color.FgRed)
			green := color.New(color.FgGreen)
			yellow := color.New(color.FgYellow)

			redTitle := color.New(color.FgRed).Add(color.Italic).Add(color.Underline)
			greenTitle := color.New(color.FgGreen).Add(color.Italic).Add(color.Underline)
			yellowTitle := color.New(color.FgYellow).Add(color.Italic).Add(color.Underline)
			blue.Println("This score is calculated based on overall System load, Memory usage, CPU usage, Network usage and Disk usage.")
			if score > 0.70 {
				greenTitle.Printf("Health Score: %.2f%%\n", score*100)
			} else if score > 0.3 {
				yellowTitle.Printf("Health Score: %.2f%%\n", score*100)
			} else {
				redTitle.Printf("Health Score: %.2f%%\n", score*100)
			}
			fmt.Println()
			if loadScore > 0.70 {
				green.Printf("Load Score: %.2f%%\n", loadScore*100)
			}
			if memScore > 0.70 {
				green.Printf("Memory Score: %.2f%%\n", memScore*100)
			}
			if cpuScore > 0.70 {
				green.Printf("CPU Score: %.2f%%\n", cpuScore*100)
			}
			if netScore > 0.70 {
				green.Printf("Network Score: %.2f%%\n", netScore*100)
			}
			if diskScore > 0.70 {
				green.Printf("Disk Score: %.2f%%\n", diskScore*100)
			}

			if loadScore > 0.30 && loadScore < 0.70 {
				yellow.Printf("Load Score: %.2f%%\n", loadScore*100)
			}
			if memScore > 0.30 && memScore < 0.70 {
				yellow.Printf("Memory Score: %.2f%%\n", memScore*100)
			}
			if cpuScore > 0.30 && cpuScore < 0.70 {
				yellow.Printf("CPU Score: %.2f%%\n", cpuScore*100)
			}
			if netScore > 0.30 && netScore < 0.70 {
				yellow.Printf("Network Score: %.2f%%\n", netScore*100)
			}
			if diskScore > 0.30 && diskScore < 0.70 {
				yellow.Printf("Disk Score: %.2f%%\n", diskScore*100)
			}

			if loadScore < 0.30 {
				red.Printf("Load Score: %.2f%%\n", loadScore*100)
			}
			if memScore < 0.30 {
				red.Printf("Memory Score: %.2f%%\n", memScore*100)
			}
			if cpuScore < 0.30 {
				red.Printf("CPU Score: %.2f%%\n", cpuScore*100)
			}
			if netScore < 0.30 {
				red.Printf("Network Score: %.2f%%\n", netScore*100)
			}
			if diskScore < 0.30 {
				red.Printf("Disk Score: %.2f%%\n", diskScore*100)
			}

			// fmt.Println()
			if showRecommendations {
				generateRecommendations(loadScore, memScore, cpuScore, netScore, diskScore)
			}

			if !continuous {
				break
			}

			time.Sleep(1 * time.Second)
			fmt.Println("\033[2J") // Clear the terminal screen
		}

	},
}

func init() {
	rootCmd.AddCommand(healthCmd)

	healthCmd.Flags().BoolVarP(&continuous, "continuous", "c", false, "Monitor system health continuously every second")
	healthCmd.Flags().BoolVarP(&showRecommendations, "recommendations", "r", false, "Show recommendations based on health scores")
}
