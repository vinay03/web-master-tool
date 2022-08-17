package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vinay03/web-master-tool/tools/benchmark"
)

var benchmarkCmd = &cobra.Command{
	Use:   "benchmark",
	Short: "Calculates average load time for a URL",
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		num, _ := cmd.Flags().GetInt("number")
		verbose, _ := cmd.Flags().GetBool("verbose")

		fmt.Println("Visiting ["+url+"] for", num, "times...")
		if verbose {
			benchmark.VerboseOutput = true
		}
		avgDuration := benchmark.Start(url, num)
		fmt.Println("Average Duration : ", avgDuration, "ms")
	},
}

func init() {
	// URL Flag to mention URL to benchmark
	benchmarkCmd.Flags().StringP("url", "u", "", "To specify URL")

	// Number of requests to be made for calculating Average
	benchmarkCmd.Flags().IntP("number", "n", 10, "Number of repetitions to perform")

	// Verbose flag
	benchmarkCmd.Flags().BoolP("verbose", "v", false, "Verbose output")

	// Must require url flag
	benchmarkCmd.MarkFlagRequired("url")

	// Add to root command
	rootCmd.AddCommand(benchmarkCmd)
}
