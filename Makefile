.PHONY: build-tmp check_swagger_installation swagger createmigration

build-tmp:
	go build -o ./tmp/goblog .

check_swagger_installation:
	which swagger || GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger: check_swagger_installation
	GO111MODULE=off swagger generate spec -o ./docs/swagger.yaml --scan-models

test:
	go test -v -race -cover ./...
