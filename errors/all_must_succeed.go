package main

import (
	"fmt"
	"log"
	"runtime"
	"strconv"
)

type Err struct {
	err error
}

func (e *Err) ok(err error) bool {
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

	var err Err

	_ = err.ok(readint("123", &a)) &&
		err.ok(readint("456", &b)) &&
		err.ok(readint("4fd", &c)) &&
		err.ok(readint("qas", &d)) &&
		err.ok(readint("boo", &e))

	if err.err != nil {
		log.Println(err.err)
	}

	fmt.Println(a, b, c, d, e)

}

func readint(s string, i *int) error {
	var err error
	*i, err = strconv.Atoi(s)
	return err
}
