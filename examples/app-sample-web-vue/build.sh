#!/bin/bash

# 获取当前脚本所在目录
current_dir=$(cd "$(dirname "$0")"; pwd)

# 进入项目目录 (假设打包脚本位于项目根目录下)
cd "$current_dir" || exit

# 安装依赖
yarn install

# 执行打包命令
yarn run build

# 创建输出目录
output_dir="../dist/web"
mkdir -p "$output_dir"

# 将打包后的内容复制到输出目录
cp -r dist/* "$output_dir/"

echo "打包完成，内容已输出到 $output_dir"