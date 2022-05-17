package cols_test

import (
	"fmt"

	"github.com/egon12/cols"
)

func ExampleCollection() {
	someSlice := []string{
		"one",
		"two",
		"three",
	}

	c := cols.From(someSlice)

	c = c.Filter(func(item string) bool { return item != "three" })

	fmt.Println(c.Slice())
	//output: [one two]
}
