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
				[]string{"A"},
			},
		},
		{
			name: "Two items",
			in:   []string{"A", "B"},
			out: [][]string{
				[]string{"A"},
				[]string{"B"},
				[]string{"A", "B"},
			},
		},
		{
			name: "Three items",
			in:   []string{"A", "B", "C"},
			out: [][]string{
				[]string{"A"},
				[]string{"B"},
				[]string{"A", "B"},
				[]string{"C"},
				[]string{"A", "C"},
				[]string{"B", "C"},
				[]string{"A", "B", "C"},
			},
		},
		{
			name: "Four items",
			in:   []string{"A", "B", "C", "D"},
			out: [][]string{
				[]string{"A"},
				[]string{"B"},
				[]string{"A", "B"},
				[]string{"C"},
				[]string{"A", "C"},
				[]string{"B", "C"},
				[]string{"A", "B", "C"},
				[]string{"D"},
				[]string{"A", "D"},
				[]string{"B", "D"},
				[]string{"A", "B", "D"},
				[]string{"C", "D"},
				[]string{"A", "C", "D"},
				[]string{"B", "C", "D"},
				[]string{"A", "B", "C", "D"},
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
