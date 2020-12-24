package a24

import (
	"log"
	"strconv"
)

type Floor struct {
	BlackTiles map[TileID]struct{}
}

func NewFloor() *Floor {
	f := &Floor{
		BlackTiles: make(map[TileID]struct{}),
	}
	return f
}

func (f *Floor) FlipTile(moves []Direction) {
	tileID := TileID{0, 0}
	for _, move := range moves {
		tileID = tileID.Move(move)
	}
	if _, exists := f.BlackTiles[tileID]; !exists {
		f.BlackTiles[tileID] = struct{}{}
	} else {
		delete(f.BlackTiles, tileID)
	}
}

func (f *Floor) NextDay() {
	nextBlack := make(map[TileID]struct{}, len(f.BlackTiles))
	blackNeighbours := make(map[TileID]struct{}, len(f.BlackTiles))
	var adjblack int
	for tileID := range f.BlackTiles {
		adjblack = 0
		for _, adjTileID := range tileID.Adj() {
			if _, isBlack := f.BlackTiles[adjTileID]; isBlack {
				adjblack++
				continue
			}
			blackNeighbours[adjTileID] = struct{}{}
		}
		if adjblack == 1 || adjblack == 2 {
			nextBlack[tileID] = struct{}{}
		}
	}

	for tileID := range blackNeighbours {
		adjblack = 0
		for _, adjTileID := range tileID.Adj() {
			if _, isBlack := f.BlackTiles[adjTileID]; isBlack {
				adjblack++
				continue
			}
		}
		if adjblack == 2 {
			nextBlack[tileID] = struct{}{}
		}
	}

	f.BlackTiles = nextBlack
}

func (f *Floor) CountBlackTiles() (res int) {
	for range f.BlackTiles {
		res++
	}
	return res
}

type TileColor int

const (
	White = iota
	Black
)

func (c TileColor) String() string {
	if c == White {
		return "White"
	} else {
		return "Black"
	}
}

type Direction int

const (
	E Direction = iota
	NE
	NW
	W
	SW
	SE
)

var allDirections = []Direction{E, SE, SW, W, NW, NE}

type TileID struct {
	XDiag int // forward-slash diagonal
	Y     int
}

func (id TileID) String() string {
	return "{" + strconv.Itoa(id.XDiag) + "," + strconv.Itoa(id.Y) + "}"
}

func (id TileID) Adj() [6]TileID {
	return [6]TileID{
		{id.XDiag + 1, id.Y},
		{id.XDiag - 1, id.Y},
		{id.XDiag + 1, id.Y - 1},
		{id.XDiag - 1, id.Y + 1},
		{id.XDiag, id.Y - 1},
		{id.XDiag, id.Y + 1},
	}
}

func (id TileID) Move(dir Direction) TileID {
	switch dir {
	case E:
		return TileID{id.XDiag + 1, id.Y}
	case SE:
		return TileID{id.XDiag + 1, id.Y - 1}
	case SW:
		return TileID{id.XDiag, id.Y - 1}
	case W:
		return TileID{id.XDiag - 1, id.Y}
	case NW:
		return TileID{id.XDiag - 1, id.Y + 1}
	case NE:
		return TileID{id.XDiag, id.Y + 1}
	}
	panic("invalid tile ID move")
}

func (d Direction) String() string {
	switch d {
	case E:
		return "e"
	case NE:
		return "ne"
	case NW:
		return "nw"
	case W:
		return "w"
	case SW:
		return "sw"
	case SE:
		return "se"
	}
	return "INVALID_DIRECTION"
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
