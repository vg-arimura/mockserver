package main

import (
	_ "fmt"
	"testing"
)

func TestDummy(t *testing.T) {
	fmt.Println(load("./data"))
}
