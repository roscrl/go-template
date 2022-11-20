# Go Template

## Setup

`git clone https://github.com/roscrl/go-template.git`  
`go mod download`

`make hotreload` requires `go install github.com/cosmtrek/air@latest`  
`make lint` requires `go install github.com/nametake/golangci-lint-langserver@latest`

## Environment Variables

`export ENV=DEV (default) / UAT / PROD`

#### What changes in each environment?

UAT/PROD: Production zap logger [instead of development](https://pkg.go.dev/go.uber.org/zap#hdr-Configuring_Zap)

## Build

`make run`  
`make hotreload` on file change - rebuild, test & rerun  
`make test`  
`make lint`  
`make build`

## Docker

`docker build . --tag go-template -f config/Dockerfile` image defaults to platform architecture
`docker build . --platform=linux/amd64 --tag go-template -f config/Dockerfile` x86-64
`docker build . --platform=linux/arm64 --tag go-template -f config/Dockerfile` arm64

## Dependencies

`chi` for routering  
`zap` for logging  
`templ` for server side template rendering

`httptest` to spin up mock server responses & http integration requests  
`is` for testing assertions  
`moq` for mocking interfaces, or prefer to [inline test data into structs](https://jrock.us/posts/go-interfaces/)

## Profiling

`curl --output profile "localhost:3000/debug/pprof/profile?seconds=30"` 30 seconds to hit endpoints you are interested in profiling  
`go tool pprof -http localhost:3001 profile`

## Structure Inspiration

[Mat Ryer - How I write HTTP services after eight years talk](https://www.youtube.com/watch?v=XGVZ0Ip4XPM)  
[Mat Ryer - Deep dive of real application](https://www.youtube.com/watch?v=VRZZeJwIAIM)

## Miscellaneous bits & bobs to look into

OpenAPI, OpenTelemetry, Grafana + Prometheus, Cucumber - Godog
