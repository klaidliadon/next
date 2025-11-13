package next

// Returns a channel of combinantions of n element from base w/o repetition
func Combination[T any](base []T, n int, repeat bool) <-chan []T {
	if n < 0 {
		n = 0
	}
	if repeat {
		return repeatCombination[T](base).of(n)
	} else {
		return combination[T](base).of(n)
	}
}

// A collection of elements for calculating combinations.
type combination[T any] []T

// Returns a channel of possible combinations of r elements.
func (c combination[T]) of(r int) <-chan []T {
	res := make(chan []T)
	go c.results(r, res)
	return res
}

// Calculates the results and send them back to the channel.
func (c combination[T]) results(r int, ch chan<- []T) {
	defer close(ch)
	base := c
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

// A collection of elements for calculating combinations.
type repeatCombination[T any] []T

func (c repeatCombination[T]) of(r int) <-chan []T {
	res := make(chan []T)
	go c.results(r, res)
	return res
}

// Calculates the results and send them back to the channel.
func (c repeatCombination[T]) results(r int, ch chan<- []T) {
	defer close(ch)
	base := []T(c)
	n, t := len(c), count(c, r)
	idxs := make([]int, r)
	sendIndex(base, idxs, ch)
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
		sendIndex(base, idxs, ch)
	}
}
