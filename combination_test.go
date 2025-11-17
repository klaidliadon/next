package next

import (
	"strings"
	"testing"
)

func TestCombinations(t *testing.T) {
	base := strings.Split("abcd", "")
	size := 3

	seq := Combination(base, size)
	r := make([]string, 0)
	for v := range seq {
		r = append(r, strings.Join(v, ""))
	}
	t.Logf("Combinations of %d elements over %d:\n%v", len(base), size, r)
}

func TestRepeatCombinations(t *testing.T) {
	base := strings.Split("abcd", "")
	size := 3

	seq := RepeatCombination(base, size)
	r := make([]string, 0)
	for v := range seq {
		r = append(r, strings.Join(v, ""))
	}
	t.Logf("Combinations of %d elements over %d with repetition:\n%v", len(base), size, r)
}
