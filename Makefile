run:
	@go run cmd/main.go

BINARY_NAME=Nefer_UI
VERSION?=0.0.1
BUILD_DIR=releases
MAIN_PATH=./cmd/

all: build-all

clean:
	rm -rf $(BUILD_DIR)

$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

# Build for all platforms
build-all: windows linux darwin

# Build for Windows
windows: $(BUILD_DIR)
	GOOS=windows go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-windows $(MAIN_PATH)

# Build for Linux
linux: $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 $(MAIN_PATH)
	GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-arm64 $(MAIN_PATH)

# Build for macOS
darwin: $(BUILD_DIR)
	GOOS=darwin go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-macos $(MAIN_PATH)

# Build for current platform
build: $(BUILD_DIR)
	go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)