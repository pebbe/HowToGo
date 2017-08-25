// This is a generated file. Do not edit.

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var err error
	var i int
	i, err = strconv.Atoi("abc")
	if err != nil {
		fmt.Fprintf(os.Stderr, "main.go4:12: error: %v\n", err)
		return
	}
	fmt.Println(i)
}
