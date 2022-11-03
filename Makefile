.PHONY clean:
clean:
	go clean

.PHONY build:
build: clean 
	go build -o bin/app .

.PHONY run:
run: 
	go run .

.PHONY hotreload:
hotreload: 
	air -c ./config/.air.toml

.PHONY lint:
lint:
	golangci-lint run

.PHONY test:
test:
	go test -v 
