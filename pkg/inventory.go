package pkg

import (
    "context"
    "fmt"
    "log"

    "github.com/Azure/azure-sdk-for-go/sdk/azcore"
    "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
    "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
)

// Returns an Azure credential for authentication
func GetAzureClient() (azcore.TokenCredential, error) {
    cred, err := azidentity.NewDefaultAzureCredential(nil)
    if err != nil {
        log.Fatalf("failed to obtain a credential: %v", err)
    }
    return cred, nil
}

// Lists resource names in a subscription
func ListResourcesSub(cred azcore.TokenCredential, subscriptionID string) ([]string, error) {
    client, err := armresources.NewClient(subscriptionID, cred, nil)
    if err != nil {
        return nil, fmt.Errorf("failed to create resources client: %v", err)
    }
	ctx := context.Background()

    pager := client.NewListPager(nil)
    var resources []string
    for pager.More() {
        page, err := pager.NextPage(ctx)
        if err != nil {
            return nil, fmt.Errorf("failed to get next page of resources: %v", err)
        }
        for _, resource := range page.Value {
            resources = append(resources, *resource.Name)
        }
    }
    return resources, nil
}