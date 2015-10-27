/*

Printing make-like warnings and errors

*/

package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func main() {

	a := "foo"

	i1, err := strconv.Atoi(a)
	if x(err) {
		fmt.Println("That didn't go well")
	}

	i2, err := strconv.Atoi(a)
	if x(err, "here is some info", 1, 2, 3) {
		fmt.Println("That didn't go well at all")
	}

	i3, err := strconv.Atoi(a)
	xx(err)

	fmt.Println("You can't see us:", i1, i2, i3)

}

func x(err error, a ...interface{}) bool {

	if err == nil {
		return false
	}

	s := err.Error()
	_, filename, lineno, ok := runtime.Caller(1)
	if ok {
		s = fmt.Sprintf("%s:%d: warning: %s", filename, lineno, s)
	} else {
		s = "warning: " + s
	}
	if len(a) > 0 {
		s = s + " | " + strings.TrimSpace(fmt.Sprintln(a...))
	}

	fmt.Println(s)

	return true

}

func xx(err error, a ...interface{}) {

	if err == nil {
		return
	}

	s := err.Error()
	_, filename, lineno, ok := runtime.Caller(1)
	if ok {
		s = fmt.Sprintf("%s:%d: error: %s", filename, lineno, s)
	} else {
		s = "error: " + s
	}
	if len(a) > 0 {
		s = s + " | " + strings.TrimSpace(fmt.Sprintln(a...))
	}

	fmt.Println(s)
	os.Exit(1)
}
