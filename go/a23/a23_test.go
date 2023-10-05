package a23_test

import (
	"aoc2020/a23"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func findMinMax(ns []int) (min, max int) {
	min = int(1e9)
	for _, n := range ns {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return
}

func Test_a23(t *testing.T) {
	// input := []int{3, 8, 9, 1, 2, 5, 4, 6, 7}
	input := []int{1, 9, 8, 7, 5, 3, 4, 6, 2}

	ring := &a23.LinkedRing{
		ItemPos: make(map[int]*a23.LinkedItem, 1e6),
	}
	// cba to fix this...
	ring.Insert(input[:1])
	ring.Insert(input[1:])

	day2 := true
	var ntotal int = 1e6
	if day2 {
		for i := 10; i <= ntotal; i++ {
			ring.InsertBefore(i)
		}
	}

	require.Equal(t, int(1e6), ring.Len)
	require.Equal(t, ring.First.Val, input[0])
	require.Equal(t, ring.First.Prev.Val, int(1e6))

	doPrint = false
	min, max := 1, ring.Len
	var nmoves int = 1e7
	for i := 1; i <= nmoves; i++ {
		// println("-- move", i, "--")
		makeMove(ring, min, max)
		// println()
	}

	// fmt.Println("-- final --")
	// fmt.Printf("cups: %v\n", ring)

	ring.ShiftTo(1)
	require.Equal(t, 693659135400, ring.First.Next.Val*ring.First.Next.Next.Val)
}

var doPrint = true

func println(msgs ...interface{}) {
	if doPrint {
		fmt.Println(msgs...)
	}
}

func print(msgs ...interface{}) {
	if doPrint {
		fmt.Print(msgs...)
	}
}

func printf(format string, args ...interface{}) {
	if doPrint {
		fmt.Printf(format, args...)
	}
}

func makeMove(ring *a23.LinkedRing, min, max int) {
	// printf("cups: %v\n", ring)

	startVal := ring.First.Val

	// Pick up three cups
	pickedUp := ring.Remove(3)
	// printf("pick up: %+v\n", pickedUp)
	// printf("after pick up: %v\n", ring)

	// Find destination cup
	targetLabel := ring.First.Val - 1
	for {
		if ring.ShiftTo(targetLabel) {
			// printf("on insert position: %+v\n", ring)
			ring.Insert(pickedUp)
			// printf("after insert: %+v\n", ring)
			break
		}

		// If target label is smaller than the label of any cup,
		// reset to the highest value
		if targetLabel < min {
			// println("target label below min value, setting to", max)
			targetLabel = max
			continue
		}

		// Reduce target and try again
		targetLabel--
	}

	// Shift to first position
	ring.ShiftTo(startVal)
	// printf("after reset: %+v\n", ring)
	// Shift to position to the right of the first position
	ring.ShiftRight(1)
	// printf("one step to the right of reset: %+v\n", ring)
}
