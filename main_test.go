package main

import (
	_ "fmt"
	"testing"
)

func TestMakePaths(t *testing.T) {
	infos := makePaths("./testdata")
	if infos[0].filePath != "testdata/fuga.json" || infos[0].urlPath != "/fuga.json" {
		t.Fatalf("invalid: %s", infos[0])
	}
	if infos[1].filePath != "testdata/hoge.html" || infos[1].urlPath != "/hoge.html" {
		t.Fatalf("invalid: %s", infos[1])
	}
}
