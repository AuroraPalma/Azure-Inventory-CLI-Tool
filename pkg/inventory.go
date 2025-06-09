package pkg

import (
    "context"
    "fmt"
    "log"
	"encoding/csv"
	"os"

    "github.com/Azure/azure-sdk-for-go/sdk/azcore"
    "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
    "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

type Resource struct {
    Name string
    Group string
	Type string
}

// Returns an Azure credential for authentication
func GetAzureClient() (azcore.TokenCredential, error) {
    cred, err := azidentity.NewDefaultAzureCredential(nil)
    if err != nil {
        log.Fatalf("failed to obtain a credential: %v", err)
    }
    return cred, nil
}

// Lists resource names in a subscription
func ListResourcesSub(cred azcore.TokenCredential, subscriptionID string) ([]Resource, error) {
    client, err := armresources.NewClient(subscriptionID, cred, nil)
    if err != nil {
        return nil, fmt.Errorf("failed to create resources client: %v", err)
    }
	ctx := context.Background()

    pager := client.NewListPager(nil)
    var resources []Resource
	fmt.Println("Listing resources in subscription:", subscriptionID)
    for pager.More() {
        page, err := pager.NextPage(ctx)
        if err != nil {
            return nil, fmt.Errorf("failed to get next page of resources: %v", err)
        }
        for _, resource := range page.Value {
			resourceGroup := ""
			resourceType := ""
			if resource.ID != nil {
				parts, err := arm.ParseResourceID(*resource.ID)
				type_resource , err := arm.ParseResourceType(*resource.Type)
				if err != nil {
					panic(fmt.Sprintf("failed to parse resource ID: %v", err))
				}
				resourceGroup = parts.ResourceGroupName
				resourceType = type_resource.Type
			}
            resources = append(resources, Resource{
				Name:  *resource.Name,
				Group: resourceGroup,
				Type:  resourceType,
			})
        }
    }
	fmt.Printf("Total resources found: %d\n", len(resources))
    return resources, nil
}

func SaveResourcesToCSV(resources []Resource, filename string) error {
	csvFile, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create CSV file: %v", err)
	}
	defer csvFile.Close()

	csvwriter := csv.NewWriter(csvFile)
	defer csvwriter.Flush()

	csvwriter.Write([]string{"Name", "Group", "Type"}) // Write header row

	for _, resource := range resources {
		err := csvwriter.Write([]string{resource.Name, resource.Group, resource.Type})
		if err != nil {
			return fmt.Errorf("failed to write to CSV file: %v", err)
		}
	}
	log.Println("Resources have been written to resources.csv")
	return nil
}