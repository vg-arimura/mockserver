package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func registerEndpoints(mux *http.ServeMux, paths []string) {
	//print file
	for _, path := range paths {
		data, error := ioutil.ReadFile(path)
		if error != nil {
			die("path doesn't exist: %s", path)
		}
		fmt.Println(string(data[:]))
	}

	//set contents type
	//register
}

func makePaths(path string) []string {
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

func die(format string, vals ...string) {
	os.Stderr.WriteString(fmt.Sprintf(format, vals) + "\n")
	os.Exit(1)
}
