/*

Defer error handling until you got all the error values

*/

package main

import (
	"fmt"
	"runtime"
	"strconv"
)

type ErrType struct {
	err      error
	filename string
	lineno   int
}

func main() {

	var err0, err1, err2, err3 ErrType
	errors := []*ErrType{&err0, &err1, &err2, &err3}
	defer func() {
		for _, err := range errors {
			if (*err).err != nil {
				fmt.Printf("%s:%d: %v\n", (*err).filename, (*err).lineno, (*err).err)
			}
		}
	}()

	a, err0 := atoi("A")
	b, err1 := atoi("B")
	x, err2 := atoi("X")
	y, err3 := atoi("Y")
	if any(errors) {
		return
	}
	fmt.Println(a + b)
	fmt.Println(x / y)
}

func atoi(s string) (int, ErrType) {
	var e ErrType
	i, err := strconv.Atoi(s)
	if err == nil {
		return i, e
	}
	e.err = err
	_, e.filename, e.lineno, _ = runtime.Caller(1)
	return 0, e
}

func any(errors []*ErrType) bool {
	for _, err := range errors {
		if (*err).err != nil {
			return true
		}
	}
	return false
}
