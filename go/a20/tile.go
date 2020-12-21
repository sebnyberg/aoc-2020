package a20

import "strings"

type pixels uint16

func pixelsFromString(s string) (p pixels) {
	for i, ch := range s {
		if i >= 10 {
			panic("max length for a pixel is 10")
		}
		p <<= 1
		if ch == '#' {
			p |= 1
		}
	}
	return p
}

func (p pixels) String() string {
	var res [10]rune
	for i := 0; i < 10; i++ {
		if p%2 == 0 {
			res[9-i] = '.'
		} else {
			res[9-i] = '#'
		}
		p >>= 1
	}
	return string(res[:])
}

type tile struct {
	rows []pixels
	cols []pixels
}

func tileFromString(s string) tile {
	rows := strings.Split(s, "\n")
	if len(rows) != 10 {
		panic("tile needs exactly 10 rows")
	}
	t := tile{
		rows: make([]pixels, 10),
		cols: make([]pixels, 10),
	}

	var colStrs [10][10]rune

	for i, row := range rows {
		for j, ch := range row {
			colStrs[j][i] = ch
		}
		t.rows[i] = pixelsFromString(row)
	}

	for i, colStr := range colStrs {
		t.cols[i] = pixelsFromString(string(colStr[:]))
	}

	return t
}

func (t tile) String() string {
	var sb strings.Builder
	sb.WriteString(t.rows[0].String())
	for _, row := range t.rows[1:] {
		sb.WriteRune('\n')
		sb.WriteString(row.String())
	}
	return sb.String()
}

func reverse(in []pixels) []pixels {
	// Reverse rows
	for i, j := 0, len(in)-1; i < j; i, j = i+1, j-1 {
		in[i], in[j] = in[j], in[i]
	}
	return in
}

func (t tile) rotateRight() {
	t.rows, t.cols = t.cols, reverse(t.rows)
}
