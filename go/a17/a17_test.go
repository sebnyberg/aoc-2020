package a17_test

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// Point in 3D space
type P struct {
	x int
	y int
	z int
	w int
}

func Test_Nearby(t *testing.T) {
	nearby := P{1, 1, 1, 1}.Nearby(1)
	require.Equal(t, 80, len(nearby))
}

func (p P) Nearby(dist int) []P {
	nearby := make([]P, 0, int(math.Pow(float64(1+2*dist), 3))-1)
	for w := p.w - dist; w <= p.w+dist; w++ {
		for z := p.z - dist; z <= p.z+dist; z++ {
			for y := p.y - dist; y <= p.y+dist; y++ {
				for x := p.x - dist; x <= p.x+dist; x++ {
					if w == p.w && z == p.z && y == p.y && x == p.x {
						continue
					}
					nearby = append(nearby, P{x, y, z, w})
				}
			}
		}
	}
	return nearby
}

func (p P) String() string {
	return fmt.Sprintf("(%v,%v,%v)", p.x, p.y, p.z)
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func Test_day17(t *testing.T) {
	grid := parseGrid("input")

	printGrid(grid)
	for i := 0; i < 6; i++ {
		grid = cycleActive(grid)
	}
	nactive := 0
	for _, active := range grid {
		if active {
			nactive++
		}
	}
	require.Equal(t, 1980, nactive)
}

func cycleActive(grid map[P]bool) map[P]bool {
	// The original grid does not include points which need to be considered,
	// We need to pad the grid first
	paddedGrid := make(map[P]bool)

	// Create padded result grid
	for p, active := range grid {
		paddedGrid[p] = active
		// Get nearby cubes for each point
		nearby := p.Nearby(1)
		for _, p := range nearby {
			paddedGrid[p] = grid[p]
		}
	}

	// Determine if each point should be active or not
	nextGrid := make(map[P]bool)
	for p, active := range paddedGrid {
		nactive := 0
		for _, nearbyP := range p.Nearby(1) {
			if paddedGrid[nearbyP] {
				nactive++
			}
		}
		// if nactive >= 2 {
		// 	fmt.Println(p)
		// }
		nextGrid[p] = active
		if active {
			remainsActive := nactive == 2 || nactive == 3
			// fmt.Printf("active %v,\tremainsActive: %v\n", p, remainsActive)
			if !remainsActive {
				delete(nextGrid, p)
			} else {
				nextGrid[p] = remainsActive
			}
		} else {
			becomesActive := nactive == 3
			// fmt.Printf("inactive %v,\tbecomesActive: %v\n", p, becomesActive)
			if !becomesActive {
				delete(nextGrid, p)
			} else {
				nextGrid[p] = true
			}
			// nextGrid[p] = becomesActive
		}
	}
	return nextGrid
}

func parseGrid(fpath string) map[P]bool {
	f, err := os.Open(fpath)
	check(err)
	sc := bufio.NewScanner(f)
	var x, y, z, w int
	grid := make(map[P]bool)
	for sc.Scan() {
		row := sc.Text()
		for _, ch := range row {
			switch ch {
			case '#':
				grid[P{x, y, z, w}] = true
			case '.':
				grid[P{x, y, z, w}] = false
			default:
				panic(ch)
			}
			x++
		}
		x = 0
		y++
	}
	return grid
}

const minInt = -int(^uint(0)>>1) - 1
const maxInt = int(^uint(0) >> 1)

func printGrid(grid map[P]bool) {
	// Put all points in a grid
	// Collect max x,y,z for padding later on
	minx, miny, minz, minw := maxInt, maxInt, maxInt, maxInt
	maxx, maxy, maxz, maxw := minInt, minInt, minInt, minInt
	for p := range grid {
		if minx > p.x {
			minx = p.x
		}
		if maxx < p.x {
			maxx = p.x
		}
		if miny > p.y {
			miny = p.y
		}
		if maxy < p.y {
			maxy = p.y
		}
		if minz > p.z {
			minz = p.z
		}
		if maxz < p.z {
			maxz = p.z
		}
		if minw > p.w {
			minw = p.w
		}
		if maxw < p.w {
			maxw = p.w
		}
	}

	fmt.Print("\n")
	for w := minw; w <= maxw; w++ {
		for z := minz; z <= maxz; z++ {
			fmt.Printf("z=%v, w=%v\n", z, w)
			for y := miny; y <= maxy; y++ {
				for x := minx; x <= maxx; x++ {
					if grid[P{x, y, z, 0}] {
						fmt.Print("#")
					} else {
						fmt.Print(".")
					}
				}
				fmt.Print("\n")
			}
			fmt.Print("\n")
		}
	}
}
