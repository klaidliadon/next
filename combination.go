package next

import "iter"

// Returns a channel of combinantions of n element from base w/o repetition
func Combination[T any](base []T, n int, repeat bool) iter.Seq[[]T] {
	if n < 0 {
		n = 0
	}
	if repeat {
		return repeatCombination[T](base).Of(n)
	} else {
		return combination[T](base).Of(n)
	}
}

// A collection of elements for calculating combinations.
type combination[T any] []T

// Returns a channel of possible combinations of r elements.
func (c combination[T]) Of(r int) func(yield func([]T) bool) {
	base := c
	n, t := len(c), count(c, r)
	if t == 0 {
		return func(yield func([]T) bool) {}
	}
	return func(yield func([]T) bool) {
		idxs := make([]int, r)
		for i := range idxs {
			idxs[i] = i
		}
		if !yield(getResult(base, idxs)) {
			return
		}
		for i, j := 1, r-1; i < t; i++ {
			if idxs[j] == j+n-r {
				for idxs[j] == j+n-r {
					j--
				}
				v := idxs[j] + 1
				for i := j; i < r; i++ {
					idxs[i] = v
					v++
				}
				j = r - 1
			} else {
				idxs[j] = idxs[j] + 1
			}
			if !yield(getResult(base, idxs)) {
				return
			}
		}
	}
}

// A collection of elements for calculating combinations.
type repeatCombination[T any] []T

func (c repeatCombination[T]) Of(r int) func(yield func([]T) bool) {
	base := []T(c)
	n, t := len(c), count(c, r)
	return func(yield func([]T) bool) {
		idxs := make([]int, r)
		if !yield(getResult(base, idxs)) {
			return
		}
		for i, j := 1, r-1; i < t; i++ {
			if idxs[j] == n-1 {
				for idxs[j] == n-1 {
					j--
				}
				v := idxs[j] + 1
				for i := j; i < r; i++ {
					idxs[i] = v
				}
				j = r - 1
			} else {
				idxs[j] = idxs[j] + 1
			}
			if !yield(getResult(base, idxs)) {
				return
			}
		}
	}
}
