package next

import "sync"

// The combination struct.
type Combination struct {
	Base []interface{}
}

// Returns a channel of possible combinations of l elements.
func (c *Combination) Results(l int) <-chan []interface{} {
	lock, ch, bs := new(sync.Mutex), make(chan []interface{}), len(c.Base)
	defer func() {
		go func() {
			lock.Lock()
			defer lock.Unlock()
			close(ch)
		}()
	}()
	if l < 1 || l > bs {
		return ch
	}
	lock.Lock()
	if l == 1 {
		// simple case, every element in the base
		go func() {
			defer lock.Unlock()
			for _, v := range c.Base {
				ch <- []interface{}{v}
			}
		}()
		return ch
	} else if l == bs {
		// simple case, 1 result (the base)
		go func() { defer lock.Unlock(); ch <- c.Base }()
		return ch
	}
	go func() {
		defer lock.Unlock()
		series := bs - l + 1
		for s := 0; s < series; s++ {
			e := &element{index: s}
			for i := s; e.depth != 1 || e.index < series+1; i++ {
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
					e = e.prev
				}
			}
		}
	}()
	return ch
}
