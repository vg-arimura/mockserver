package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type EndpointInfo struct {
	urlPath  string
	filePath string
}

func main() {
	dataPath := flag.String("data", "./data", "specify response dir")
	port := flag.String("port", "8080", "specify port")
	flag.Parse()

	endpointInfos := makePaths(*dataPath)

	info("start server on port " + *port)
	mux := http.NewServeMux()
	registerEndpoints(mux, endpointInfos)
	http.ListenAndServe(":"+*port, mux)
}

func registerEndpoints(mux *http.ServeMux, endpointInfos []EndpointInfo) {
	for _, endpointInfo := range endpointInfos {
		data, error := ioutil.ReadFile(endpointInfo.filePath)
		if error != nil {
			die("file doesn't exist: %s", endpointInfo.filePath)
		}
		info("register endpoint on " + endpointInfo.urlPath)
		mux.HandleFunc(endpointInfo.urlPath, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", http.DetectContentType(data))
			fmt.Fprintf(w, string(data))
		})
	}
}

func makePaths(dirPath string) []EndpointInfo {
	fi, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		die("not exists: %s", dirPath)
	}
	if !fi.Mode().IsDir() {
		die("not dir: %s", dirPath)
	}

	if strings.HasPrefix(dirPath, "./") {
		dirPath = dirPath[2:]
	}

	var endpointInfos []EndpointInfo
	filepath.Walk(dirPath, func(filePath string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		urlPath := strings.Replace(filePath, dirPath, "", 1)
		endpointInfos = append(endpointInfos, EndpointInfo{
			urlPath,
			filePath,
		})
		return nil
	})

	return endpointInfos
}

func info(v string) {
	log.Println(v)
}

func die(format string, vals ...string) {
	os.Stderr.WriteString(fmt.Sprintf(format, vals) + "\n")
	os.Exit(1)
}
