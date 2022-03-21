package cols_test

import (
	"fmt"

	"github.com/egon12/cols"
)

func ExampleMapString() {
	in := []Obj{
		{Name: "Peter", BirthDate: 31},
		{Name: "Simon", BirthDate: 29},
		{Name: "Kikorosuma", BirthDate: 10},
	}

	got := cols.MapString(in, func(o Obj) string { return o.Name })

	fmt.Println(got)
	// Output: [Peter Simon Kikorosuma]
}

func ExampleJoinString() {
	in := []Obj{
		{Name: "Peter", BirthDate: 31},
		{Name: "Simon", BirthDate: 29},
		{Name: "Kikorosuma", BirthDate: 10},
	}

	got := cols.JoinString(in, func(o Obj) string { return o.Name })

	fmt.Println(got)
	// Output: Peter,Simon,Kikorosuma
}
