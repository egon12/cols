# cols
Colllections function that use go1.18 generic

Examples:

```go
import (
	"fmt"

	"github.com/egon12/cols"
)

func ExampleMap() {
	in := []Obj{
		{Name: "Peter", BirthDate: 31},
		{Name: "Simon", BirthDate: 29},
		{Name: "Kikorosuma", BirthDate: 10},
	}

	got := cols.Map(in, func(o Obj) string {
		return o.Name
	})

	fmt.Println(got)
	// Output: [Peter Simon Kikorosuma]
}
```

For complete documentation you can go to pkg.go.dev/github.com/egon12/cols
