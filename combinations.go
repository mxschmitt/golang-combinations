// Package combinations provides a method to generate all combinations out of a given string array.
package combinations

// All will return all combinations for a given string array
func All(in []string) [][]string {
	length := uint(len(in))
	maxCount := 1 << length
	var out [][]string
	for i := 1; i < maxCount; i++ {
		var item []string
		for j := uint(0); j < length; j++ {
			if i&(1<<j) != 0 {
				item = append(item, in[j])
			}
		}
		out = append(out, item)
	}
	return out
}
