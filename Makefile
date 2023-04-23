PACKAGES := $(shell go list ./... | grep -v /vendor/)

install-tools:
	cat tools/tools.go | grep "_" | awk -F '"' '{print $$2}' | xargs -L1 go install

tools:
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

vendor:
	go mod vendor

clean:
	rm -rf vendor docs/swagger

install: tools
	go mod vendor && go mod tidy

lint:
	golangci-lint run -n

test/cover:
	@echo "mode: count" > coverage-all.out
	@$(foreach pkg,$(PACKAGES), \
		go test -p=1 -cover -covermode=count -coverprofile=coverage.out ${pkg}; \
		tail -n +2 coverage.out >> coverage-all.out;)
	go tool cover -html=coverage-all.out

run/api:
	go run main.go api

swagger:
	swag init -g ./api/api.go -o ./docs/swagger