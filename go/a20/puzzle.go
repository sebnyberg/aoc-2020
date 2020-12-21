package a20

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
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

	// Find corner pieces
	edgeTileEdges := make(map[TileID][]uint)
	for tileID, tile := range p.tiles {
		for _, tileBorder := range BorderValues(tile.Pixels) {
			switch len(borderTiles[tileBorder]) {
			case 1:
				if _, exists := edgeTileEdges[tileID]; !exists {
					edgeTileEdges[tileID] = make([]uint, 0)
				}
				edgeTileEdges[tileID] = append(edgeTileEdges[tileID], tileBorder)
			case 2:
			default:
				log.Fatalf("invalid number of matching edges: %v", len(borderTiles[tileBorder]))
			}
		}
	}
	sum := 1
	for tileID, edges := range edgeTileEdges {
		if len(edges) == 4 {
			n, _ := strconv.Atoi(string(tileID))
			sum *= n
		}
	}

	fmt.Println(sum)
}

// func (p *Puzzle) Solve() bool {

// 	return true
// }
