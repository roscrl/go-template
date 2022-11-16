# Go Template

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
`clock` for mocking time
`templ` for server side template rendering

`httptest` to spin up mock server responses & http integration requests
`is` for testing assertions
`moq` for mocking

## Structure Inspiration

[Mat Ryer - How I write HTTP services after eight years talk](https://www.youtube.com/watch?v=XGVZ0Ip4XPM)  
[Mat Ryer - Deep dive of real application](https://www.youtube.com/watch?v=VRZZeJwIAIM)
