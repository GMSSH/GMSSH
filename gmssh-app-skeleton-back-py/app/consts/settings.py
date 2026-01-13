"""
@文件        :__init__.py
@说明        :This is an example
@时间        :2025/06/30 09:17:23
@作者        :xxx
@邮箱        :
@版本        :1.0.0
"""
import os

def mkdir_dir(path):
    if not os.path.exists(path):
        os.makedirs(path)
    return path

# 方式 1：使用当前文件所在目录进行拼接
CURRENT_PATH = os.path.dirname(__file__)
APP_TMP_DIR_PATH = mkdir_dir(os.path.join(os.path.dirname(CURRENT_PATH), "tmp"))
APP_CONFIG_FILE_PATH = os.path.join(CURRENT_PATH, "config.yaml")
APP_I18N_DIR_PATH = os.path.join(os.path.dirname(CURRENT_PATH), "i18n")
APP_SOCKET_FILE_PATH = os.path.join(APP_TMP_DIR_PATH, "app.sock")


# 方式 2：使用绝对路径进行拼接
# 应用名称：{组织名}/{应用名}，需要根据实际情况修改
APP_NAME = "official/example"

GMSSH_PATH = "/.__gmssh"
PLUGIN_PATH = os.path.join(GMSSH_PATH, "plugin")
APP_INSTALLED_PATH = os.path.join(PLUGIN_PATH, APP_NAME)

APP_TMP_PATH = os.path.join(APP_INSTALLED_PATH, "tmp")
APP_DIR_PATH = os.path.join(APP_INSTALLED_PATH, "app")
APP_BIN_PATH = os.path.join(APP_DIR_PATH, "bin")

# 在应用上架的时候请使用下面路径更加稳妥
# APP_SOCKET_FILE_PATH = os.path.join(APP_TMP_PATH, "app.sock")
# APP_CONFIG_FILE_PATH = os.path.join(APP_BIN_PATH, "config.yaml")
# APP_I18N_DIR_PATH = os.path.join(APP_BIN_PATH, "app/i18n")

