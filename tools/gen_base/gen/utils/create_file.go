package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	// 创建失败
	ErrCreateFile = fmt.Errorf("failed to create file")
)

// CreateFile creates a file with the given path and content.
func CreateFile(file_path string) error {
	dir := filepath.Dir(file_path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return ErrCreateFile
		}
	}
	// 创建文件
	_, err := os.Create(file_path)
	if err != nil {
		return err
	}
	return nil
}
