package a13_test

import (
	"log"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// func Test_Day13_part1(t *testing.T) {
// 	buses := make([]int, 0)
// 	for _, n := range strings.Split(input, ",") {
// 		if n == "x" {
// 			continue
// 		}
// 		busID, err := strconv.Atoi(n)
// 		check(err)
// 		buses = append(buses, busID)
// 	}
// 	earliestDepart := 1002462
// 	actualDepart := earliestDepart
// 	earliestBus := 0
// 	for {
// 		for _, bus := range buses {
// 			if actualDepart%bus == 0 {
// 				earliestBus = bus
// 				goto EndLoop
// 			}
// 		}
// 		actualDepart++
// 	}
// EndLoop:
// 	require.Equal(t, 601, earliestBus)
// 	require.Equal(t, 1002468, actualDepart)
// 	require.Equal(t, 3606, (actualDepart-earliestDepart)*earliestBus)
// }

func Test_Day13_part2(t *testing.T) {
	// input := `37,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,41,x,x,x,x,x,x,x,x,x,601,x,x,x,x,x,x,x,x,x,x,x,19,x,x,x,x,17,x,x,x,x,x,23,x,x,x,x,x,29,x,443,x,x,x,x,x,x,x,x,x,x,x,x,13`
	// res := findSubseq(input)
	// require.Equal(t, 1, res)
}

func Test_findSubseq(t *testing.T) {
	for _, tc := range []struct {
		in   string
		want int
	}{
		{"17,x,13,19", 3417},
		{"67,7,59,61", 754018},
		{"67,x,7,59,61", 779210},
		{"67,7,x,59,61", 1261476},
		{"1789,37,47,1889", 1202161486},
	} {
		t.Run(tc.in, func(t *testing.T) {
			require.Equal(t, tc.want, findSubseq(tc.in))
		})
	}
}

func findSubseq(input string) int {
	busOffsets := make(map[int]int)
	maxBus := 0

	offset := 0
	for _, n := range strings.Split(input, ",") {
		if n == "x" {
			offset++
			continue
		}
		i, err := strconv.Atoi(n)
		check(err)
		if i > maxBus {
			maxBus = i
		}
		busOffsets[i] = offset
		offset++
	}

	// find sequence
	for i := busOffsets[maxBus]; ; i += busOffsets[maxBus] {
		for bus, busOffset := range busOffsets {
			if bus == maxBus {
				continue
			}
			if (i+busOffset-busOffsets[maxBus])%bus != 0 {
				goto ContinueLoop
			}
		}

		// Success!
		return i

	ContinueLoop:
	}

	// return 0
}
