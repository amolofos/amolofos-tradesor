.PHONY: init clean fmt lint build test run

init:
	./scripts/init.sh

clean:
	./scripts/clean.sh

fmt:
	./scripts/fmt.sh

lint:
	./scripts/lint.sh

build:
	./scripts/build.sh

test:
	./scripts/test.sh

run:
	./scripts/run.sh
