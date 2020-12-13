package a12_test

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"gotest.tools/assert"
)

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func Test_day12(t *testing.T) {
	f, err := os.Open("input")
	check(err)
	sc := bufio.NewScanner(f)
	var x, y int

	printPos := func(markerX, markerY, x, y int) {
		fmt.Printf("Pos = (%03d, %03d) M = (%03d, %03d)\n", x, y, markerX, markerY)
	}
	markerX := 10
	markerY := 1
	for sc.Scan() {
		row := sc.Text()
		action := row[0]
		val, err := strconv.Atoi(row[1:])
		check(err)
		switch action {
		case 'N':
			markerY += val
		case 'S':
			markerY -= val
		case 'E':
			markerX += val
		case 'W':
			markerX -= val
		case 'L':
			markerX, markerY = rotateMarker(markerX, markerY, val)
		case 'R':
			markerX, markerY = rotateMarker(markerX, markerY, -val)
		case 'F':
			x += val * markerX
			y += val * markerY
		}
		printPos(markerX, markerY, x, y)
	}
	require.Equal(t, 28885, abs(x)+abs(y))
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Test_rotateMarker(t *testing.T) {
	for _, tc := range []struct {
		inX     int
		inY     int
		degrees int
		wantX   int
		wantY   int
	}{
		{10, 1, -90, 1, -10},
		{10, 1, 0, 10, 1},
		{10, 1, 90, -1, 10},
		{10, 1, 180, -10, -1},
		{10, 1, 270, 1, -10},
		{10, 1, 360, 10, 1},
		{10, 1, 450, -1, 10},
	} {
		t.Run(fmt.Sprintf("(%v,%v)+%v", tc.inX, tc.inY, tc.degrees), func(t *testing.T) {
			gotX, gotY := rotateMarker(tc.inX, tc.inY, tc.degrees)
			assert.Equal(t, tc.wantX, gotX)
			assert.Equal(t, tc.wantY, gotY)
		})
	}
}

func rotateMarker(curX, curY int, degrees int) (newX, newY int) {
	if degrees < 0 {
		degrees = 360 + (degrees % 360)
	}
	newX, newY = curX, curY
	switch degrees % 360 {
	case 90:
		newX = -curY
		newY = curX
	case 180:
		newX = -curX
		newY = -curY
	case 270:
		newX = curY
		newY = -curX
	}
	return newX, newY
}
