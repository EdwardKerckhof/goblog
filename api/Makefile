.PHONY: build-tmp check_swagger_installation swagger createmigration

build-tmp:
	go build -o ./tmp/goblog ./cmd/httpserver

check_swagger_installation:
	which swagger || GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger: check_swagger_installation
	GO111MODULE=off swagger generate spec -o ./docs/swagger.yaml --scan-models

test:
	go test -v -race -cover ./...

mock-repos:
	mockgen -source=internal/core/ports/repositories.go -destination=mocks/repositories/mock.go -package=repositories_mock
