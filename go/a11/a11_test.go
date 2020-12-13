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

func Test_day11(t *testing.T) {
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
		after := simulateArrival(cur)
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

	require.Equal(t, 0, noccupied)
}

func simulateArrival(before [][]rune) [][]rune {
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
