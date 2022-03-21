package cols_test

import (
	"fmt"

	"github.com/egon12/cols"
)

func ExampleSumBy() {
	in := []SumAbleObj{
		{Name: "trans_1", Group: "account_1", Amount: 1.0},
		{Name: "trans_2", Group: "account_1", Amount: 3.0},
		{Name: "trans_3", Group: "account_2", Amount: 2.5},
		{Name: "trans_4", Group: "account_2", Amount: 7.9},
	}

	got := cols.SumBy(
		in,
		func(o SumAbleObj) string { return o.Group },
		func(o SumAbleObj) float64 { return o.Amount },
	)

	fmt.Println(got)
	// Output: map[account_1:4 account_2:10.4]
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

type SumAbleObj struct {
	Name   string
	Group  string
	Amount float64
}
