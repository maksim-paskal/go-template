.PHONY: test
test:
	./scripts/validate-license.sh
	go fmt ./cmd
	go mod tidy
	SOMEVAR=some-env-value go test --race ./cmd
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run -v
build:
	make test
	@./scripts/build-all.sh
	ls -lah _dist
	go mod tidy
run:
	go run --race ./cmd $(args)