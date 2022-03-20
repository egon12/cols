package examples

import (
	"fmt"

	"github.com/egon12/cols"
)

type Person struct {
	name string
	age  int32
}

var persons = []Person{
	{name: "gilbert", age: 32},
	{name: "johnson", age: 16},
	{name: "dwayne", age: 45},
}

func ExampleFilterMap() {
	in := persons
	temp := cols.Filter(in, func(p Person) bool { return p.age > 20 })
	res := cols.Map(temp, func(p Person) string { return p.name })
	fmt.Println(res)
}
