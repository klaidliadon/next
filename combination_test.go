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
	c := Combination{Base: []interface{}{
		"0a", "0b", "0c", "0d",
		"1a", "1b", "1c", "1d",
		"2a", "2b", "2c", "2d",
		"3a", "3b", "3c", "3d",
		"4a", "4b", "4c", "4d",
	}}
	size := len(c.Base)
	for i := 0; i < size+2; i++ {
		count := countResults(c, i)
		if expected := combinationSize(size, i); count != expected {
			t.Errorf("(%2d %2d) = %d (expected %d)", size, i, count, expected)
		} else {
			t.Logf("(%2d %2d) = %d", size, i, count)
		}
	}
}

func countResults(c Combination, i int) int {
	var count int
	for range c.Results(i) {
		count++
	}
	return count
}
