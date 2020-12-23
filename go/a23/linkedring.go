package a23

import (
	"strconv"
	"strings"
)

type LinkedItem struct {
	Next *LinkedItem
	Prev *LinkedItem
	Val  int
}

type LinkedRing struct {
	First   *LinkedItem
	Len     int
	ItemPos map[int]*LinkedItem
}

func (r *LinkedRing) String() string {
	var sb strings.Builder
	// sb.WriteString(strconv.Itoa(r.First.Val))
	// Print current
	sb.WriteString("(" + strconv.Itoa(r.First.Val) + ")")
	start := r.First.Val
	for cur := r.First.Next; cur.Val != start; cur = cur.Next {
		sb.WriteString(" " + strconv.Itoa(cur.Val))
	}
	return sb.String()
}

// Shift the current position in the ring n times to the right
func (r *LinkedRing) ShiftRight(n int) {
	for i := 0; i < n; i++ {
		r.First = r.First.Next
	}
}

// Rotate ring until we find n
// If n could not be find, the ring will rotate an entire round
func (r *LinkedRing) ShiftTo(n int) bool {
	if item, exists := r.ItemPos[n]; !exists || item == nil {
		return false
	}
	r.First = r.ItemPos[n]
	return true
}

func (r *LinkedRing) Insert(items []int) {
	// Insert each item between the current and next
	// To keep order of inserts, we insert in inverse order
	for i := 0; i < len(items); i++ {
		linkItem := &LinkedItem{Val: items[len(items)-1-i]}
		r.ItemPos[linkItem.Val] = linkItem
		// Edge case
		if r.First == nil {
			r.First = linkItem
			r.First.Next = r.First
			r.First.Prev = r.First
			r.Len++
			continue
		}
		linkItem.Next = r.First.Next
		linkItem.Prev = r.First
		r.First.Next.Prev = linkItem
		r.First.Next = linkItem
		r.Len++
	}
}

func (r *LinkedRing) InsertBefore(item int) {
	linkItem := &LinkedItem{Val: item}
	r.ItemPos[linkItem.Val] = linkItem
	linkItem.Next = r.First
	linkItem.Prev = r.First.Prev
	r.First.Prev.Next = linkItem
	r.First.Prev = linkItem
	r.Len++
}

// Remove n items returning the removed elements
// returns an error if removal would Remove the current position
func (r *LinkedRing) Remove(n int) []int {
	removed := make([]int, n)
	for i := 0; i < n; i++ {
		removed[i] = r.First.Next.Val
		r.ItemPos[r.First.Next.Val] = nil
		r.First.Next.Next.Prev = r.First
		r.First.Next = r.First.Next.Next
		r.Len--
	}
	return removed
}
