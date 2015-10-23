package next

// A combination of elements.
type Combination []interface{}

// Returns a channel of possible combinations of l elements.
func (c *Combination) Results(r int) <-chan []interface{} {
	res := make(chan []interface{})
	go c.results(r, res)
	return res
}

func (c *Combination) results(r int, ch chan<- []interface{}) {
	defer close(ch)
	n := len(*c)
	if r > n {
		return
	}
	idxs := make([]int, r)
	for i := range idxs {
		idxs[i] = i
	}
	cmb := make([]interface{}, r)
	for i, el := range idxs {
		cmb[i] = (*c)[el]
	}
	ch <- cmb
	for {
		i := r - 1
		for i >= 0 && idxs[i] == i+n-r {
			i -= 1
		}
		if i < 0 {
			break
		}
		idxs[i] += 1
		for j := i + 1; j < r; j += 1 {
			idxs[j] = idxs[j-1] + 1
		}
		for ; i < len(idxs); i += 1 {
			cmb[i] = (*c)[idxs[i]]
		}
		ch <- cmb
	}
}
