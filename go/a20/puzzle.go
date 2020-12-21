package a20

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Puzzle struct {
	tiles       map[TileID]*Tile
	puzzleTiles map[TileID]*PuzzleTile
}

type PuzzleTile struct {
	ID     TileID
	tile   *Tile
	top    *PuzzleTile
	right  *PuzzleTile
	bottom *PuzzleTile
	left   *PuzzleTile
}

func (p Puzzle) String() string {
	var sb strings.Builder
	// if len(p.solution) == 0 {
	for _, tile := range p.tiles {
		sb.WriteRune('\n')
		sb.WriteString(tile.String())
		sb.WriteRune('\n')
	}
	sb.WriteRune('\n')
	return sb.String()
	// }
	// return ""
}

func NewPuzzle(tiles map[TileID]*Tile) Puzzle {
	p := Puzzle{
		tiles:       tiles,
		puzzleTiles: make(map[TileID]*PuzzleTile),
	}
	for tileID, tile := range tiles {
		p.puzzleTiles[tileID] = &PuzzleTile{
			ID:   tileID,
			tile: tile,
		}
	}
	return p
}

// ParsePuzzle parses tiles from the provided reader
func ParsePuzzle(r io.Reader) Puzzle {
	sc := bufio.NewScanner(r)
	tiles := make(map[TileID]*Tile)
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
		tile := TileFromString(tileSB.String())
		tiles[TileID(id)] = &tile
	}

	return NewPuzzle(tiles)
}

func (p *Puzzle) Solve() {
	borderTiles := make(map[uint][]TileID)
	for tileID, tile := range p.tiles {
		for _, border := range BorderValues(tile.Pixels) {
			if _, exists := borderTiles[border]; !exists {
				borderTiles[border] = make([]TileID, 0)
			}
			borderTiles[border] = append(borderTiles[border], tileID)
		}
	}

	// Pick any starting tile
	var anyTileID TileID
	for tileID := range p.puzzleTiles {
		anyTileID = tileID
		break
	}

	done := make(map[TileID]bool)
	todo := map[TileID]bool{anyTileID: true}
	addTodo := func(id TileID) {
		if _, exists := done[id]; !exists {
			todo[id] = true
		}
	}

	for {
		if len(todo) == 0 {
			break
		}
		// pick first tileID in todo
		var curID TileID
		for curID = range todo {
			break
		}

		borders := BorderValues(p.tiles[curID].Pixels)
		// For each border with current rotation / flip
		for j := 0; j < 4; j++ {
			border := borders[2*j]
			// Check if there is a bordering tile ID which is not the current tile
			if len(borderTiles[border]) > 2 {
				panic("ERROR")
			}
			for _, borderingTileID := range borderTiles[border] {
				if borderingTileID == curID {
					continue
				}
				addTodo(p.puzzleTiles[borderingTileID].ID)
				switch j {
				case 0: // top
					if !p.tiles[borderingTileID].Orient(nil, nil, []uint{border}, nil) {
						panic("failed to orient tile")
					}
					p.puzzleTiles[curID].top = p.puzzleTiles[borderingTileID]
				case 1: // right
					if !p.puzzleTiles[borderingTileID].tile.Orient(nil, nil, nil, []uint{border}) {
						panic("failed to orient tile")
					}
					p.puzzleTiles[curID].right = p.puzzleTiles[borderingTileID]
				case 2: // bottom
					if !p.puzzleTiles[borderingTileID].tile.Orient([]uint{border}, nil, nil, nil) {
						panic("failed to orient tile")
					}
					p.puzzleTiles[curID].bottom = p.puzzleTiles[borderingTileID]
				case 3: // left
					if !p.puzzleTiles[borderingTileID].tile.Orient(nil, []uint{border}, nil, nil) {
						panic("failed to orient tile")
					}
					p.puzzleTiles[curID].left = p.puzzleTiles[borderingTileID]
				}
			}
		}

		done[curID] = true
		delete(todo, curID)
	}

	// Find any tile
	for _, puzzleTile := range p.puzzleTiles {
		// Move up until nil
		for {
			if puzzleTile.top == nil {
				break
			}
			puzzleTile = puzzleTile.top
		}
		for ; puzzleTile.left != nil; puzzleTile = puzzleTile.left {
		}
		fmt.Println(puzzleTile.ID)
		fmt.Println(puzzleTile.tile)
		break
	}
}
