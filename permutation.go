package next

// A combination of elements.
type Permutation []interface{}

// Returns a channel of possible combinations of l elements.
func (p *Permutation) Results(r int) <-chan []interface{} {
	res := make(chan []interface{})
	go p.results(r, res)
	return res
}

func (p *Permutation) results(r int, ch chan<- []interface{}) {
	defer close(ch)
	src := *p
	n := len(src)
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
		cmb[i] = (src)[el]
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
					cmb[k] = src[idxs[k]]
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
