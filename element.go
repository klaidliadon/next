package next

// An element of the combination
type element struct {
	prev         *element
	depth, index int
}

// Creates a slice of indexes
func (e *element) list() []int {
	if e.prev == nil {
		return []int{e.index}
	}
	return append(e.prev.list(), e.index)
}

// Returns the actual combination using an element
func (e *element) value(base []interface{}) []interface{} {
	var result = make([]interface{}, 0, e.depth+1)
	for _, v := range e.list() {
		result = append(result, base[v])
	}
	return result
}

func (e *element) join(i int) *element {
	return &element{index: i + 1, depth: e.depth + 1, prev: e}
}
