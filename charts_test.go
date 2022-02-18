package main

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/gookit/color"
)

func TestReq(t *testing.T) {
	fmt.Printf("t: %v\n", t)
	//fmt.Println(request)
	color.Set(color.Magenta)
	spew.Dump(request)
	color.Reset()
	fmt.Println()
}
