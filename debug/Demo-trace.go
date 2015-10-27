package main

import (
	"fmt"
	"log"
	"time"
)

var (
	tracecount = 0
)

func main() {
	defer trace("main()")()

	time.Sleep(time.Second)
	process("test")
	time.Sleep(time.Second)
}

func process(arg string) {
	defer trace("process(%q)", arg)()

	time.Sleep(time.Second)
}

func trace(format string, args ...interface{}) func() {
	start := time.Now()
	tracecount++
	n := tracecount
	s := fmt.Sprintf(format, args...)
	log.Printf("ENTER [%d] %v", n, s)
	return func() {
		log.Printf("EXIT  [%d] %v  %v", n, s, time.Since(start))
	}
}
