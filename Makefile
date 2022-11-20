.PHONY clean:
clean:
	go clean

.PHONY build:
build: clean 
	go build -o bin/main .

.PHONY run:
run: 
	go run .

.PHONY hotreload:
hotreload: 
	air -c ./config/.air.toml

.PHONY lint:
lint:
	golangci-lint run --config config/.golangci.yml

.PHONY test:
test:
	go test -v 

