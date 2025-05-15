package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	cobra.OnInitialize(func() { initClient() })
	rootCmd.AddCommand(trendsCmd)
}

var trendsCmd = &cobra.Command{
	Use:   "trends",
	Short: "Analyze trends in emissions",
	Long:  "Analyze trends in emissions",
	Run: func(cmd *cobra.Command, args []string) {
		domain := args[0]
		fmt.Printf("Analyzing trends for %s...\n", domain)
	},
}
