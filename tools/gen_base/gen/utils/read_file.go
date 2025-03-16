package utils

import (
	"fmt"
	"os"
)

var (
	ErrReadFile = fmt.Errorf("failed to read file")
)

// ReadFile reads the content of a file.
func ReadFile(file_path string) (string, error) {
	file, err := os.Open(file_path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return "", err
	}

	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)
	_, err = file.Read(buffer)
	if err != nil {
		return "", err
	}

	return string(buffer), nil
}
