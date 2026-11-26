# 项目构建Makefile

# 定义构建目标目录
DIST_DIR := dist

# 构建后端应用
build:
	@echo "开始构建后端应用..."
	@chmod +x build.sh
	@./build.sh
	@echo "后端应用构建完成!"

# 清理构建结果
clean:
	@echo "清理构建结果..."
	@rm -rf $(DIST_DIR)
	@echo "清理完成!"

# 帮助信息
help:
	@echo "可用命令:"
	@echo "  make build - 构建后端应用"
	@echo "  make clean     - 清理构建结果"
	@echo "  make help      - 显示此帮助信息"

.PHONY: build clean help
