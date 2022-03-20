package cols

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	in := []ExampleObj{
		{Name: "Egon", BirthDate: 31},
		{Name: "Simon", BirthDate: 29},
		{Name: "Kikorosuma", BirthDate: 10},
	}

	want := []string{"Egon", "Simon", "Kikorosuma"}

	fun := func(arg ExampleObj) string { return arg.Name }

	got := Map(in, fun)

	assert.Equal(t, want, got)
}

func TestFilter(t *testing.T) {
	in := []ExampleObj{
		{Name: "Egon", BirthDate: 31},
		{Name: "Simon", BirthDate: 29},
		{Name: "Kikorosuma", BirthDate: 10},
	}

	want := []ExampleObj{
		{Name: "Egon", BirthDate: 31},
		{Name: "Simon", BirthDate: 29},
	}

	fun := func(arg ExampleObj) bool { return arg.BirthDate > 15 }

	got := Filter(in, fun)

	assert.Equal(t, want, got)
}

type ExampleObj struct {
	Name      string
	BirthDate int32
}
