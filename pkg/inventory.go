import (
	"context"
	"errors"
	"fmt"
    "github.com/Azure/azure-sdk-for-go/sdk/azcore"
    "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

func GetAzureClient() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)

	}
	return cred, nil
}

func ListResourcesSub(cred NewDefaultAzureCredential, subscriptionID string) ([]string, error) {
	client, err := armresources.NewResourcesClient(subscriptionID, cred, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create resources client: %v", err)
	}

	pager := client.NewListPager(nil)

	var resources []string
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		if err != nil {
			return nil, fmt.Errorf("failed to get next page of resources: %v", err)
		}
		for _, resource := range page.Value {
			resources = append(resources, *resource.Name)
		}
	}
	return resources, nil
}