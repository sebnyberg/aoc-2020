package a24

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_a24(t *testing.T) {
	floor := NewFloor()
	tileMoves := ParseAllMoves("input")
	for _, moves := range tileMoves {
		floor.FlipTile(moves)
	}
	for i := 0; i < 100; i++ {
		floor.NextDay()
	}
}

func Test_tileIDMove(t *testing.T) {
	id := TileID{0, 0}
	id = id.Move(NE)
	require.Equal(t, []int{0, 1}, []int{id.XDiag, id.Y})
	id = id.Move(NE)
	require.Equal(t, []int{0, 2}, []int{id.XDiag, id.Y})
}

func Test_flipTiles(t *testing.T) {
	floor := NewFloor()
	floor.FlipTile(ParseMoves("neseswwnwne"))
	require.Equal(t, 1, floor.CountBlackTiles())
	floor.FlipTile(ParseMoves("nw"))
	require.Equal(t, 0, floor.CountBlackTiles())
}

func Test_parseMoves(t *testing.T) {
	for _, tc := range []struct {
		in   string
		want []Direction
	}{
		{"sesenwne", []Direction{SE, SE, NW, NE}},
	} {
		t.Run(tc.in, func(t *testing.T) {
			require.Equal(t, tc.want, ParseMoves(tc.in))
		})
	}
}
