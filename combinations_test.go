package combinations

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStringCombinations(t *testing.T) {
	tt := []struct {
		name string
		in   []string
		out  [][]string
	}{
		{
			name: "Empty slice",
			in:   []string{},
			out:  nil,
		},
		{
			name: "Single item",
			in:   []string{"A"},
			out: [][]string{
				{"A"},
			},
		},
		{
			name: "Two items",
			in:   []string{"A", "B"},
			out: [][]string{
				{"A"},
				{"B"},
				{"A", "B"},
			},
		},
		{
			name: "Three items",
			in:   []string{"A", "B", "C"},
			out: [][]string{
				{"A"},
				{"B"},
				{"A", "B"},
				{"C"},
				{"A", "C"},
				{"B", "C"},
				{"A", "B", "C"},
			},
		},
		{
			name: "Four items",
			in:   []string{"A", "B", "C", "D"},
			out: [][]string{
				{"A"},
				{"B"},
				{"A", "B"},
				{"C"},
				{"A", "C"},
				{"B", "C"},
				{"A", "B", "C"},
				{"D"},
				{"A", "D"},
				{"B", "D"},
				{"A", "B", "D"},
				{"C", "D"},
				{"A", "C", "D"},
				{"B", "C", "D"},
				{"A", "B", "C", "D"},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			out := All(tc.in)
			if !reflect.DeepEqual(out, tc.out) {
				t.Errorf("error: \nreturn:\t%v\nwant:\t%v", out, tc.out)
			}
		})
	}
}

func ExampleAll() {
	combinations := All([]string{"A", "B", "C"})
	fmt.Println(combinations)
	// Output:
	// [[A] [B] [A B] [C] [A C] [B C] [A B C]]
}

func TestStringCombinationsN(t *testing.T) {
	tt := []struct {
		name string
		in   []string
		n    int
		out  [][]string
	}{
		{
			name: "Empty slice",
			in:   []string{},
			n:    1,
			out:  nil,
		},
		{
			name: "Single item",
			in:   []string{"A"},
			n:    1,
			out: [][]string{
				{"A"},
			},
		},
		{
			name: "Two items, n = 0",
			in:   []string{"A", "B"},
			n:    0,
			out: [][]string{
				{"A"},
				{"B"},
				{"A", "B"},
			},
		},
		{
			name: "Two items, n = 1",
			in:   []string{"A", "B"},
			n:    1,
			out: [][]string{
				{"A"},
				{"B"},
			},
		}, {
			name: "Two items, n = 2",
			in:   []string{"A", "B"},
			n:    2,
			out: [][]string{
				{"A", "B"},
			},
		},
		{
			name: "Three items, n = 0",
			in:   []string{"A", "B", "C"},
			n:    0,
			out: [][]string{
				{"A"},
				{"B"},
				{"A", "B"},
				{"C"},
				{"A", "C"},
				{"B", "C"},
				{"A", "B", "C"},
			},
		},
		{
			name: "Three items, n = 1",
			in:   []string{"A", "B", "C"},
			n:    1,
			out: [][]string{
				{"A"},
				{"B"},
				{"C"},
			},
		},
		{
			name: "Three items, n = 2",
			in:   []string{"A", "B", "C"},
			n:    2,
			out: [][]string{
				{"A", "B"},
				{"A", "C"},
				{"B", "C"},
			},
		},
		{
			name: "Three items, n = 3",
			in:   []string{"A", "B", "C"},
			n:    3,
			out: [][]string{
				{"A", "B", "C"},
			},
		},
		{
			name: "Three items, n = 4",
			in:   []string{"A", "B", "C"},
			n:    4,
			out: [][]string{
				{"A", "B", "C"},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			out := Combinations(tc.in, tc.n)
			if !reflect.DeepEqual(out, tc.out) {
				t.Errorf("error: \nreturn:\t%v\nwant:\t%v", out, tc.out)
			}
		})
	}
}
