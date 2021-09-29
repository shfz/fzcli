run:
	go run main.go run -t ../demo-typescript/index.js -o log/ -p 5 -n 20

cli:
	go run main.go run -t ../demo-typescript/index.js -o log/ -p 5 -n 20 -v off

build:
	make pre
	go build

fmt:
	gofmt -s -w .

lint:
	golangci-lint run --tests --enable=golint

test:
	-go test -v -race -cover ./...

pre:
	@make fmt
	@make lint
	@make test

update:
	rm go.sum
	go get -u
	go mod tidy
	@make pre
