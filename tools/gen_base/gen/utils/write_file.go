package utils

import (
	"fmt"
	"os"
)

var (
	ErrWriteFile = fmt.Errorf("failed to write file")
	ErrOpenFile  = fmt.Errorf("failed to open file for writing")
)

// WriteFile writes content to a file with the given path.
func WriteFile(file_path, content string) error {
	err := CreateFile(file_path)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(file_path, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.WriteString(content); err != nil {
		return err
	}

	return nil
}
