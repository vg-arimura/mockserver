package main

import (
	_ "fmt"
	"testing"
)

func TestLoad(t *testing.T) {
	paths := makePaths("./testdata")
	if paths[0] != "testdata/fuga.json" {
		t.Fatalf("invalid path: %s", paths[0])
	}
	if paths[1] != "testdata/hoge.html" {
		t.Fatalf("invalid path: %s", paths[1])
	}
}

func TestRegisterEndpoints(t *testing.T) {
	registerEndpoints(nil, []string{"testdata/fuga.json", "testdata/hoge.html"})
}
