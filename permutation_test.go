package next

import (
	"strings"
	"testing"
)

func TestPermutations(t *testing.T) {
	base := strings.Split("abcd", "")
	size := 3

	seq := Permutation(base, size)
	r := make([]string, 0)
	for v := range seq {
		r = append(r, strings.Join(v, ""))
	}
	t.Logf("Permutations of %d elements over %d:\n%v", len(base), size, r)
}

func TestRepeatPermutations(t *testing.T) {
	base := strings.Split("abcd", "")
	size := 3
	seq := RepeatPermutation(base, size)
	r := make([]string, 0)
	for v := range seq {
		r = append(r, strings.Join(v, ""))
	}
	t.Logf("Permutations of %d elements over %d with repetition:\n%v", len(base), size, r)
}
