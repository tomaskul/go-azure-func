Simple Hello World Go API endpoint, configured to be ran as an Azure Function. Code based on following articles:
- [Build serverless app with Go](https://learn.microsoft.com/en-gb/training/modules/serverless-go/3-web-api), and
- [Quickstart: Create a Go or Rust function in Azure using Visual Studio Code](https://learn.microsoft.com/en-us/azure/azure-functions/create-first-function-vs-code-other?tabs=go%2Cwindows)

# To run this API as Azure Function
- Update `defaultExecutablePath` value in `host.json` to match built API (if building API in Linux/MacOS).
- Build handler - In root repo folder, run `go build handler.go`
- Run Azure Functions locally - In root repo folder, run `func start`
- Via Postman/HTTP client of choice, send a GET/POST request to `http://localhost:7071/api/HttpExample`

# Version check
| Tool | Version |
| --- | --- |
| Go | 1.19.1 |
| Azure Functions Core Tools | 4.0.5198 |
| Function Runtime | 4.21.1.20667 |