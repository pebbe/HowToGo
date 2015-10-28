To unmarshall XML that is not in UTF-8:

```go
import (
	"github.com/paulrosania/go-charset/charset"
	_ "github.com/paulrosania/go-charset/data"
	"encoding/xml"
)
```

...

```go

	d := xml.NewDecoder(os.Stdin)
	d.CharsetReader = charset.NewReader
	var t myType
	err := d.Decode(&t)
```
