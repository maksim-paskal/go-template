.PHONY: test
test:
	go fmt ./cmd
	go mod tidy
	SOMEVAR=some-env-value go test --race ./cmd
	golangci-lint run -v
build:
	@./scripts/validate-license.sh
	@./scripts/build-all.sh
	ls -lah _dist
	go mod tidy