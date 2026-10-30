.PHONY: help run-concepts run-practice fmt vet build clean

help:
	@echo "Usage:"
	@echo "  make run-concepts   Run the concepts demos"
	@echo "  make run-practice   Run the practice module"
	@echo "  make fmt            Format Go sources"
	@echo "  make vet            Run go vet on both modules"
	@echo "  make build          Build both modules into bin/"
	@echo "  make clean          Remove build artifacts"

run-concepts:
	cd src/Concepts && go run .

run-practice:
	cd src/Practice && go run .

fmt:
	gofmt -w src/Concepts src/Practice

vet:
	cd src/Concepts && go vet ./...
	cd src/Practice && go vet ./...

build:
	mkdir -p bin
	cd src/Concepts && go build -o ../../bin/concepts .
	cd src/Practice && go build -o ../../bin/practice .

clean:
	rm -rf bin
