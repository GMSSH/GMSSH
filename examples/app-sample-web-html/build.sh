#!/bin/bash
# 获取当前脚本所在目录
current_dir=$(cd $(dirname $0); pwd)

# 定义源文件路径（index.html）
source_file="$current_dir/index.html"

# 定义目标目录路径（上一层级的 dist/web）
target_dir="$current_dir/../dist/web"

# 检查源文件是否存在
if [ -f "$source_file" ]; then
  # 如果目标目录不存在，则创建
  if [ ! -d "$target_dir" ]; then
    mkdir -p "$target_dir"
  fi

  # 复制文件到目标目录
  cp "$source_file" "$target_dir/"

  echo "文件已成功复制到: $target_dir"
else
  echo "错误: 源文件 $source_file 不存在!"
  exit 1
fi