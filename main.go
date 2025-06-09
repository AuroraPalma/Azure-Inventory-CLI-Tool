package main

import (
    "os"
    "log"

    "github.com/joho/godotenv"
    "azure-inventory-cli-tool/pkg"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")

    cred, err := pkg.GetAzureClient()
    if err != nil {
        log.Fatalf("Failed to get Azure client: %v", err)
    }
    resources, err := pkg.ListResourcesSub(cred, subscriptionID)
    if err != nil {
        log.Fatalf("Failed to list resources: %v", err)
    }

    for _, name := range resources {
        log.Println(name)
    }

}