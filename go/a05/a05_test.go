package a05_test

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

var input = mustReadAll("input")

var exampleInput = `BFFFBBFRRR
FFFBBBFRRR
BBFFBBFRLL`

func Test_findMissingSeat(t *testing.T) {
	for i, tc := range []struct {
		in   []int
		want int
	}{
		{[]int{1, 3}, 2},
		{[]int{1, 2, 3, 4, 7}, 5},
		{getSeatIDs(bytes.NewBufferString(string(input))), 696},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			require.Equal(t, tc.want, findMissingSeat(tc.in))
		})
	}
}

// Returns -1 if no seat was found
func findMissingSeat(seatIDs []int) int {
	if len(seatIDs) <= 1 {
		panic("list of seat IDs must contain at least one element")
	}
	sort.Ints(seatIDs)
	for i := 1; i < len(seatIDs); i++ {
		if seatIDs[i] != seatIDs[i-1]+1 {
			return seatIDs[i-1] + 1
		}
	}
	return -1
}

func Test_getSeatIDs(t *testing.T) {
	tcs := []struct {
		in          io.Reader
		expectedLen int
		expectedMax int
	}{
		{bytes.NewBufferString(exampleInput), 3, 820},
		{bytes.NewBufferString(input), 868, 938},
	}
	for _, tc := range tcs {
		seatIDs := getSeatIDs(tc.in)
		require.Equal(t, tc.expectedLen, len(seatIDs))
		sort.Ints(seatIDs)
		require.Equal(t, tc.expectedMax, seatIDs[len(seatIDs)-1])
	}
}

func getSeatIDs(r io.Reader) []int {
	seatIDs := make([]int, 0, 50)
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		row := sc.Text()
		_, _, seatID := parseSeating(row)
		seatIDs = append(seatIDs, seatID)
	}
	return seatIDs
}

func Test_parseSeating(t *testing.T) {
	type testCase struct {
		input  string
		row    int
		col    int
		seatID int
	}
	tcs := []testCase{
		{"FBFBBFFRLR", 44, 5, 357},
		{"BFFFBBFRRR", 70, 7, 567},
		{"FFFBBBFRRR", 14, 7, 119},
		{"BBFFBBFRLL", 102, 4, 820},
	}
	for _, tc := range tcs {
		t.Run(tc.input, func(t *testing.T) {
			row, col, seatID := parseSeating(tc.input)
			require.Equal(t, tc.row, row)
			require.Equal(t, tc.col, col)
			require.Equal(t, tc.seatID, seatID)
		})
	}
}

func parseSeating(s string) (row, col, seatID int) {
	if s == "" {
		panic("parseSeating received an empty string")
	}
	for i, ch := range s[:7] {
		if ch == 'B' {
			row |= 1 << (6 - i)
		}
	}
	for i, ch := range s[7:] {
		if ch == 'R' {
			col |= 1 << (2 - i)
		}
	}
	return row, col, row*8 + col
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func mustReadAll(filepath string) string {
	f, err := os.Open(filepath)
	check(err)
	res, err := ioutil.ReadAll(f)
	check(err)
	return string(res)
}
