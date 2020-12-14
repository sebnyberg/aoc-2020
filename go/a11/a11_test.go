package a11_test

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func Test_day11_part1(t *testing.T) {
	// 	input := `L.LL.LL.LL
	// LLLLLLL.LL
	// L.L.L..L..
	// LLLL.LL.LL
	// L.LL.LL.LL
	// L.LLLLL.LL
	// ..L.L.....
	// LLLLLLLLLL
	// L.LLLLLL.L
	// L.LLLLL.LL`
	f, err := os.Open("input")
	check(err)

	// 	input := `L.
	// LL`
	// f := bytes.NewBufferString(input)
	sc := bufio.NewScanner(f)

	seats := make([][]rune, 0)
	for sc.Scan() {
		seats = append(seats, []rune(sc.Text()))
	}

	nochange := func(before, after [][]rune) bool {
		for i := range before {
			for j := range before[i] {
				if before[i][j] != after[i][j] {
					return false
				}
			}
		}
		return true
	}

	cur := seats
	i := 0
	for {
		after := simulateArrival1(cur)
		if nochange(cur, after) {
			break
		}
		cur = after
		i++
	}

	noccupied := 0
	for _, row := range cur {
		for _, seat := range row {
			if seat == '#' {
				noccupied++
			}
		}
	}

	require.Equal(t, 2277, noccupied)
}
func Test_day11_part2(t *testing.T) {
	// 	input := `L.LL.LL.LL
	// LLLLLLL.LL
	// L.L.L..L..
	// LLLL.LL.LL
	// L.LL.LL.LL
	// L.LLLLL.LL
	// ..L.L.....
	// LLLLLLLLLL
	// L.LLLLLL.L
	// L.LLLLL.LL`
	// f := bytes.NewBufferString(input)
	f, err := os.Open("input")
	check(err)

	// 	input := `L.
	// LL`
	sc := bufio.NewScanner(f)

	seats := make([][]rune, 0)
	for sc.Scan() {
		seats = append(seats, []rune(sc.Text()))
	}

	nochange := func(before, after [][]rune) bool {
		for i := range before {
			for j := range before[i] {
				if before[i][j] != after[i][j] {
					return false
				}
			}
		}
		return true
	}

	cur := seats
	i := 0
	for {
		after := simulateArrival2(cur)
		if nochange(cur, after) {
			break
		}
		cur = after
		i++
	}

	noccupied := 0
	for _, row := range cur {
		for _, seat := range row {
			if seat == '#' {
				noccupied++
			}
		}
	}

	require.Equal(t, 2066, noccupied)
}

func simulateArrival2(before [][]rune) [][]rune {
	debugRow, debugCol := 3, 2
	debug := false

	after := make([][]rune, len(before))
	// adj := make([]rune, 0, 9)
	for row := range before {
		after[row] = make([]rune, len(before[row]))
		for col := range before[row] {
			// Floor - no change
			if before[row][col] == '.' {
				after[row][col] = '.'
				continue
			}
			after[row][col] = before[row][col]

			nadjoccupied := 0

			// Go upwards left until we find a seat or edge of the row
			for i := 1; row-i >= 0 && col-i >= 0; i++ {
				if before[row-i][col-i] == '#' {
					nadjoccupied++
					break
				}
				if before[row-i][col-i] == 'L' {
					break
				}
			}
			if debug && row == debugRow && col == debugCol {
				fmt.Print(nadjoccupied)
			}

			// Go up until we find a seat or edge of the row
			for i := row - 1; i >= 0; i-- {
				if before[i][col] == '#' {
					nadjoccupied++
					break
				}
				if before[i][col] == 'L' {
					break
				}
			}
			if debug && row == debugRow && col == debugCol {
				fmt.Print(nadjoccupied)
			}

			// Go upwards right until we find a seat or edge of the row
			for i := 1; row-i >= 0 && col+i < len(before[row]); i++ {
				if before[row-i][col+i] == '#' {
					nadjoccupied++
					break
				}
				if before[row-i][col+i] == 'L' {
					break
				}
			}
			if debug && row == debugRow && col == debugCol {
				fmt.Printf("%v\n", nadjoccupied)
			}

			// Go left until we find a seat or the edge of the row
			for i := col - 1; i >= 0; i-- {
				if before[row][i] == '#' {
					nadjoccupied++
					break
				}
				if before[row][i] == 'L' {
					break
				}
			}
			if debug && row == debugRow && col == debugCol {
				fmt.Printf("%vX", nadjoccupied)
			}

			// Go right until we find a seat or the edge of the row
			for i := col + 1; i < len(before[row]); i++ {
				if before[row][i] == '#' {
					nadjoccupied++
					break
				}
				if before[row][i] == 'L' {
					break
				}
			}
			if debug && row == debugRow && col == debugCol {
				fmt.Printf("%v\n", nadjoccupied)
			}

			// Go downwards left until we find a seat or edge of the row
			for i := 1; row+i < len(before) && col-i >= 0; i++ {
				if before[row+i][col-i] == '#' {
					nadjoccupied++
					break
				}
				if before[row+i][col-i] == 'L' {
					break
				}
			}
			if debug && row == debugRow && col == debugCol {
				fmt.Print(nadjoccupied)
			}

			// Go down until we find a seat or the edge of the row
			for i := row + 1; i < len(before); i++ {
				if before[i][col] == '#' {
					nadjoccupied++
					break
				}
				if before[i][col] == 'L' {
					break
				}
			}
			if debug && row == debugRow && col == debugCol {
				fmt.Print(nadjoccupied)
			}

			// Go downwards right until we find a seat or edge of the row
			for i := 1; row+i < len(before) && col+i < len(before[row]); i++ {

				if before[row+i][col+i] == '#' {
					nadjoccupied++
					break
				}
				if before[row+i][col+i] == 'L' {
					break
				}
			}
			if debug && row == debugRow && col == debugCol {
				fmt.Printf("%v\n", nadjoccupied)
			}

			switch before[row][col] {
			case 'L':
				if nadjoccupied == 0 {
					after[row][col] = '#'
				}
			case '#':
				if nadjoccupied >= 5 {
					after[row][col] = 'L'
				}
			}
		}
	}

	return after
}

func simulateArrival1(before [][]rune) [][]rune {
	after := make([][]rune, len(before))
	// adj := make([]rune, 0, 9)
	for row := range before {
		after[row] = make([]rune, len(before[row]))
		for col := range before[row] {
			// Floor - no change
			if before[row][col] == '.' {
				after[row][col] = '.'
				continue
			}
			after[row][col] = before[row][col]

			// Fetch adjecent seats
			// adj = adj[:0]

			adjrowStart := max(0, row-1)
			adjrowEnd := min(len(before), row+2)
			adjcolStart := max(0, col-1)
			adjcolEnd := min(len(before[row]), col+2)
			nadjoccupied := 0

			for i, adjrow := range before[adjrowStart:adjrowEnd] {
				for j, seat := range adjrow[adjcolStart:adjcolEnd] {
					// Skip current seat when collecting adjecent
					if adjrowStart+i == row && adjcolStart+j == col {
						continue
					}
					if seat == '#' {
						nadjoccupied++
					}
					// adj = append(adj, seat)
				}
			}

			switch before[row][col] {
			case 'L':
				if nadjoccupied == 0 {
					after[row][col] = '#'
				}
			case '#':
				if nadjoccupied >= 4 {
					after[row][col] = 'L'
				}
			}
		}
	}

	return after
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func printSeats(seats [][]rune) {
	for _, row := range seats {
		for _, ch := range row {
			fmt.Print(string(ch))
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}
