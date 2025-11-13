package next

import (
	"fmt"
	"testing"
)

type testbase bool

func (t testbase) of(int) <-chan []testbase { return nil }

func newCases(n int) []base[string] {
	b := make([]string, n)
	for i := range b {
		b[i] = fmt.Sprintf("a%d", i)
	}
	return []base[string]{
		combination[string](b[:]),
		permutation[string](b[:]),
		repeatCombination[string](b[:]),
		repeatPermutation[string](b[:]),
	}
}

func TestUnknown(t *testing.T) {
	if c := count(testbase(false), 2); c != 0 {
		t.Fail()
	}
}

func TestSizes(t *testing.T) {
	size := 7
	cases := newCases(size)
	for _, c := range cases {
		for i := 0; i < size+2; i++ {
			var tot int
			ch := c.of(i)
			for _ = range ch {
				tot++
			}
			if expected := count(c, i); tot != expected {
				t.Errorf("%T (%2d %2d) = %d (expected %d)", c, size, i, tot, expected)
			} else {
				t.Logf("%T (%2d %2d) = %d", c, size, i, tot)
			}
		}
	}
}

func TestCreation(t *testing.T) {
	base := []int{1, 2, 3}
	Combination(base, 2, false)
	Combination(base, -1, true)
	Permutation(base, 2, false)
	Permutation(base, -1, true)
}

func TestShowcase(t *testing.T) {
	size := 5
	cases := newCases(size)
	for _, c := range cases {
		var r = make([][]string, 0, count(c, size))
		for v := range c.of(3) {
			r = append(r, v)
		}
		t.Logf("Result for %T: %v", c, r)
	}
}
