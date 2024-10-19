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

func TestIntegerCombinations(t *testing.T) {
	tt := []struct {
		name string
		in   []int
		out  [][]int
	}{
		{
			name: "Empty slice",
			in:   []int{},
			out:  nil,
		},
		{
			name: "Single item",
			in:   []int{1},
			out: [][]int{
				{1},
			},
		},
		{
			name: "Two items",
			in:   []int{1, 2},
			out: [][]int{
				{1},
				{2},
				{1, 2},
			},
		},
		{
			name: "Three items",
			in:   []int{1, 2, 3},
			out: [][]int{
				{1},
				{2},
				{1, 2},
				{3},
				{1, 3},
				{2, 3},
				{1, 2, 3},
			},
		},
		{
			name: "Four items",
			in:   []int{1, 2, 3, 4},
			out: [][]int{
				{1},
				{2},
				{1, 2},
				{3},
				{1, 3},
				{2, 3},
				{1, 2, 3},
				{4},
				{1, 4},
				{2, 4},
				{1, 2, 4},
				{3, 4},
				{1, 3, 4},
				{2, 3, 4},
				{1, 2, 3, 4},
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

func TestAllWithRepetitions(t *testing.T) {
	tt := []struct {
		name string
		in   []string
		m    int
		out  [][]string
	}{
		{
			name: "Empty slice",
			in:   []string{},
			m:    1,
			out:  nil,
		},
		{
			name: "Single item, m = 1",
			in:   []string{"A"},
			m:    1,
			out: [][]string{
				{"A"},
			},
		},
		{
			name: "Single item, m = 2",
			in:   []string{"A"},
			m:    2,
			out: [][]string{
				{"A"},
				{"A", "A"},
			},
		},
		{
			name: "Two items, m = 1",
			in:   []string{"A", "B"},
			m:    1,
			out: [][]string{
				{"A"},
				{"B"},
			},
		},
		{
			name: "Two items, m = 2",
			in:   []string{"A", "B"},
			m:    2,
			out: [][]string{
				{"A"},
				{"B"},
				{"A", "A"},
				{"A", "B"},
				{"B", "A"},
				{"B", "B"},
			},
		},
		{
			name: "Three items, m = 2",
			in:   []string{"A", "B", "C"},
			m:    2,
			out: [][]string{
				{"A"},
				{"B"},
				{"C"},
				{"A", "A"},
				{"A", "B"},
				{"A", "C"},
				{"B", "A"},
				{"B", "B"},
				{"B", "C"},
				{"C", "A"},
				{"C", "B"},
				{"C", "C"},
			},
		},
		{
			name: "Two items, m = 3",
			in:   []string{"A", "B"},
			m:    3,
			out: [][]string{
				{"A"},
				{"B"},
				{"A", "A"},
				{"A", "B"},
				{"B", "A"},
				{"B", "B"},
				{"A", "A", "A"},
				{"A", "A", "B"},
				{"A", "B", "A"},
				{"A", "B", "B"},
				{"B", "A", "A"},
				{"B", "A", "B"},
				{"B", "B", "A"},
				{"B", "B", "B"},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			out := AllRepeat(tc.in, tc.m)
			if !reflect.DeepEqual(out, tc.out) {
				t.Errorf("error: \nreturn:\t%v\nwant:\t%v", out, tc.out)
			}
		})
	}
}

func TestAllWithRepetitionsInt(t *testing.T) {
	tt := []struct {
		name string
		in   []int
		m    int
		out  [][]int
	}{
		{
			name: "Empty slice",
			in:   []int{},
			m:    1,
			out:  nil,
		},
		{
			name: "Single item, m = 1",
			in:   []int{1},
			m:    1,
			out: [][]int{
				{1},
			},
		},
		{
			name: "Single item, m = 2",
			in:   []int{1},
			m:    2,
			out: [][]int{
				{1},
				{1, 1},
			},
		},
		{
			name: "Two items, m = 2",
			in:   []int{1, 2},
			m:    2,
			out: [][]int{
				{1},
				{2},
				{1, 1},
				{1, 2},
				{2, 1},
				{2, 2},
			},
		},
		{
			name: "Three items, m = 2",
			in:   []int{1, 2, 3},
			m:    2,
			out: [][]int{
				{1},
				{2},
				{3},
				{1, 1},
				{1, 2},
				{1, 3},
				{2, 1},
				{2, 2},
				{2, 3},
				{3, 1},
				{3, 2},
				{3, 3},
			},
		},
		{
			name: "Two items, m = 3",
			in:   []int{1, 2},
			m:    3,
			out: [][]int{
				{1},
				{2},
				{1, 1},
				{1, 2},
				{2, 1},
				{2, 2},
				{1, 1, 1},
				{1, 1, 2},
				{1, 2, 1},
				{1, 2, 2},
				{2, 1, 1},
				{2, 1, 2},
				{2, 2, 1},
				{2, 2, 2},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			out := AllRepeat(tc.in, tc.m)
			if !reflect.DeepEqual(out, tc.out) {
				t.Errorf("error: \nreturn:\t%v\nwant:\t%v", out, tc.out)
			}
		})
	}
}
