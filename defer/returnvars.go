/*

Over het resultaat van testb():

If you look a the spec [http://golang.org/ref/spec#Return_statements]:
        A "return" statement that specifies results sets the result parameters before any deferred functions are executed.

In other words, your return statement sets a to its own value, which is 1,
and then the deferred function runs and sets a to 2.
It would be the same if you wrote 'return 5' instead of 'return a'.

*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(testa()) // 1
	fmt.Println(testb()) // 2
	fmt.Println(testc()) // 2
}

func testa() int {
	a := 1

	defer func() {
		a = 2
	}()

	return a
}

func testb() (a int) {
	defer func() {
		a = 2
	}()

	a = 1
	return a
}

func testc() (a int) {
	defer func() {
		a = 2
	}()

	a = 1
	return
}
