// Package combinations provides a method to generate all combinations out of a given generic array.
package combinations

import "math/bits"

// All returns all combinations for a given generic array.
// This is essentially a powerset of the given set except that the empty set is disregarded.
func All[T any](set []T) (subsets [][]T) {
	length := uint(len(set))

	// Go through all possible combinations of objects
	// from 1 (only first object in subset) to 2^length (all objects in subset)
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		var subset []T

		for object := uint(0); object < length; object++ {
			// checks if object is contained in subset
			// by checking if bit 'object' is set in subsetBits
			if (subsetBits>>object)&1 == 1 {
				// add object to subset
				subset = append(subset, set[object])
			}
		}
		// add subset to subsets
		subsets = append(subsets, subset)
	}
	return subsets
}

// AllRepeat returns all combinations with repetitions for a given slice,
// from 1 up to a maximum combination length of m.
func AllRepeat[T any](set []T, m int) (subsets [][]T) {
	if m < 1 {
		return nil
	}

	var generateCombos func([]T, int)
	generateCombos = func(current []T, depth int) {
		if depth == 0 {
			subset := make([]T, len(current))
			copy(subset, current)
			subsets = append(subsets, subset)
			return
		}

		for _, item := range set {
			generateCombos(append(current, item), depth-1)
		}
	}

	for length := 1; length <= m; length++ {
		generateCombos([]T{}, length)
	}

	return subsets
}

// Combinations returns combinations of n elements for a given generic array.
// For n < 1, it equals to All and returns all combinations.
func Combinations[T any](set []T, n int) (subsets [][]T) {
	length := uint(len(set))

	if n > len(set) {
		n = len(set)
	}

	// Go through all possible combinations of objects
	// from 1 (only first object in subset) to 2^length (all objects in subset)
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		if n > 0 && bits.OnesCount(uint(subsetBits)) != n {
			continue
		}

		var subset []T

		for object := uint(0); object < length; object++ {
			// checks if object is contained in subset
			// by checking if bit 'object' is set in subsetBits
			if (subsetBits>>object)&1 == 1 {
				// add object to subset
				subset = append(subset, set[object])
			}
		}
		// add subset to subsets
		subsets = append(subsets, subset)
	}
	return subsets
}
