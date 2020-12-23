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

// Find the provided item, returning its offset from the current index
// returns -1 if the item could not be found
func (r *Ring) Find(item int) (idx int) {
	nitems := len(r.Items)
	for i := 0; i < len(r.Items); i++ {
		currentPos := (r.Pos + i) % nitems
		if r.Items[currentPos] == item {
			return i
		}
	}
	return -1
}

// Return the CurrentItem item in the ring
func (r *Ring) CurrentItem() int {
	return r.Items[r.Pos]
}

// Return the Current position in the ring
func (r *Ring) CurrentPos() int {
	return r.Pos
}

func (r *Ring) Insert(items []int, offset int) {
	insertPos := (r.Pos + offset) % len(r.Items)
	// edge case - insert at the end
	if insertPos == 0 {
		r.Items = append(r.Items, items...)
		return
	}
	// make room for items (not in correct location yet)
	r.Items = append(r.Items, items...)

	// copy tail to correct location (at the end)
	copy(r.Items[insertPos+len(items):], r.Items[insertPos:])

	// copy items to correct location (in the middle)
	copy(r.Items[insertPos:], items)

	// adjust position to account for new items
	if insertPos <= r.Pos {
		r.Pos += len(items)
	}
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
