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
	res := make([]int, n)
	nitems := len(r.Items)
	nremoved := 0
	startPos := (r.Pos % offset) % nitems
	removePos := startPos
	for ; nremoved < n; nremoved, removePos = nremoved+1, (removePos+1)%nitems {
		if removePos == r.Pos {
			return nil, errors.New("current position cannot be removed")
		}
		res = append(res, r.Items[removePos])
	}

	// remove slice from start of list
	if startPos <= removePos {
		r.Items = append(r.Items[:startPos], r.Items[removePos+1:]...)
		return res, nil
	}

	// start position is greater than the remove position
	// the wrap-around requires removal from both the head and tail
	r.Items = r.Items[startPos-1:]  // remove tail
	r.Items = r.Items[removePos+1:] // remove head
	return res, nil
}
