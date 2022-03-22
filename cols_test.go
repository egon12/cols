package cols_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/egon12/cols"
)

type Obj struct {
	Name      string
	BirthDate int32
}

type GroupObj struct {
	Name      string
	GroupName string
}

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

func ExampleFilter() {
	in := []Obj{
		{Name: "Peter", BirthDate: 31},
		{Name: "Simon", BirthDate: 29},
		{Name: "Kikorosuma", BirthDate: 10},
	}

	got := cols.Filter(in, func(o Obj) bool {
		return o.BirthDate > 15
	})

	fmt.Println(got)
	// Output: [{Peter 31} {Simon 29}]
}

func ExampleReduce() {
	in := []int{1, 2, 3}

	got := cols.Reduce(in, func(a float32, i int) float32 {
		return a + float32(i)
	}, 0.0)

	fmt.Println(got)
	// Output: 6
}

func ExampleFind() {
	in := []Obj{
		{Name: "Peter", BirthDate: 31},
		{Name: "Simon", BirthDate: 29},
		{Name: "Kikorosuma", BirthDate: 10},
	}

	got, _ := cols.Find(in, func(o Obj) bool {
		return o.Name == "Simon"
	})

	fmt.Println(got)
	// Output: {Simon 29}
}

func ExampleAny() {
	in := []Obj{
		{Name: "Peter", BirthDate: 31},
		{Name: "Simon", BirthDate: 29},
		{Name: "Kikorosuma", BirthDate: 10},
	}

	got := cols.Any(in, func(o Obj) bool {
		return o.Name == "Simon"
	})

	fmt.Println(got)
	// Output: true
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

	got := cols.GroupBy(in, func(o GroupObj) string {
		return o.GroupName
	})

	fmt.Println(got["Group 1"])
	// Output: [{Person 1 Group 1} {Person 2 Group 1} {Person 3 Group 1} {Person 4 Group 1}]
}

func ExampleUnique() {
	in := []string{
		"one",
		"two",
		"three",
		"two",
		"two",
		"one",
		"three",
	}

	got := cols.Unique(in)

	sort.Strings(got)
	fmt.Println(got)
	// Output: [one three two]
}

func TestFind(t *testing.T) {
	t.Run("func Find if not that same should return false",
		func(t *testing.T) {
			got, ok := cols.Find([]int{1, 2, 3}, func(it int) bool {
				return it == 4
			})

			if ok {
				t.Error("should not ok, cause not found any")
			}

			if got != 0 {
				t.Errorf("got should be default value (0) got %v", got)
			}
		})
}
