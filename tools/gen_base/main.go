package main

import (
	"fmt"
	"genbase/gen/genfiles"
	"os"
	"path/filepath"
	"runtime"
)

func init() {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Printf("failed to determine current file path")
	}
	dir := filepath.Dir(file)
	if err := os.Chdir(dir); err != nil {
		fmt.Printf("failed to change working directory: %v", err)
	}
}

func main() {
	genfiles.GenClient()
	genfiles.GenCmd()
	genfiles.GenDig()
	genfiles.GenInternal()
}
