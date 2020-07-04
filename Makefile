lint-install:
	@curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin

lint:
	golangci-lint run

test:
	go clean -testcache
	go test ./... --race

test-coverage:
	go test ./... -covermode=atomic -coverprofile=coverage.out -coverpkg ./...
	go tool cover -func=coverage.out

test-coverage-html:
	go test ./... -covermode=atomic -coverprofile=coverage.out -coverpkg ./...
	go tool cover -html=coverage.out

grpc-api:
	protoc --proto_path=proto --go_out=plugins=grpc:. iticapital.proto

generate-event-handler:
	go run handlergen.go