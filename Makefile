.PHONY: help run-concepts run-concept run-practice run-practice-demo fmt vet build clean

CATEGORY ?= basics
DEMO ?=
PRACTICE_CATEGORY ?= concurrency
PRACTICE_DEMO ?=

help:
	@echo "Usage:"
	@echo "  make run-concepts   Show the concepts category menu"
	@echo "  make run-concept    Run a concept demo"
	@echo "                      e.g. make run-concept CATEGORY=concurrency DEMO=select"
	@echo "  make run-practice   Show the practice category menu"
	@echo "  make run-practice-demo"
	@echo "                      Run a practice demo"
	@echo "                      e.g. make run-practice-demo PRACTICE_CATEGORY=concurrency PRACTICE_DEMO=select"
	@echo "  make fmt            Format Go sources"
	@echo "  make vet            Run go vet on both modules"
	@echo "  make build          Build all packages and binaries"
	@echo "  make clean          Remove build artifacts"

run-concepts:
	cd src/Concepts && go run .

run-concept:
	cd src/Concepts && go run ./$(CATEGORY) $(DEMO)

run-practice:
	cd src/Practice && go run .

run-practice-demo:
	cd src/Practice && go run ./$(PRACTICE_CATEGORY) $(PRACTICE_DEMO)

fmt:
	gofmt -w src/Concepts src/Practice

vet:
	cd src/Concepts && go vet ./...
	cd src/Practice && go vet ./...

build:
	mkdir -p bin
	cd src/Concepts && go build ./...
	cd src/Concepts && go build -o ../../bin/concepts .
	cd src/Practice && go build -o ../../bin/practice .

clean:
	rm -rf bin
