package next

// Returns a channel of permutations of n element from base w/o repetition
func Permutation[T any](base []T, n int, repeat bool) <-chan []T {
	if n < 0 {
		n = 0
	}
	if repeat {
		return repeatPermutation[T](base).of(n)
	} else {
		return permutation[T](base).of(n)
	}
}

// A combination of elements.
type permutation[T any] []T

// Returns a channel of possible combinations of l elements.
func (p permutation[T]) of(r int) <-chan []T {
	res := make(chan []T)
	go p.results(r, res)
	return res
}

func (p permutation[T]) results(r int, ch chan<- []T) {
	defer close(ch)
	n := len(p)
	if r > n {
		return
	}
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
		cmb[i] = p[el]
	}
	copy(res, cmb)
	ch <- res
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
					cmb[k] = p[idxs[k]]
				}
				res := make([]T, r)
				copy(res, cmb)
				ch <- res
				break
			}
		}
		if i < 0 {
			return
		}

	}
}

// A combination of elements.
type repeatPermutation[T any] []T

// Returns a channel of possible combinations of l elements.
func (p repeatPermutation[T]) of(r int) <-chan []T {
	res := make(chan []T)
	go p.results(r, res)
	return res
}

func (p repeatPermutation[T]) results(r int, ch chan<- []T) {
	defer close(ch)
	n, t := len(p), count[T](p, r)
	for i := 0; i < t; i++ {
		v := make([]T, r)
		j := i
		for k := 0; k < r; k++ {
			x := j % n
			j = int(j / n)
			v[k] = p[x]
		}
		ch <- v
	}
}
