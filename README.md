# Azure-Inventory-CLI-Tool
Azure Resource Inventory CLI Tool: Lists Azure resources given a subscription ID and export to CSV file

## Requirements

- Create Service Principal or Manage Identity to login in Azure
- Assign at least Reader role on your subscription
- Create .env file with your Azure credenctials, for example using a Service Principal:

```
AZURE_CLIENT_ID=xxxxx
AZURE_CLIENT_SECRET=xxxxx
AZURE_TENANT_ID=xxxxx
```

- Install the binary on your computer with go:

```bash
go install github.com/AuroraPalma/Azure-Inventory-CLI-Tool@latest
```

or 

- Download the binary and save in a folder like C:\azure-inventory-cli-tool, add that path to your system PATH (via System Properties > Environment Variables)

## Usage

```bash
azure-inventory-cli-tool --subscription-id <your-subscription-id> --env-file <your path to .env file>
```
