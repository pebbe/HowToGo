To unmarshall XML that is not in UTF-8:

```go
import (
    "golang.org/x/net/html/charset"
    "encoding/xml"
)
```

...

```go

    d := xml.NewDecoder(os.Stdin)
    d.CharsetReader = charset.NewReaderLabel
    var t myType
    err := d.Decode(&t)
```
