package a20

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_parsePixels(t *testing.T) {
	for _, tc := range []string{"##..#.....", "..##.#..#."} {
		t.Run(tc, func(t *testing.T) {
			pixels := pixelsFromString(tc)
			require.Equal(t, tc, pixels.String())
		})
	}
}

func Test_tileFromString(t *testing.T) {
	for _, tc := range []string{
		`..##.#..#.
##..#.....
#...##..#.
####.#...#
##.##.###.
##...#.###
.#.#.#..##
..#....#..
###...#.#.
..###..###`,
		`#.##...##.
#.####...#
.....#..##
#...######
.##.#....#
.###.#####
###.##.##.
.###....#.
..#.#..#.#
#...##.#..`,
	} {
		t.Run(tc, func(t *testing.T) {
			tile := tileFromString(tc)
			require.Equal(t, tc, tile.String())
		})
	}
}

// func Test_rotateRight(t *testing.T) {
// 	in := tile{
// 		rows: []pixels{
// 			pixelsFromString("##..#....."),
// 	}
// 	for _, tc := range []struct{
// 		in tile
// 		expected []tile
// 	}
// }
