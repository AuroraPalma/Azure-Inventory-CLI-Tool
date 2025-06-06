package main

import (
    "azure-inventory-cli-tool/pkg"
    "os"
    "log"

    "github.com/joho/godotenv"
)

func main() {
    // Load .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Get subscription ID from environment variable
    subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")
}