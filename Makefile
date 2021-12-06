.PHONY: test
test:
	./scripts/validate-license.sh
	go fmt ./cmd
	go mod tidy
	SOMEVAR=some-env-value go test --race ./cmd
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run -v
build:
	go run github.com/goreleaser/goreleaser@latest build --rm-dist --skip-validate
run:
	go run --race ./cmd $(args)