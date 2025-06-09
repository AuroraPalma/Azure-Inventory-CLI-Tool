# Azure-Inventory-CLI-Tool
Azure Resource Inventory CLI Tool: Lists Azure resources across subscriptions

1. Initialize the Go module (replace azure-inventory-cli-tool with your desired module name, usually your repo path)

```bash
go mod init azure-inventory-cli-tool
```

2. Add dependencies (this will also create/update go.sum):

```bash
go get github.com/Azure/azure-sdk-for-go/sdk/azidentity
go get github.com/Azure/azure-sdk-for-go/sdk/azcore
go get github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources
```

Now, go.mod and go.sum will be created and managed automatically.

Summary:

```bash
go.mod tracks your module and its requirements.
go.sum ensures dependency integrity.
Use go mod tidy anytime to clean up unused dependencies.
```