package next

import "math"

type base[T any] interface {
	of(int) <-chan []T
}

func count[T any](b base[T], r int) int {
	switch c := b.(type) {
	case combination[T]:
		n := len(c)
		if r > n {
			return 0
		}
		var t = 1
		for i := 1; i <= r; i++ {
			t = t * (n - r + i) / i
		}
		return t
	case permutation[T]:
		n := len(c)
		if r > n {
			return 0
		}
		var t = 1
		for i := n - r + 1; i <= n; i++ {
			t = t * i
		}
		return t
	case repeatCombination[T]:
		n := len(c)
		var t = 1
		for i := r + 1; i < n+r; i++ {
			t = t * i
		}
		for i := 1; i < n; i++ {
			t = t / i
		}
		return t
	case repeatPermutation[T]:
		return int(math.Pow(float64(len(c)), float64(r)))
	}
	return 0
}

// Creates a new result using the selected indexes and sends them to the channel
func sendIndex[T any](base []T, index []int, ch chan<- []T) {
	r := len(index)
	res := make([]T, r)
	for i, idx := range index {
		res[i] = base[idx]
	}
	ch <- res
}
