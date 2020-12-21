package a20

import (
	"log"
	"strings"
)

type TileID string

type rotation int

const (
	rotationNone rotation = iota
	rotationRight
	rotationDown
	rotationLeft
)

type Tile struct {
	width    int
	Pixels   [][]bool
	Flipped  bool
	Rotation rotation
}

func NewTile(width int) Tile {
	t := Tile{
		Pixels: make([][]bool, width),
		width:  width,
	}
	for i := range t.Pixels {
		t.Pixels[i] = make([]bool, width)
	}
	return t
}

func TileFromString(s string) Tile {
	rows := strings.Split(s, "\n")
	var t Tile
	t.width = len(rows[0])

	// Parse pixels
	t.Pixels = make([][]bool, t.width)
	for i, row := range rows {
		t.Pixels[i] = make([]bool, t.width)
		for j, ch := range row {
			if ch == '#' {
				t.Pixels[i][j] = true
			}
		}
	}
	return t
}

func (t Tile) String() string {
	var sb strings.Builder
	for i := range t.Pixels {
		for j := range t.Pixels[i] {
			if t.Pixels[i][j] {
				sb.WriteRune('#')
				continue
			}
			sb.WriteRune('.')
		}
		if i < len(t.Pixels)-1 {
			sb.WriteRune('\n')
		}
	}
	return sb.String()
}

func (t *Tile) RotateRight() {
	if t.Flipped {
		t.Rotation = t.Rotation - 1
	} else {
		t.Rotation = t.Rotation + 1
	}
	t.Rotation %= 4
	newPixels := make([][]bool, t.width)
	for i := range t.Pixels {
		newPixels[i] = make([]bool, t.width)
	}

	for i := range t.Pixels {
		for j := range t.Pixels[i] {
			newPixels[j][t.width-1-i] = t.Pixels[i][j]
		}
	}
	t.Pixels = newPixels
}

func (t *Tile) FlipX() {
	t.Flipped = !t.Flipped
	newPixels := make([][]bool, t.width)
	for i := range t.Pixels {
		newPixels[i] = make([]bool, t.width)
	}
	for i := range t.Pixels {
		for j := range t.Pixels[i] {
			newPixels[i][t.width-1-j] = t.Pixels[i][j]
		}
	}
	t.Pixels = newPixels
}

func (t *Tile) FlipY() {
	t.FlipX()
	t.RotateRight()
	t.RotateRight()
}

type Borders struct {
	Top           uint
	TopFlipped    uint
	Right         uint
	RightFlipped  uint
	Bottom        uint
	BottomFlipped uint
	Left          uint
	LeftFlipped   uint
}

// Return (possible) borders for the tile
func BorderValues(pixels [][]bool) (res [8]uint) {
	height := len(pixels)
	width := len(pixels[0])
	if height != width {
		log.Fatalln("grids are expected to be NxN")
	}

	// create columns
	firstCol := make([]bool, height)
	lastCol := make([]bool, height)
	for i := range pixels {
		firstCol[i] = pixels[i][0]
		lastCol[i] = pixels[i][height-1]
	}

	res[0] = BoolSliceToUint(pixels[0])
	res[1] = BoolSliceToUint(BoolSliceReverse(pixels[0]))
	res[2] = BoolSliceToUint(lastCol)
	res[3] = BoolSliceToUint(BoolSliceReverse(lastCol))
	res[4] = BoolSliceToUint(pixels[width-1])
	res[5] = BoolSliceToUint(BoolSliceReverse(pixels[width-1]))
	res[6] = BoolSliceToUint(firstCol)
	res[7] = BoolSliceToUint(BoolSliceReverse(firstCol))

	return
}

func (t *Tile) MatchBorders(topBorders []uint, rightBorders []uint, bottomBorders []uint, leftBorders []uint) bool {
	matches := func(borders [8]uint) bool {
		if len(topBorders) == 2 {
			// Either must match
			if borders[0] != topBorders[0] && borders[0] != topBorders[1] {
				return false
			}
		} else if len(topBorders) == 1 {
			if borders[0] != topBorders[0] {
				return false
			}
		}

		if len(rightBorders) == 2 {
			// Either must match
			if borders[2] != rightBorders[0] && borders[2] != rightBorders[1] {
				return false
			}
		} else if len(rightBorders) == 1 {
			if borders[2] != rightBorders[0] {
				return false
			}
		}

		if len(bottomBorders) == 2 {
			// Either must match
			if borders[4] != bottomBorders[0] && borders[4] != bottomBorders[1] {
				return false
			}
		} else if len(bottomBorders) == 1 {
			if borders[4] != bottomBorders[0] {
				return false
			}
		}

		if len(leftBorders) == 2 {
			// Either must match
			if borders[6] != leftBorders[0] && borders[0] != leftBorders[1] {
				return false
			}
		} else if len(leftBorders) == 1 {
			if borders[6] != leftBorders[0] {
				return false
			}
		}
		return true
	}

	for i := 0; i < 4; i++ {
		if matches(BorderValues(t.Pixels)) {
			return true
		}
		t.RotateRight()
	}

	t.FlipX()
	for i := 0; i < 4; i++ {
		if matches(BorderValues(t.Pixels)) {
			return true
		}
		t.RotateRight()
	}

	return false
}

func BoolSliceReverse(in []bool) []bool {
	res := make([]bool, len(in))
	for i := range in {
		res[len(in)-1-i] = in[i]
	}
	return res
}

func BoolSliceToUint(in []bool) (res uint) {
	for _, b := range in {
		res <<= 1
		if b {
			res |= 1
		}
	}
	return res

}
