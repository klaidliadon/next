package next

import "sync"

// An element of the combination
type element struct {
	prev         *element
	depth, index uint
}

// The combination struct
type Combination struct {
	Base []interface{}
}

// Returns a channel of possible combinations
func (c *Combination) Results(l uint) <-chan []interface{} {
	wg, ch, bs := new(sync.WaitGroup), make(chan []interface{}), uint(len(c.Base))
	defer func() { go func() { wg.Wait(); close(ch) }() }()
	if l < 1 || l > bs {
		return ch
	}
	wg.Add(1)
	if l == 1 {
		go func() {
			defer wg.Done()
			for _, v := range c.Base {
				ch <- []interface{}{v}
			}
		}()
		return ch
	} else if l == bs {
		go func() { defer wg.Done(); ch <- c.Base }()
		return ch
	}
	max := bs - l
	wg.Add(int(max))
	for i := uint(0); i < max+1; i++ {
		go func(l, i uint) { defer wg.Done(); c.thread(wg, ch, l, i) }(l, i)
	}
	return ch
}

// A single thread where the first element is fixed
func (c *Combination) thread(wg *sync.WaitGroup, ch chan []interface{}, l, s uint) {
	e, bs := &element{index: s}, uint(len(c.Base))
	for i, max := s, bs+2-l; e.depth != 1 || e.index < max; i++ {
		if i+1 == bs {
			if e.prev == nil {
				break
			}
			i = e.index - 1
			e = e.prev
			continue
		}
		e = &element{index: i + 1, depth: e.depth + 1, prev: e}
		if e.depth+1 == l {
			ch <- value(c, e)
			i = e.index - 1
			e = e.prev
		}
	}
}

// Returns the actual combination using an element
func value(c *Combination, e *element) []interface{} {
	result := make([]interface{}, e.depth+1)
	var curr = e
	for i := uint(0); i <= e.depth; i++ {
		result[e.depth-i] = c.Base[curr.index]
		curr = curr.prev
	}
	return result
}
