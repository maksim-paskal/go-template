.PHONY: test
test:
	go fmt ./cmd
	go mod tidy
	go test ./cmd
	golangci-lint run --allow-parallel-runners -v --enable-all --disable testpackage --fix
build:
	@./scripts/validate-license.sh
	@./scripts/build-all.sh
	ls -lah _dist
	go mod tidy