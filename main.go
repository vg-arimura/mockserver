package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	dataPath := flag.String("data", "./data", "specify response dir")
	port := flag.String("port", "8080", "specify port")
	flag.Parse()
	paths := makePaths(*dataPath)
	fmt.Println(paths)
	fmt.Println(port)
}

func registerEndpoints(mux *http.ServeMux, paths []string) {
	//print file
	for _, path := range paths {
		data, error := ioutil.ReadFile(path)
		if error != nil {
			die("path doesn't exist: %s", path)
		}
		//		fmt.Println(string(data[:]))
		mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(data)
		})

		//set
	}

	//set contents type
	//register
}

func makePaths(path string) []string {
	fi, err := os.Stat(path)
	if os.IsNotExist(err) {
		die("not exists: %s", path)
	}
	if !fi.Mode().IsDir() {
		die("not dir: %s", path)
	}
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
