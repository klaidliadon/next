package next

import (
	"iter"
	"math"
)

// Permutation returns an iterator of permutations of n element from base without repetition
func Permutation[T any](elements []T, r int) iter.Seq[[]T] {
	if r < 0 {
		r = 0
	}

	base := elements
	n := len(elements)
	if r > n {
		return func(yield func([]T) bool) {}
	}

	return func(yield func([]T) bool) {
		idxs := make([]int, n)
		for i := range idxs {
			idxs[i] = i
		}
		cycles := make([]int, r)
		for i := range cycles {
			cycles[i] = n - i
		}
		cmb := make([]T, r)
		res := make([]T, r)
		for i, el := range idxs[:r] {
			cmb[i] = base[el]
		}
		copy(res, cmb)
		if !yield(res) {
			return
		}
		for n > 0 {
			i := r - 1
			for ; i >= 0; i -= 1 {
				cycles[i] -= 1
				if cycles[i] == 0 {
					index := idxs[i]
					for j := i; j < n-1; j += 1 {
						idxs[j] = idxs[j+1]
					}
					idxs[n-1] = index
					cycles[i] = n - i
				} else {
					j := cycles[i]
					idxs[i], idxs[n-j] = idxs[n-j], idxs[i]
					for k := i; k < r; k += 1 {
						cmb[k] = base[idxs[k]]
					}
					res := make([]T, r)
					copy(res, cmb)
					if !yield(res) {
						return
					}
					break
				}
			}
			if i < 0 {
				return
			}
		}
	}
}

func RepeatPermutation[T any](elements []T, r int) iter.Seq[[]T] {
	if r < 0 {
		r = 0
	}

	return func(yield func([]T) bool) {
		base := elements
		n := len(elements)
		t := int(math.Pow(float64(n), float64(r)))
		for i := range t {
			v := make([]T, r)
			j := i
			for k := range r {
				x := j % n
				j = int(j / n)
				v[k] = base[x]
			}
			if !yield(v) {
				return
			}
		}
	}

}
