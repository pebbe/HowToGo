If you have this:

```go
type MyType int
```

You can't do this:
``` go
a := []int{1, 2, 3}
b := []MyType(a)
```

Instead, you have to do this:
```go
a := []int{1, 2, 3}
b := make([]MyType, len(a))
for i, v := range a {
    b[i] = MyType(v)
}
```

See https://github.com/golang/go/issues/20621
