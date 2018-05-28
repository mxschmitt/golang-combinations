package permutations

import (
	"math"
)

func Strings(in []string) [][]string {
	lenght := len(in)
	switch lenght {
	case 0:
		return [][]string{}
	case 1:
		return [][]string{in}
	}

	maxCount := int(math.Pow(2, float64(lenght)))
	out := [][]string{}
	for i := 1; i < maxCount; i++ {
		item := []string{}
		for j := 0; j < lenght; j++ {
			if (i & (1 << uint(j))) != 0 {
				item = append(item, in[j])
			}
		}
		out = append(out, item)
	}
	return out
}
