package cols_test

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

func ExampleSum() {
	in := []Obj{
		{Name: "Peter", BirthDate: 31},
		{Name: "Simon", BirthDate: 29},
		{Name: "Kikorosuma", BirthDate: 10},
	}

	got := cols.Sum(in, func(o Obj) int32 {
		return o.BirthDate
	})

	fmt.Println(got)
	// Output: 70
}

func ExampleMax() {
	in := []Obj{
		{Name: "Peter", BirthDate: 31},
		{Name: "Simon", BirthDate: 29},
		{Name: "Kikorosuma", BirthDate: 10},
	}

	got := cols.Max(in, func(o Obj) int32 {
		return o.BirthDate
	})

	fmt.Println(got)
	// Output: 31
}

func ExampleMin() {
	in := []Obj{
		{Name: "Peter", BirthDate: 31},
		{Name: "Simon", BirthDate: 29},
		{Name: "Kikorosuma", BirthDate: 10},
	}

	got := cols.Min(in, func(o Obj) int32 {
		return o.BirthDate
	})

	fmt.Println(got)
	// Output: 10
}

func ExampleSumNumber() {
	in1 := []int{1, 2, 3, 4}
	got := cols.SumNumber(in1)
	fmt.Println(got)
	// Output: 10
}

func ExampleMaxNumber() {
	got := cols.MaxNumber([]int32{10, 20, 100, 1, 2, 5})
	fmt.Println(got)
	// Output: 100
}

func ExampleMinNumber() {
	got := cols.MaxNumber([]int32{10, 20, 100, 1, 2, 5})
	fmt.Println(got)
	// Output: 100
}

type Obj struct {
	Name      string
	BirthDate int32
}

type GroupObj struct {
	Name      string
	GroupName string
}
