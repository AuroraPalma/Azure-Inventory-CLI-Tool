package cmd

import (
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
    "azure-inventory-cli-tool/pkg"
    "github.com/spf13/cobra"
)

var subscriptionID string
var envfile string

func init() {
	rootCmd.Flags().StringVar(&subscriptionID, "subscription-id", "", "Azure Subscription ID")
	rootCmd.MarkFlagRequired("subscription-id")
	if err := rootCmd.MarkFlagRequired("subscription-id"); err != nil {
		log.Fatalf("Error marking subscription-id flag as required: %v", err)
	}
	rootCmd.Flags().StringVar(&envfile, "env-file", ".env", "Path to the .env file containing Azure credentials")
	rootCmd.MarkFlagRequired("env-file")
	if err := rootCmd.MarkFlagRequired("env-file"); err != nil {
		log.Fatalf("Error marking env-file flag as required: %v", err)
	}
}

var rootCmd = &cobra.Command{
	Use:   "azure-inventory-cli-tool",
	Short: "A CLI tool to list Azure resources in a subscription and save to CSV",
	Long:    `azure-inventory-cli-tool is a command-line tool that lists all resources in your Azure subscription and exports them to a CSV file.`,
	Example: `azure-inventory-cli-tool --subscription-id <your-subscription-id> --env-file <path-to-env-file>`,
	Run: func(cmd *cobra.Command, args []string) {
		err := godotenv.Load(envfile)
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	
		cred, err := pkg.GetAzureClient()
		if err != nil {
			log.Fatalf("Failed to get Azure client: %v", err)
		}
		resources, err := pkg.ListResourcesSub(cred, subscriptionID)
		if err != nil {
			log.Fatalf("Failed to list resources: %v", err)
		}
		filename := fmt.Sprintf("resources_%s.csv", subscriptionID)
		err = pkg.SaveResourcesToCSV(resources, filename)
		if err != nil {
			log.Fatalf("Failed to save resources to CSV: %v", err)
		}
		log.Printf("Resources saved to %s\n", filename)
	},
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
