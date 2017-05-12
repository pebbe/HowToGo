package main

import (
	"fmt"
	"log"
	"runtime"
	"strconv"
)

type Error struct {
	err error
}

func (e *Error) ok(err error) bool {
	if err == nil {
		return true
	}
	_, filename, lineno, ok := runtime.Caller(1)
	if ok {
		e.err = fmt.Errorf("%s -- %s:%d", err, filename, lineno)
	} else {
		e.err = err
	}
	return false
}

func main() {

	var a, b, c, d, e int

	var x Error

	_ = true &&
		x.ok(readint("123", &a)) &&
		x.ok(readint("456", &b)) &&
		x.ok(readint("4fd", &c)) &&
		x.ok(readint("qas", &d)) &&
		x.ok(readint("boo", &e))

	if x.err != nil {
		log.Println(x.err)
	}

	fmt.Println(a, b, c, d, e)

}

func readint(s string, i *int) error {
	var err error
	*i, err = strconv.Atoi(s)
	return err
}
