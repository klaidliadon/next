package next

import "testing"

func combSize(n, r int) int {
	if r > n {
		return 0
	}
	var t = 1
	for i := 1; i <= r; i++ {
		t = t * (n - r + i) / i
	}
	return t
}

// n!/(n-r)!
func permSize(n, r int) int {
	if r > n {
		return 0
	}
	var t = 1
	for i := n - r + 1; i <= n; i++ {
		t = t * i
	}
	return t
}

func newComb(a ...interface{}) result {
	c := Combination(a)
	return &c
}

func newPerm(a ...interface{}) result {
	c := Permutation(a)
	return &c
}

type result interface {
	Results(int) <-chan []interface{}
}
type TestCase struct {
	New  func(...interface{}) result
	Size func(int, int) int
}

func Test(t *testing.T) {
	var base = []interface{}{
		"a", "b", "c", "d", "e",
		"f", "g", "h", "i", "j",
	}
	size := len(base)
	var cases = []TestCase{
		TestCase{newComb, combSize},
		TestCase{newPerm, permSize},
	}
	for _, tc := range cases {
		c := tc.New(base...)
		for i := 0; i < size+2; i++ {
			var tot int
			ch := c.Results(i)
			for _ = range ch {
				tot++
			}
			if expected := tc.Size(size, i); tot != expected {
				t.Errorf("(%2d %2d) = %d (expected %d)", size, i, tot, expected)
			} else {
				t.Logf("(%2d %2d) = %d", size, i, tot)
			}
		}
	}
}
