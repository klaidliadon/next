package next

// Returns a channel of permutations of n element from base w/o repetition
func Permutation(base []interface{}, n int, repeat bool) <-chan []interface{} {
	if n < 0 {
		n = 0
	}
	if repeat {
		return repeatPermutation(base).of(n)
	} else {
		return permutation(base).of(n)
	}
}

// A combination of elements.
type permutation []interface{}

// Returns a channel of possible combinations of l elements.
func (p permutation) of(r int) <-chan []interface{} {
	res := make(chan []interface{})
	go p.results(r, res)
	return res
}

func (p permutation) results(r int, ch chan<- []interface{}) {
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
	cmb := make([]interface{}, r)
	res := make([]interface{}, r)
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
				res := make([]interface{}, r)
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
type repeatPermutation []interface{}

// Returns a channel of possible combinations of l elements.
func (p repeatPermutation) of(r int) <-chan []interface{} {
	res := make(chan []interface{})
	go p.results(r, res)
	return res
}

func (p repeatPermutation) results(r int, ch chan<- []interface{}) {
	defer close(ch)
	n, t := len(p), count(p, r)
	for i := 0; i < t; i++ {
		v := make([]interface{}, r)
		j := i
		for k := 0; k < r; k++ {
			x := j % n
			j = int(j / n)
			v[k] = p[x]
		}
		ch <- v
	}
}
