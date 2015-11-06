package next

// Returns a channel of combinantions of n element from base w/o repetition
func Combination(base []interface{}, n int) <-chan []interface{} {
	if n < 0 {
		n = 0
	}
	return combination(base).of(n)
}

// A collection of elements for calculating combinations.
type combination []interface{}

// Returns a channel of possible combinations of r elements.
func (c combination) of(r int) <-chan []interface{} {
	res := make(chan []interface{})
	go c.results(r, res)
	return res
}

// Calculates the results and send them back to the channel.
func (c combination) results(r int, ch chan<- []interface{}) {
	defer close(ch)
	base := []interface{}(c)
	n, t := len(c), count(c, r)
	if t == 0 {
		return
	}
	idxs := make([]int, r)
	for i := range idxs {
		idxs[i] = i
	}
	sendIndex(base, idxs, ch)
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
		sendIndex(base, idxs, ch)
	}
}
