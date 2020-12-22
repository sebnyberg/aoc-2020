package a20_test

import (
	"aoc2020/a20"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Day20(t *testing.T) {
	f, err := os.Open("input")
	check(err)
	puzzle := a20.ParsePuzzle(f)
	puzzle.Solve()
	// puzzle.PrintTileIDs()
	tile := a20.TileFromString(puzzle.String())
	tile.FlipX()
	for i := 0; i < 8; i++ {
		if i == 4 {
			tile.FlipX()
		}
		if findMonster(tile.Pixels) > 0 {
			break
		}
		tile.RotateRight()
	}
	monsterHashes := trueCount(parseCanvas(monster))
	nmonsters := findMonster(tile.Pixels)
	fmt.Println(monsterHashes * nmonsters)
	fmt.Println(trueCount(tile.Pixels))
	require.Equal(t, 2155, trueCount(tile.Pixels)-nmonsters*trueCount(parseCanvas(monster)))
}

func trueCount(in [][]bool) (res int) {
	for i := range in {
		for j := range in[i] {
			if in[i][j] {
				res++
			}
		}
	}
	return res
}

var monster = `                  # 
#    ##    ##    ###
 #  #  #  #  #  #   `

var doubleMonster = `                  #     # 
#    ##    ##    ###   ###
 #  #  #  #  #  #  #  #   `

var monsterTest = `.#.#..#.##...#.##..#####
###....#.#....#..#......
##.##.###.#.#..######...
###.#####...#.#####.#..#
##.#....#.##.####...#.##
...########.#....#####.#
....#..#...##..#.#.###..
.####...#..#.....#......
#..#.##..#..###.#.##....
#.####..#.####.#.#.###..
###.#.#...#.######.#..##
#.####....##..########.#
##..##.#...#...#.#.#.#..
...#..#..#.#.##..###.###
.#.#....#.##.#...###.##.
###.#...#..#.##.######..
.#.#.###.##.##.#..#.##..
.####.###.#...###.#..#.#
..#.#..#..#.#.#.####.###
#..####...#.#.#.###.###.
#####..#####...###....##
#.##..#..#...#..####...#
.#.###..##..##..####.##.
...###...##...#...#..###`

// func Test_findMonster(t *testing.T) {
// 	// canvas := monster
// 	// require.Equal(t, 1, findMonster(parseCanvas(canvas)))
// 	// require.Equal(t, 2, findMonster(parseCanvas(doubleMonster)))
// 	testTile := a20.TileFromString(monsterTest)
// 	testTile.FlipX()
// 	testTile.RotateRight()
// 	testTile.RotateRight()
// 	testTile.RotateRight()
// 	require.Equal(t, 2, findMonster(testTile.Pixels))
// }

func parseCanvas(s string) [][]bool {
	rows := strings.Split(s, "\n")
	res := make([][]bool, len(rows))
	for i, row := range rows {
		res[i] = make([]bool, len(row))
		for j, ch := range row {
			if ch == '#' {
				res[i][j] = true
			}
		}
	}
	return res
}

func findMonster(canvas [][]bool) (nfound int) {
	monsterBools := parseCanvas(monster)
	// fmt.Println(boolsString(monsterBools))
	if len(canvas) < 3 {
		return 0
	}
	for row := 0; row < len(canvas)-2; row++ {
		for col := 0; col < len(canvas[row])-19; col++ {
			// fmt.Println("--->")
			checkArea := make([][]bool, 3)
			for i := range checkArea {
				checkArea[i] = canvas[row+i][col : col+20]
			}
			matched := true
			for checkRow := range checkArea {
				for checkCol := range checkArea[checkRow] {
					if matched && monsterBools[checkRow][checkCol] {
						if !checkArea[checkRow][checkCol] {
							matched = false
						}
					}
				}
			}
			if matched {
				// fmt.Println("FOUND!!")
				// fmt.Println("row", row, "col", col)
				// fmt.Println("----------")
				// fmt.Println(boolsString(checkArea))
				// fmt.Println(boolsString(monsterBools))
				// fmt.Println("FOUND!!")
				nfound++
			}
			// fmt.Println(boolsString(checkArea))
		}
		// fmt.Println("---")
		// fmt.Println("DOWN")
		// fmt.Println("---")
	}
	return nfound
}

func boolsString(bools [][]bool) string {
	var sb strings.Builder
	for row := 0; row < len(bools); row++ {
		for col := 0; col < len(bools[row]); col++ {
			if bools[row][col] {
				sb.WriteRune('#')
			} else {
				sb.WriteRune('.')
			}
		}
		sb.WriteRune('\n')
	}
	return strings.TrimRight(sb.String(), "\n")
}
