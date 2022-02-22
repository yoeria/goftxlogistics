package main

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/gookit/color"
)

func TestReq(t *testing.T) {
	fmt.Printf("t: %v\n", t)
	color.Set(color.HiGreen, color.BgBlack)
	spew.Dump(request)
	color.Reset()
	fmt.Println()
}

func TestChart(t *testing.T) {
	fmt.Print("TestChart func:\n\n")
	color.Set(color.Magenta)
	//spew.Dump(t)
	color.Reset()
	serveChart()
}
