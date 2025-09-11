package cli

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"scope3/go-interview/api"

	"github.com/spf13/cobra"
)

var (
	client *api.Client
)

func initClient() *api.Client {
	if client == nil {
		client = api.NewClient()
	}
	return client
}

func init() {
	cobra.OnInitialize(func() { initClient() })
	rootCmd.AddCommand(probeCmd)
	rootCmd.AddCommand(compareCmd)
}

var rootCmd = &cobra.Command{
	Use:   "measure-cli",
	Short: "a cli for interacting and parsing the /measure endpoint",
	Run: func(cmd *cobra.Command, args []string) {
		log.Warn("root does nothing")
	},
}

var probeCmd = &cobra.Command{
	Use:   "probe",
	Short: "tests that the API is reachable and API_KEY is set",
	Run: func(cmd *cobra.Command, args []string) {
		client := initClient()
		response, err := client.Measure([]string{"yahoo.com"}, "2025-05-01")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(response.String())
	},
}

var compareCmd = &cobra.Command{
	Use:   "compare [properties...]",
	Short: "compares emission data across properties",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		date, _ := cmd.Flags().GetString("date")

		options := map[string]interface{}{
			"date": date,
		}

		result := map[string]interface{}{
			"properties": args,
			"options":    options,
		}

		fmt.Printf("%+v compare arguments!\n", result)
	},
}

func init() {
	compareCmd.Flags().StringP("date", "d", "", "Date to compare against")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
