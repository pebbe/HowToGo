package p // import "p"

import "C"

//export Foo
func Foo() int32 { return 42 }
