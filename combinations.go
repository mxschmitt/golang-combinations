// Package combinations provides a method to generate all combinations out of a given string array.
package combinations

import "math/bits"

// All returns all combinations for a given string array.
// This is essentially a powerset of the given set except that the empty set is disregarded.
func All(set []string, fs ...AddFunction) (subsets [][]string) {
	length := uint(len(set))

	fLen := len(fs)

	// Go through all possible combinations of objects
	// from 1 (only first object in subset) to 2^length (all objects in subset)
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		var subset []string

		for object := uint(0); object < length; object++ {
			// checks if object is contained in subset
			// by checking if bit 'object' is set in subsetBits
			if (subsetBits>>object)&1 == 1 {
				// add object to subset
				subset = append(subset, set[object])
			}
		}

		if addFunc(subset, fLen, fs...){
			// add subset to subsets
			subsets = append(subsets, subset)
		}

	}
	return subsets
}

// Combinations returns combinations of n elements for a given string array.
// For n < 1, it equals to All and returns all combinations.
func Combinations(set []string, n int, fs ...AddFunction) (subsets [][]string) {
	length := uint(len(set))

	if n > len(set) {
		n = len(set)
	}

	fLen := len(fs)

	// Go through all possible combinations of objects
	// from 1 (only first object in subset) to 2^length (all objects in subset)
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		if n > 0 && bits.OnesCount(uint(subsetBits)) != n {
			continue
		}

		var subset []string

		for object := uint(0); object < length; object++ {
			// checks if object is contained in subset
			// by checking if bit 'object' is set in subsetBits
			if (subsetBits>>object)&1 == 1 {
				// add object to subset
				subset = append(subset, set[object])
			}
		}

		if addFunc(subset, fLen, fs...){
			// add subset to subsets
			subsets = append(subsets, subset)
		}
	}
	return subsets
}

type AddFunction func(subset []string)bool

func addFunc(subset []string, fLen int ,fs ...AddFunction)bool{
	// add subset to subsets
	if fLen == 0{
		return true
	}
	for _, f := range fs {
		if f(subset){
			return true
		}
	}
	return false
}