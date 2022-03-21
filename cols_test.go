package cols

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleMap() {
	in := []Obj{
		{Name: "Peter", BirthDate: 31},
		{Name: "Simon", BirthDate: 29},
		{Name: "Kikorosuma", BirthDate: 10},
	}

	got := Map(in, func(o Obj) string {
		return o.Name
	})

	fmt.Println(got)
	// Output: [Egon Simon Kikorosuma]
}

func ExampleFilter() {
	in := []Obj{
		{Name: "Peter", BirthDate: 31},
		{Name: "Simon", BirthDate: 29},
		{Name: "Kikorosuma", BirthDate: 10},
	}

	got := Filter(in, func(o Obj) bool {
		return o.BirthDate > 15
	})

	fmt.Println(got)
	// Output: [{Egon 31} {Simon 29}]
}

func ExampleGroupBy() {
	in := []GroupObj{
		{Name: "Person 1", GroupName: "Group 1"},
		{Name: "Person 2", GroupName: "Group 1"},
		{Name: "Person 3", GroupName: "Group 1"},
		{Name: "Person 4", GroupName: "Group 1"},
		{Name: "Person 5", GroupName: "Group 2"},
		{Name: "Person 6", GroupName: "Group 2"},
	}

	got := GroupBy(in, func(o GroupObj) string {
		return o.GroupName
	})

	fmt.Println(got["Group 1"])
	// Output: [{Person 1 Group 1} {Person 2 Group 1} {Person 3 Group 1} {Person 4 Group 1}]
}

func TestMap(t *testing.T) {
	in := []Obj{
		{Name: "Egon", BirthDate: 31},
		{Name: "Simon", BirthDate: 29},
		{Name: "Kikorosuma", BirthDate: 10},
	}

	want := []string{"Egon", "Simon", "Kikorosuma"}

	fun := func(arg Obj) string { return arg.Name }

	got := Map(in, fun)

	assert.Equal(t, want, got)
}

func TestFilter(t *testing.T) {
	in := []Obj{
		{Name: "Egon", BirthDate: 31},
		{Name: "Simon", BirthDate: 29},
		{Name: "Kikorosuma", BirthDate: 10},
	}

	want := []Obj{
		{Name: "Egon", BirthDate: 31},
		{Name: "Simon", BirthDate: 29},
	}

	fun := func(arg Obj) bool { return arg.BirthDate > 15 }

	got := Filter(in, fun)

	assert.Equal(t, want, got)
}

type Obj struct {
	Name      string
	BirthDate int32
}

type GroupObj struct {
	Name      string
	GroupName string
}
