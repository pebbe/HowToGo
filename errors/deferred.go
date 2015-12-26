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
	filename string
	lineno   int
	err      error
}

func main() {

	var errA, errB, errX, errY ErrType
	errors := []*ErrType{&errA, &errB, &errX, &errY} // using this slice causes a compile time error if you forget a variable
	defer func() {
		for _, err := range errors {
			if !err.isNil() {
				fmt.Println(err)
			}
		}
	}()

	a, errA := atoi("2")
	b, errB := atoi("B")
	x, errX := atoi("4")
	y, errY := atoi("Y")
	if any(errors) {
		return
	}
	fmt.Println(a + b)
	fmt.Println(x / y)
}

func atoi(s string) (i int, e ErrType) {
	_, e.filename, e.lineno, _ = runtime.Caller(1)
	i, e.err = strconv.Atoi(s)
	return
}

func any(errors []*ErrType) bool {
	for _, err := range errors {
		if !err.isNil() {
			return true
		}
	}
	return false
}

func (e *ErrType) isNil() bool {
	return (*e).err == nil
}

func (e *ErrType) String() string {
	return fmt.Sprintf("%s:%d: %v", (*e).filename, (*e).lineno, (*e).err)
}
