package next

import "sync"

// The combination struct.
type Combination struct {
	Base []interface{}
}

// Returns a channel of possible combinations of l elements.
func (c *Combination) Results(l int) <-chan []interface{} {
	wg, ch, bs := new(sync.WaitGroup), make(chan []interface{}), len(c.Base)
	defer func() {
		go func() { wg.Wait(); close(ch) }()
	}()
	if l < 1 || l > bs {
		return ch
	}
	wg.Add(1)
	if l == 1 {
		// simple case, every element in the base
		go func() {
			defer wg.Done()
			for _, v := range c.Base {
				ch <- []interface{}{v}
			}
		}()
		return ch
	} else if l == bs {
		// simple case, 1 result (the base)
		go func() { defer wg.Done(); ch <- c.Base }()
		return ch
	}
	max := bs - l
	wg.Add(max)
	for i := int(0); i < max+1; i++ {
		go func(l, i int) { defer wg.Done(); c.thread(wg, ch, l, i) }(l, i)
	}
	return ch
}

// A single thread where the first element is fixed
func (c *Combination) thread(wg *sync.WaitGroup, ch chan []interface{}, l, s int) {
	e, bs := &element{index: s}, len(c.Base)
	for i, max := s, bs+2-l; e.depth != 1 || e.index < max; i++ {
		// if the next element is out of range
		if i+1 == bs {
			// if there is no previous element stop
			if e.prev == nil {
				break
			}
			// or go to the previous element using the index
			i, e = e.index-1, e.prev
			continue
		}
		e = &element{index: i + 1, depth: e.depth + 1, prev: e}
		if e.depth+1 == l {
			ch <- e.value(c.Base)
			i = e.index - 1
			e = e.prev
		}
	}
}
