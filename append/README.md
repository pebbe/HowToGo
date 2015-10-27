## append behavior

see: (message by Jan Mercl)[https://plus.google.com/u/0/100144763948435845718/posts/cRuK8cJhEU4]

```go
package main

import "fmt"

func main() {
    a := make([]int, 1, 100)
    a[0] = 10
    fmt.Println(a)

    fmt.Println()

    b := append(a, 20)
    fmt.Println(a)
    fmt.Println(b)

    fmt.Println()

    c := append(a, 30)
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
}
```

output:
```
[10]

[10]
[10 20]

[10]
[10 30]
[10 30]
```
