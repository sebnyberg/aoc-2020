package a12_test

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

const (
	East  int = 0
	South int = 1
	West  int = 2
	North int = 3
)

func Test_day12(t *testing.T) {
	// 	f := bytes.NewBufferString(`F10
	// N3
	// F7
	// R90
	// F11`)
	f, err := os.Open("input")
	check(err)
	sc := bufio.NewScanner(f)
	var x, y int
	dir := 0
	for sc.Scan() {
		row := sc.Text()
		action := row[0]
		val, err := strconv.Atoi(row[1:])
		check(err)
		switch action {
		case 'N':
			y += val
		case 'S':
			y -= val
		case 'E':
			x += val
		case 'W':
			x -= val
		case 'L':
			dir = rotateDir(dir, -val)
		case 'R':
			dir = rotateDir(dir, val)
		case 'F':
			switch dir {
			case East:
				x += val
			case West:
				x -= val
			case North:
				y += val
			case South:
				y -= val
			}
		}
	}
	// assert.Equal(t, 10, pos(x))
	// assert.Equal(t, 10, pos(y))
	require.Equal(t, 100, pos(x)+pos(y))
}

func Test_rotateDir(t *testing.T) {
	testCases := []struct {
		inDir int
		inVal int
		want  int
	}{
		{0, -90, 3},
		{0, 0, 0},
		{3, -90, 2},
		{3, 90, 0},
		{3, 180, 1},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v/%v", tc.inDir, tc.inVal), func(t *testing.T) {
			require.Equal(t, tc.want, rotateDir(tc.inDir, tc.inVal))
		})
	}
}

func rotateDir(dir int, val int) int {
	rot := (val / 90)
	dir = dir + rot
	if dir < 0 {
		dir = 4 + dir
	}
	return dir % 4
}

func pos(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
