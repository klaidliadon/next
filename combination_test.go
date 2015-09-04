package next

import (
	"runtime"
	"testing"
)

func combinationSize(n, r int) int {
	if r <= 0 || r > n {
		return 0
	}
	var result = 1
	for i := 1; i <= r; i++ {
		result = result * (n - r + i) / i
	}
	return result
}

var proc = runtime.NumCPU()

func init() {
	runtime.GOMAXPROCS(proc)
}

func TestCombinationSize(t *testing.T) {
	c := Combination{Base: []interface{}{"a", "b", "c", "d", "e", "f", -1, 3, 76, 1000, []interface{}{}, 0, 898, 23, 239, 29, 0x22, "sed"}}
	size := len(c.Base)
	for i := 0; i < size+2; i++ {
		var count int
		for range c.Results(uint(i)) {
			count++
		}
		if expected := combinationSize(size, i); count != expected {
			t.Errorf("(%2d %2d) = %d (expected %d)", size, i, count, expected)
		} else {
			t.Logf("(%2d %2d) = %d", size, i, count)
		}
	}
}
