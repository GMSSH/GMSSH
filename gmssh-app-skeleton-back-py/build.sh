#!/bin/bash
set -euo pipefail
IFS=$'\n\t'

# 设置 PATH
PATH=/bin:/sbin:/usr/bin:/usr/sbin:/usr/local/bin:/usr/local/sbin:~/bin
export PATH

# 获取当前脚本所在目录
CURDIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

## 参数配置
VENV_DIR="$(mktemp -d -t example_venv_XXXX)"
ENTRY_FILE="main.py"
DIST_NAME="main"
FINAL_DIST_DIR="$CURDIR/../dist"

# 清理函数
function cleanup {
    echo "🧹 正在清理临时虚拟环境..."
    rm -rf "$VENV_DIR"
}
trap cleanup EXIT

# 创建虚拟环境
echo "🚀 创建虚拟环境：$VENV_DIR"
python3 -m venv "$VENV_DIR"
source "$VENV_DIR/bin/activate"

# 安装依赖（如果 requirements.txt 存在）
if [[ -f "$CURDIR/requirements.txt" ]]; then
    echo "📦 安装依赖..."
    "$VENV_DIR/bin/pip" install -r "$CURDIR/requirements.txt"
else
    echo "⚠️  未找到 requirements.txt，跳过依赖安装。"
fi

# 编译
echo "⚙️  使用 Nuitka 编译 $ENTRY_FILE..."
python3 -m nuitka \
    --follow-imports \
    --standalone \
    --show-progress \
    --include-data-files="$CURDIR/config.yaml=config.yaml" \
    --include-data-dir="$CURDIR/app/i18n=app/i18n" \
    "$CURDIR/$ENTRY_FILE"

# 检查是否存在目标编译产物
if [[ ! -f "${DIST_NAME}.dist/${DIST_NAME}.bin" ]]; then
    echo "❌ 编译失败，未找到产物 ${DIST_NAME}.bin"
    exit 1
fi

# 替换为无后缀二进制文件
echo "📦 重命名编译产物..."
rm -rf "${DIST_NAME}.build"
mv "${DIST_NAME}.dist/${DIST_NAME}.bin" "${DIST_NAME}.dist/${DIST_NAME}"

# 移动至最终目录
echo "📁 拷贝编译产物到 $FINAL_DIST_DIR"
mkdir -p "$FINAL_DIST_DIR"
mv "${DIST_NAME}.dist" "$FINAL_DIST_DIR/backend/bin"

echo "✅ 编译完成：$FINAL_DIST_DIR/${DIST_NAME}.dist/${DIST_NAME}"