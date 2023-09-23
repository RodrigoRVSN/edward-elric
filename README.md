# Testing Azure services with Golang

## Deploying a Golang REST API in Azure App Service from CLI

- Create the application in Golang
- Push the repository
- Run the `az web app` commands to deploy the application 

```bash
az webapp up --runtime GO:1.19 --os linux --sku B1 --name edward-elric
```

TBD...