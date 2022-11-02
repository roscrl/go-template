# Go Scaffold

## Setup

Hot reload, on file change rebuild app

`go install github.com/cosmtrek/air@latest`

Linting

`go install github.com/nametake/golangci-lint-langserver@latest`

## Build

`make run`  
`make hotreload`  
`make test`  
`make lint`  
`make build`

## Dependencies

`chi` for routering  
`zap` for logging  
`templ` for server side template rendering

`is` for testing  
`moq` for mocking

## Structure Inspiration

[Mat Ryer - How I write HTTP services after eight years talk](https://www.youtube.com/watch?v=XGVZ0Ip4XPM)
[Mat Ryer - Deep dive of real application](https://www.youtube.com/watch?v=VRZZeJwIAIM)
