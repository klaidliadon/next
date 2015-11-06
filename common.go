package next

type base interface {
	of(int) <-chan []interface{}
}

func count(b base, r int) int {
	switch c := b.(type) {
	case combination:
		n := len(c)
		if r > n {
			return 0
		}
		var t = 1
		for i := 1; i <= r; i++ {
			t = t * (n - r + i) / i
		}
		return t
	case permutation:
		n := len(c)
		if r > n {
			return 0
		}
		var t = 1
		for i := n - r + 1; i <= n; i++ {
			t = t * i
		}
		return t
	}
	return 0
}

// Creates a new result using the selected indexes and sends them to the channel
func sendIndex(base []interface{}, index []int, ch chan<- []interface{}) {
	r := len(index)
	res := make([]interface{}, r)
	for i, idx := range index {
		res[i] = base[idx]
	}
	ch <- res
}
