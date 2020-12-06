package a03

import (
	"bufio"
	"fmt"
	"io"
)

const rowlen = len("...#...###......##.#..#.....##.")

// Get next position
func NextPos(rowlen int, cur int, right int, down int) int {
	if cur%rowlen+right >= rowlen {
		down--
	}
	cur += down*rowlen + right
	return cur
}

// Count number of valid passwords found in the file
func Run(f io.ReadSeeker) {
	multipliedTrees := 1
	for _, pair := range [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}} {
		f.Seek(0, io.SeekStart)
		right, down := pair[0], pair[1]
		trees := 0
		sc := bufio.NewScanner(f)
		sc.Split(bufio.ScanRunes)
		i := 0
		pos := NextPos(rowlen, 0, right, down)
		for sc.Scan() {
			ch := sc.Text()
			if ch == "\n" {
				continue
			}
			if i == pos {
				pos = NextPos(rowlen, pos, right, down)
				if ch == "#" {
					trees++
				}
			}
			i++
		}
		fmt.Println(trees)
		multipliedTrees *= trees
	}
	fmt.Println(multipliedTrees)
}
