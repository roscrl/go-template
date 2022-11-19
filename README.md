# Go Template

## Setup

`make hotreload` requires `go install github.com/cosmtrek/air@latest`  
`make lint` requires `go install github.com/nametake/golangci-lint-langserver@latest`

## Build

`make run`  
`make hotreload` on file change - rebuild, test & rerun  
`make test`  
`make lint`  
`make build`

## Docker

`docker build . --tag go-template -f config/Dockerfile` image defaults to platform architecture
`docker build . --platform=linux/amd64 --tag go-template -f config/Dockerfile` x86-64

## Dependencies

`chi` for routering  
`zap` for logging  
`templ` for server side template rendering

`httptest` to spin up mock server responses & http integration requests  
`is` for testing assertions  
`moq` for mocking

## Structure Inspiration

[Mat Ryer - How I write HTTP services after eight years talk](https://www.youtube.com/watch?v=XGVZ0Ip4XPM)  
[Mat Ryer - Deep dive of real application](https://www.youtube.com/watch?v=VRZZeJwIAIM)
