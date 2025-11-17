package next

import "iter"

// Combination returns an iterator of combinations of n element from base without repetition
func Combination[T any](elements []T, r int) iter.Seq[[]T] {
	if r < 0 {
		r = 0
	}

	base := elements
	n := len(elements)
	if r > n {
		return func(yield func([]T) bool) {}
	}

	results := 1
	for i := 1; i <= r; i++ {
		results = results * (n - r + i) / i
	}

	return func(yield func([]T) bool) {
		idxs := make([]int, r)
		for i := range idxs {
			idxs[i] = i
		}
		if !yieldResult(yield, base, idxs) {
			return
		}
		for i, j := 1, r-1; i < results; i++ {
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
			if !yieldResult(yield, base, idxs) {
				return
			}
		}
	}
}

// RepeatCombination returns a channel of combinantions of n element from base with repetition
func RepeatCombination[T any](elements []T, r int) iter.Seq[[]T] {
	if r < 0 {
		r = 0
	}

	base := elements
	n := len(elements)

	results := 1
	for i := r + 1; i < n+r; i++ {
		results *= i
	}
	for i := 1; i < n; i++ {
		results /= i
	}
	return func(yield func([]T) bool) {
		idxs := make([]int, r)
		if !yieldResult(yield, base, idxs) {
			return
		}
		for i, j := 1, r-1; i < results; i++ {
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
			if !yieldResult(yield, base, idxs) {
				return
			}
		}
	}
}

func yieldResult[T any](yield func([]T) bool, base []T, index []int) bool {
	res := make([]T, len(index))
	for i, idx := range index {
		res[i] = base[idx]
	}
	return yield(res)
}
