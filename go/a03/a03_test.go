package a03_test

import (
	"aoc2020/a03"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_SkiSlope(t *testing.T) {
	rowlen := 3
	for _, tc := range []struct {
		right int
		down  int
		want  []int
	}{
		{1, 1, []int{0, rowlen + 1, rowlen*2 + 2, rowlen * 3, rowlen*4 + 1}},
		{1, 2, []int{0, rowlen*2 + 1, rowlen*4 + 2, rowlen * 6, rowlen*8 + 1}},
	} {
		t.Run(fmt.Sprintf("right: %v, down: %v", tc.right, tc.down), func(t *testing.T) {
			var cur int
			for _, expected := range tc.want {
				require.Equal(t, expected, cur)
				cur = a03.NextPos(rowlen, cur, tc.right, tc.down)
			}
		})
	}
}
