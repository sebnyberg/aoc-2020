package a20

import (
	"bufio"
	"io"
	"strings"
)

type Puzzle struct {
	tiles    map[TileID]Tile
	solution [][]string
}

func (p Puzzle) String() string {
	var sb strings.Builder
	if len(p.solution) == 0 {
		for _, tile := range p.tiles {
			sb.WriteRune('\n')
			sb.WriteString(tile.String())
			sb.WriteRune('\n')
		}
		sb.WriteRune('\n')
		return sb.String()
	}
	return ""
}

func NewPuzzle(tiles map[TileID]Tile) Puzzle {
	p := Puzzle{
		tiles: tiles,
	}
	return p
}

// ParsePuzzle parses tiles from the provided reader
func ParsePuzzle(r io.Reader) Puzzle {
	sc := bufio.NewScanner(r)
	tiles := make(map[TileID]Tile)
	for sc.Scan() {
		p := strings.Split(sc.Text(), " ")
		id := strings.TrimRight(p[1], ":")
		var tileSB strings.Builder
		sc.Scan()
		tileSB.WriteString(sc.Text())
		for sc.Scan() {
			row := sc.Text()
			if row == "" {
				break
			}
			tileSB.WriteRune('\n')
			tileSB.WriteString(row)
		}
		tiles[TileID(id)] = TileFromString(tileSB.String())
	}

	return NewPuzzle(tiles)
}

// func FindCornerTiles(tiles []Tile) (res [4]Tile) {
// Corner tiles are tiles which have two borders that are not
// shared with other tiles
// }

// func (p *Puzzle) Solve() bool {

// 	return true
// }
