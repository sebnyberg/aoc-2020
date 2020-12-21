package a20

import (
	"bufio"
	"fmt"
	"io"
	"log"
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

func (p *Puzzle) Solve() {
	// Create a map of border values -> tile
	borderTiles := make(map[uint][]TileID)
	for tileID, tile := range p.tiles {
		for _, border := range BorderValues(tile.Pixels) {
			if _, exists := borderTiles[border]; !exists {
				borderTiles[border] = make([]TileID, 0, 1)
			}
			borderTiles[border] = append(borderTiles[border], tileID)
		}
	}

	// List shared borders for each tile
	tileSharedBorders := make(map[TileID][]uint)
	for tileID, tile := range p.tiles {
		// For each border around the tile
		for _, tileBorder := range BorderValues(tile.Pixels) {
			// Check how many tiles share that border
			switch len(borderTiles[tileBorder]) {
			case 1: // This tile is the only one for this border - it's an edge
			case 2: // Two tiles share this border, it's a shared edge
				if _, exists := tileSharedBorders[tileID]; !exists {
					tileSharedBorders[tileID] = make([]uint, 0)
				}
				tileSharedBorders[tileID] = append(tileSharedBorders[tileID], tileBorder)
			default: // There should only be 1 or 2 tiles for a given border
				log.Fatalf("invalid number of matching edges: %v", len(borderTiles[tileBorder]))
			}
		}
	}

	cornerTiles := make([]TileID, 0)
	edgeTiles := make([]TileID, 0)
	innerTiles := make([]TileID, 0)
	for tileID, edges := range tileSharedBorders {
		switch len(edges) {
		case 4:
			cornerTiles = append(cornerTiles, tileID)
		case 6:
			edgeTiles = append(edgeTiles, tileID)
		case 8:
			innerTiles = append(innerTiles, tileID)
		default:
			log.Fatalf("invalid number of tiles: %v", len(edges))
		}
	}

	// Pick a corner tile, flip + rotate until its edges are top/left
	fmt.Println("corner tiles", cornerTiles, len(cornerTiles))
	fmt.Println("edge tiles", edgeTiles, len(edgeTiles))
	fmt.Println("inner tiles", innerTiles, len(innerTiles))

	// cornerTile := p.tiles[cornerTiles[0]]
	// cornerTileSharedBorders := tileSharedBorders[cornerTiles[0]]
}

// func (p *Puzzle) Solve() bool {

// 	return true
// }
