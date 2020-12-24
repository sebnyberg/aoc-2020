package a24

import (
	"bufio"
	"log"
	"os"
	"unicode/utf8"
)

func ParseAllMoves(filename string) [][]Direction {
	f, err := os.Open(filename)
	check(err)
	sc := bufio.NewScanner(f)
	tileMoves := make([][]Direction, 0)
	for sc.Scan() {
		tileMoves = append(tileMoves, ParseMoves(sc.Text()))
	}
	return tileMoves
}

func ParseMoves(s string) []Direction {
	var pos int

	accept := func(acceptCh rune) bool {
		ch, width := utf8.DecodeRuneInString(s[pos:])
		if acceptCh == ch {
			pos += width
			return true
		}
		return false
	}
	res := make([]Direction, 0)
	for {
		switch {
		case accept('e'):
			res = append(res, E)
		case accept('s'):
			switch {
			case accept('e'):
				res = append(res, SE)
			case accept('w'):
				res = append(res, SW)
			default:
				log.Fatalln("invalid direction", s[pos:])
			}
		case accept('w'):
			res = append(res, W)
		case accept('n'):
			switch {
			case accept('w'):
				res = append(res, NW)
			case accept('e'):
				res = append(res, NE)
			default:
				log.Fatalln("invalid direction", s[pos:])
			}
		default:
			return res
		}
	}
}
