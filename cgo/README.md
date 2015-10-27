## cgo practices

```go
// make sure the two C calls run without interference from another go routine
// 'mu' is a global (package) variable of type sync.Mutex
mu.Lock()
foo := C.foo_new()    // sets global C variable
err := C.foo_err()    // reads global C variable
mu.Unlock()

// free C resources before foo is garbage collected
runtime.SetFinalizer(foo, (*Foo).Free)
```

C struct with or without typedef

```go
package main

/*
typedef struct { int a, b; } tst1;
struct tst2 { int a, b; };
*/
import "C"

import "fmt"

func main() {
    var tt1 C.tst1
    tt1.a = C.int(1)
    var tt2 C.struct_tst2
    tt2.b = C.int(1)
    fmt.Printf("%#v\n%#v\n", tt1, tt2)
    /*
    Output:
    main._Ctype_tst1{a:1, b:0}
    main._Ctype_struct_tst2{a:0, b:1}
    */
}
```

Put non-standard path to libraries into binary:

```sh
go build -ldflags="-r /path/to/lib1:/path/to/lib2" program.go
```
