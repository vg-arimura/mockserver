package main

import (
	_ "fmt"
	"os"
	"path/filePath"
)

type Endpoint struct {
	urlPath  string
	filePath string
}

func load(path string) []string {
	var paths []string
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		paths = append(paths, path)
		return nil
	})

	return paths
}
