package a20_test

import (
	"aoc2020/a20"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

const exampleTile = `#.#
.##
..#`

func Test_TileFromString(t *testing.T) {
	tile := a20.TileFromString(exampleTile)
	require.Equal(t, exampleTile, tile.String())
}

const exampleRotated = `..#
.#.
###`

func Test_RotateRight(t *testing.T) {
	tile := a20.TileFromString(exampleTile)
	rot := tile.RotateRight()
	require.Equal(t, exampleRotated, rot.String())

	// verify full rotation
	rot = rot.RotateRight()
	rot = rot.RotateRight()
	rot = rot.RotateRight()
	require.Equal(t, exampleTile, rot.String())
}

const exampleFlipX = `#.#
##.
#..`

func Test_FlipX(t *testing.T) {
	tile := a20.TileFromString(exampleTile)
	rot := tile.FlipX()
	fmt.Println(rot)
	require.Equal(t, exampleFlipX, rot.String())
}

const exampleFlipY = `..#
.##
#.#`

func Test_FlipY(t *testing.T) {
	tile := a20.TileFromString(exampleTile)
	rot := tile.FlipY()
	fmt.Println(rot)
	require.Equal(t, exampleFlipY, rot.String())
}
