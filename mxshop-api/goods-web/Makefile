.PHONY: init
# init env
init:
	go mod tidy

.PHONY: test
# test
test:
	go test -v -cover ./...

.PHONY: build
# build
build:
	go build  -o ./bin/main ./cmd/

.PHONY: run
# run
run:
	go run  ./cmd/...