.PHONY: build test vet lint clean install cover bench

BIN := certcheck

build:
	go build -trimpath -o $(BIN) ./...

test:
	go test -race ./...

vet:
	go vet ./...

cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

bench:
	go test -bench=. -benchmem -run=^$$ ./...

install:
	go install ./...

clean:
	rm -f $(BIN) coverage.out
