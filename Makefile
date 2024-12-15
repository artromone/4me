.PHONY: build run test docker-build docker-run clean

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod

# Binary names and directories
BINDIR=bin
SERVER_BINARY=$(BINDIR)/task-server
CLI_BINARY=$(BINDIR)/task-cli

# Build targets
build: $(BINDIR)
	$(GOBUILD) -o $(SERVER_BINARY) ./cmd/server/main.go
	$(GOBUILD) -o $(CLI_BINARY) ./cmd/cli/main.go

$(BINDIR):
	mkdir -p $(BINDIR)

run: build
	./$(SERVER_BINARY)

test:
	$(GOTEST) ./...

# Docker targets
docker-build:
	docker build -t task-management .

docker-run: docker-build
	docker run -p 8080:8080 task-management

# Dependency management
deps:
	$(GOMOD) download
	$(GOMOD) tidy

# Clean up
clean:
	rm -rf $(BINDIR)
	docker rmi -f task-management
