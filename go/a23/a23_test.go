package a23_test

import (
	"aoc2020/a23"
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
