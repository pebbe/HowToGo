package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	// requires go1.12 or newer
	if bi, ok := debug.ReadBuildInfo(); ok {
		// go1.18 and newer: latest git tag
		fmt.Println(bi.Main.Version)
	} else {
		fmt.Println("No build info")
	}
}
