.PHONY: build test lint clean install generate fmt vet tidy help

BINARY   := forge
MAIN     := .
BIN_DIR  := bin
VERSION  := 0.1.0

# ── Build ─────────────────────────────────────────────────────────────────────

build:
	@mkdir -p $(BIN_DIR)
	go build -ldflags="-s -w -X github.com/miladmahmoodi/forge/cmd.Version=$(VERSION)" \
		-o $(BIN_DIR)/$(BINARY) $(MAIN)
	@echo "  built  $(BIN_DIR)/$(BINARY)"

install:
	go install -ldflags="-s -w -X github.com/miladmahmoodi/forge/cmd.Version=$(VERSION)" $(MAIN)
	@echo "  installed forge"

# ── Generate ──────────────────────────────────────────────────────────────────

generate: build
	./$(BIN_DIR)/$(BINARY) build
	@echo "  README.md regenerated"

# ── Quality ───────────────────────────────────────────────────────────────────

test:
	go test -race -coverprofile=coverage.out ./...
	@go tool cover -func=coverage.out | tail -1

vet:
	go vet ./...

fmt:
	gofmt -w -s .

lint:
	golangci-lint run ./...

tidy:
	go mod tidy

# ── Clean ─────────────────────────────────────────────────────────────────────

clean:
	rm -rf $(BIN_DIR) coverage.out

# ── Help ──────────────────────────────────────────────────────────────────────

help:
	@echo ""
	@echo "  Forge — $(VERSION)"
	@echo ""
	@echo "  make build       Build the forge binary"
	@echo "  make install     Install forge to GOPATH/bin"
	@echo "  make generate    Build binary and regenerate README.md"
	@echo "  make test        Run tests with race detector"
	@echo "  make lint        Run golangci-lint"
	@echo "  make fmt         Format source code"
	@echo "  make vet         Run go vet"
	@echo "  make tidy        Run go mod tidy"
	@echo "  make clean       Remove build artifacts"
	@echo ""
