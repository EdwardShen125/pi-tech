# 定义目录变量
API_DIR := .
PROTO_DIR := ../proto
GEN_DIR := ../proto-gen-go

# 为确保输出结构与 api 中一致，我们将处理相对路径

# Step 1: 根据 .api 生成 .proto 文件
.PHONY: gen-proto
gen-proto:
	@find $(API_DIR) -type f -name "*.api" | while read api; do \
		rel=`dirname "$$api" | sed 's#^./##'`; \
		if [ "$$rel" = "." ]; then \
			outdir="$(PROTO_DIR)"; \
		else \
			outdir="$(PROTO_DIR)/$$rel"; \
		fi; \
		mkdir -p $$outdir; \
		echo "Generating proto for $$api -> $$outdir"; \
		goctl-proto proto --input "$$api" --output "$$outdir"; \
	done


# Step 2: 根据 proto 生成 pb 文件（包含 client 和 server）
.PHONY: gen-pb
gen-pb:
	@find $(PROTO_DIR) -type f -name "*.proto" | while read proto; do \
		echo "Generating pb for $$proto -> $(GEN_DIR)"; \
		protoc --proto_path=$(PROTO_DIR) --go_out=$(GEN_DIR) --go_opt=paths=source_relative \
		       --go-grpc_out=$(GEN_DIR) --go-grpc_opt=paths=source_relative \
		       "$$proto"; \
	done
#.PHONY: gen-pb
#gen-pb:
#	@find $(PROTO_DIR) -type f -name "*.proto" | while read proto; do \
#	    rel=`dirname "$$proto" | sed "s#$(PROTO_DIR)/##"`; \
#	    base=`basename "$$proto"`; \
#	    if [ "$$rel" = "" ]; then \
#	        outdir="$(GEN_DIR)"; \
#	        pkg="default"; \
#	    else \
#	        outdir="$(GEN_DIR)/$$rel"; \
#	        pkg=`basename "$$rel"`; \
#	    fi; \
#	    mkdir -p $$outdir; \
#	    echo "Generating pb for $$proto -> $$outdir with package $$pkg"; \
#	    protoc --proto_path=$(PROTO_DIR) \
#	    	   --go_out=$(GEN_DIR) \
#	    	   --go_opt=M$$base=$$pkg,paths=source_relative \
#	    	   --go-grpc_out=$(GEN_DIR) \
#	           --go-grpc_opt=M$$base=$$pkg,paths=source_relative \
#	           "$$proto"; \
#	done



# Step 3: 在 proto-gen-go 目录下执行 go mod tidy
.PHONY: tidy
tidy:
	@echo "Running go mod tidy in $(GEN_DIR)..."
	@cd $(GEN_DIR) && go mod tidy

.PHONY: all
all: gen-proto gen-pb tidy
