# Variables
PROTO_DIR := $(SERVICE)
OUT_DIR := $(PROTO_DIR)/proto
PROTO_FILES := $(wildcard $(PROTO_DIR)/*.proto)

# Check if SERVICE is provided
check-service:
	@if [ -z "$(SERVICE)" ]; then \
		echo "Error: SERVICE is not specified. Use 'make <target> SERVICE=<service_name>'."; \
		exit 1; \
	fi

# Generate gRPC code
.PHONY: proto
proto: check-service
	@echo "Generating proto files for service: $(SERVICE)"
	@mkdir -p $(OUT_DIR)
	protoc --proto_path=$(PROTO_DIR) \
	       --go_out=$(OUT_DIR) \
	       --go-grpc_out=$(OUT_DIR) \
	       $(PROTO_FILES)

# Clean generated files
.PHONY: clean
clean: check-service
	@echo "Cleaning proto files for service: $(SERVICE)"
	@rm -rf $(OUT_DIR)
