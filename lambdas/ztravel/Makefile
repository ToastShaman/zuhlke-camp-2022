BUILD_DIR=build
BUNDLING_DIR=/asset-output

all: clean test build

test:
	go test ./...
.PHONY: test

build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/hello cmd/api/hello/hello.go
	go build -o $(BUILD_DIR)/world cmd/api/world/world.go
	go build -o $(BUILD_DIR)/pigtail cmd/api/pigtail/pigtail.go
	chmod u+x $(BUILD_DIR)/*
.PHONY: build

clean:
	rm -rf $(BUILD_DIR)
.PHONY: clean

bundling: clean test build
	mkdir -p $(BUNDLING_DIR) && mv $(BUILD_DIR)/* $(BUNDLING_DIR)
.PHONY: bundling
