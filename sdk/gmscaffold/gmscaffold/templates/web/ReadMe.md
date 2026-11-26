# {{ description }}


## FAQ
#### 处理多进程编译的问题
"""python3
import multiprocessing

# 处理多进程编译的问题
# freeze_support: 主脚本被二次执行时进程会退出
multiprocessing.freeze_support()
"""
