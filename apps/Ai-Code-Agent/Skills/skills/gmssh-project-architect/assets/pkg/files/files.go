package files

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Exists 判断文件或目录是否存在
func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// IsDir 判断路径是否为目录
func IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// IsFile 判断路径是否为文件
func IsFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// ReadFile 读取文件内容
func ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

// WriteFile 写入文件内容
func WriteFile(path string, data []byte, perm os.FileMode) error {
	// 确保父目录存在
	dir := filepath.Dir(path)
	if !Exists(dir) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}
	return os.WriteFile(path, data, perm)
}

// ReadJSON 读取JSON文件并解析到结构体
func ReadJSON(path string, v any) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

// WriteJSON 将结构体序列化为JSON并写入文件
func WriteJSON(path string, v any, indent bool) error {
	var data []byte
	var err error

	if indent {
		data, err = json.MarshalIndent(v, "", "  ")
	} else {
		data, err = json.Marshal(v)
	}

	if err != nil {
		return err
	}

	return WriteFile(path, data, 0644)
}

// EnsureDir 确保目录存在,不存在则创建
func EnsureDir(path string) error {
	if !Exists(path) {
		return os.MkdirAll(path, 0755)
	}
	return nil
}

// Remove 删除文件或目录
func Remove(path string) error {
	return os.RemoveAll(path)
}

// Copy 复制文件
func Copy(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	return WriteFile(dst, data, 0644)
}

// ListDir 列出目录下的所有文件和子目录
func ListDir(path string) ([]os.FileInfo, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return f.Readdir(-1)
}

// GetFileSize 获取文件大小(字节)
func GetFileSize(path string) (int64, error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}
