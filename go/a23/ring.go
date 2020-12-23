package a23

import "errors"

type Ring struct {
	Items []int
	Pos   int
}

// Shift the current position in the ring n times to the right
func (r *Ring) ShiftRight(n int) {
	r.Pos += n
	r.Pos %= len(r.Items)
}

// Return the CurrentItem item in the ring
func (r *Ring) CurrentItem() int {
	return r.Items[r.Pos]
}

// Return the Current position in the ring
func (r *Ring) CurrentPos() int {
	return r.Pos
}

// Remove n items starting at the provided offset
// returning the removed elements
// returns an error if removal would Remove the current position
func (r *Ring) Remove(offset int, n int) ([]int, error) {
	if offset <= 0 {
		return nil, errors.New("offset must be greater than zero")
	}
	if n >= len(r.Items) {
		return nil, errors.New("cannot remove all items from the ring")
	}
	res := make([]int, 0, n)
	nitems := len(r.Items)
	nremoved := 0
	startPos := (r.Pos + offset) % nitems

	for removePos := (r.Pos + offset) % nitems; nremoved < n; nremoved++ {
		if removePos == r.Pos {
			return nil, errors.New("current position cannot be removed")
		}
		res = append(res, r.Items[removePos])
		removePos++
		removePos %= nitems
	}

	// remove slice from start of list
	finalRemovedPos := (r.Pos + offset + n - 1) % nitems
	if startPos <= finalRemovedPos {
		r.Items = append(r.Items[:startPos], r.Items[finalRemovedPos+1:]...)
		return res, nil
	}

	// start position is greater than the remove position
	// the wrap-around requires removal from both the head and tail
	r.Items = r.Items[finalRemovedPos+1 : startPos]
	return res, nil
}
