package a23_test

import (
	"aoc2020/a23"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// var input = "198753462"

func Test_Ring_ShiftRight(t *testing.T) {
	for _, tc := range []struct {
		in      []int
		shifts  int
		wantPos int
	}{
		{[]int{3, 2, 4, 1, 5}, 1, 1},
		{[]int{3, 2, 4, 1, 5}, 3, 3},
		{[]int{3, 2, 4, 1, 5}, 5, 0},
	} {
		t.Run(fmt.Sprintf("in:%+v\tshifts:%v", tc.in, tc.shifts), func(t *testing.T) {
			ring := a23.Ring{
				Items: tc.in,
			}
			ring.ShiftRight(tc.shifts)
			require.Equal(t, tc.wantPos, ring.CurrentPos())
		})
	}
}

func Test_Ring_Remove(t *testing.T) {
	for _, tc := range []struct {
		in          []int
		offset      int
		n           int
		wantErr     error
		wantRemoved []int
		wantRemains []int
	}{
		{[]int{3, 2, 4, 1, 5}, 1, 3, nil, []int{2, 4, 1}, []int{3, 5}},
		{[]int{3, 5}, 1, 2, errors.New("cannot remove all items from the ring"), nil, []int{3, 5}},
		{[]int{3, 5}, 2, 1, errors.New("current position cannot be removed"), nil, []int{3, 5}},
		{[]int{3, 5, 4}, 2, 2, errors.New("current position cannot be removed"), nil, []int{3, 5, 4}},
	} {
		testName := fmt.Sprintf("in:%+v\toffset:%v\tn:%v", tc.in, tc.offset, tc.n)
		t.Run(testName, func(t *testing.T) {
			ring := a23.Ring{
				Items: tc.in,
			}
			gotRemoved, err := ring.Remove(tc.offset, tc.n)
			require.Equal(t, tc.wantErr, err)
			require.Equal(t, tc.wantRemoved, gotRemoved)
			require.Equal(t, tc.wantRemains, ring.Items)
		})
	}
}

func Test_Ring_RemoveWrapAround(t *testing.T) {
	ring := a23.Ring{
		Items: []int{1, 2, 3, 4},
		Pos:   2,
	}
	removed, err := ring.Remove(1, 3)
	require.Nil(t, err)
	require.Equal(t, []int{4, 1, 2}, removed)
	require.Equal(t, []int{3}, ring.Items)
}
