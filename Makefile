.PHONY: build test vet lint clean install

BIN := certcheck

build:
	go build -trimpath -o $(BIN) ./...

test:
	go test -race ./...

vet:
	go vet ./...

install:
	go install ./...

clean:
	rm -f $(BIN)
