.PHONY: test
test:
	@./scripts/test.sh
build:
	@./scripts/validate-license.sh
	@./scripts/build-all.sh