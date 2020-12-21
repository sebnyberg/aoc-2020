package a20

import "strings"

type Tile struct {
	width  int
	pixels [][]bool
}

func NewTile(width int) Tile {
	t := Tile{
		pixels: make([][]bool, width),
		width:  width,
	}
	for i := range t.pixels {
		t.pixels[i] = make([]bool, width)
	}
	return t
}

func TileFromString(s string) Tile {
	rows := strings.Split(s, "\n")
	var t Tile
	t.width = len(rows[0])
	t.pixels = make([][]bool, t.width)
	for i, row := range rows {
		t.pixels[i] = make([]bool, t.width)
		for j, ch := range row {
			if ch == '#' {
				t.pixels[i][j] = true
			}
		}
	}
	return t
}

func (t Tile) String() string {
	var sb strings.Builder
	for i := range t.pixels {
		for j := range t.pixels[i] {
			if t.pixels[i][j] {
				sb.WriteRune('#')
				continue
			}
			sb.WriteRune('.')
		}
		if i < len(t.pixels)-1 {
			sb.WriteRune('\n')
		}
	}
	return sb.String()
}

func (t Tile) RotateRight() Tile {
	rot := NewTile(t.width)
	for i := range t.pixels {
		for j := range t.pixels[i] {
			rot.pixels[j][t.width-1-i] = t.pixels[i][j]
		}
	}
	return rot
}

func (t Tile) FlipX() Tile {
	flip := NewTile(t.width)
	for i := range t.pixels {
		for j := range t.pixels[i] {
			flip.pixels[i][t.width-1-j] = t.pixels[i][j]
		}
	}
	return flip
}

func (t Tile) FlipY() Tile {
	flip := NewTile(t.width)
	for i := range t.pixels {
		for j := range t.pixels[i] {
			flip.pixels[t.width-1-i][j] = t.pixels[i][j]
		}
	}
	return flip
}
