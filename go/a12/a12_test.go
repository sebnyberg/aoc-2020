package a12_test

import (
	"bufio"
	"bytes"
	"log"
	"math"
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

type ship struct {
	x          int
	y          int
	rotDegrees int
}

func Test_move(t *testing.T) {
	rot := 90
	dir := int(math.Sin((float64(rot) / 360) * 2 * math.Pi))
	require.Equal(t, 1, dir)
}

func (p *ship) move(units int) {
	y := int(math.Sin((float64(p.rotDegrees)/360)*2*math.Pi)) * units
	p.y += y
	x := int(math.Cos((float64(p.rotDegrees)/360)*2*math.Pi)) * units
	p.x += x
	// fmt.Println(p.x, p.y, x, y)
}

func Test_day12_part1(t *testing.T) {
	f := bytes.NewBufferString(`F10
N3
F7
R90
F11`)
	// f, err := os.Open("input")
	// check(err)
	sc := bufio.NewScanner(f)
	var s ship
	for sc.Scan() {
		row := sc.Text()
		action := row[0]
		val, err := strconv.Atoi(row[1:])
		check(err)
		switch action {
		case 'N':
			s.y += val
		case 'S':
			s.y -= val
		case 'E':
			s.x += val
		case 'W':
			s.x -= val
		case 'L':
			s.rotDegrees += val
		case 'R':
			s.rotDegrees -= val
		case 'F':
			s.move(val)
		}
	}
	assert.Equal(t, 17, s.x)
	assert.Equal(t, -8, s.y)
}

// func Test_rotateDir(t *testing.T) {
// 	testCases := []struct {
// 		inDir int
// 		inVal int
// 		want  int
// 	}{
// 		{0, -90, 3},
// 		{0, 0, 0},
// 		{3, -90, 2},
// 		{3, 90, 0},
// 		{3, 180, 1},
// 	}
// 	for _, tc := range testCases {
// 		t.Run(fmt.Sprintf("%v/%v", tc.inDir, tc.inVal), func(t *testing.T) {
// 			require.Equal(t, tc.want, rotateDir(tc.inDir, tc.inVal))
// 		})
// 	}
// }

// func rotateDir(dir int, val int) int {
// 	rot := (val / 90)
// 	dir = dir + rot
// 	if dir < 0 {
// 		dir = 4 + dir
// 	}
// 	return dir % 4
// }

func pos(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
