package main

import (
	"fmt"

	"github.com/alecthomas/kong"
)

var CLI struct {
	One struct {
		OneArg string `arg:"" required:""`
	} `cmd: ""`
	Two struct {
		TwoArg string `arg:"" enum:"a,b,c" required:""`
	} `cmd: ""`
}

func main() {
	ctx := kong.Parse(&CLI)
	switch cmd := ctx.Command(); cmd {
	case "one <one-arg>", "two <two-arg>":
		fmt.Println("command:", cmd)
	default:
		panic(ctx.Command())
	}
}
