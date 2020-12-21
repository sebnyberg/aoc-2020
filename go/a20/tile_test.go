package a20_test

import (
	"aoc2020/a20"
	"log"
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

func Test_Tile_RotateRight(t *testing.T) {
	tile := a20.TileFromString(exampleTile)
	tile.RotateRight()
	require.Equal(t, exampleRotated, tile.String())

	// verify full rotation
	tile.RotateRight()
	tile.RotateRight()
	tile.RotateRight()
	require.Equal(t, exampleTile, tile.String())
}

const exampleFlipX = `#.#
##.
#..`

func Test_Tile_FlipX(t *testing.T) {
	tile := a20.TileFromString(exampleTile)
	tile.FlipX()
	require.Equal(t, exampleFlipX, tile.String())
}

const exampleFlipY = `..#
.##
#.#`

func Test_Tile_FlipY(t *testing.T) {
	tile := a20.TileFromString(exampleTile)
	tile.FlipY()
	require.Equal(t, exampleFlipY, tile.String())
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func Test_Tile_Borders(t *testing.T) {
	// ..#
	// #..
	// ###
	tile := a20.TileFromString("..#\n#..\n###")
	got := a20.BorderValues(tile.Pixels)
	want := [8]uint{1, 4, 5, 5, 7, 7, 3, 6}
	require.Equal(t, got, want)
}
