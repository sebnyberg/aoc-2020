package a23_test

import (
	"aoc2020/a23"
	"fmt"
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
	min, max := findMinMax(input)

	ring := &a23.Ring{
		Items: input,
	}

	doPrint = false
	nmoves := 100
	for i := 1; i <= nmoves; i++ {
		println("-- move", i, "--")
		makeMove(ring, min, max)
		println()
	}

	fmt.Println("-- final --")
	fmt.Printf("cups: %v\n", ring)

	fmt.Println("shifting so that 1 is first")
	ring.ShiftRight(ring.Find(1))
	fmt.Printf("cups: %v\n", ring)

	resultInts, err := ring.Remove(1, len(ring.Items)-1)
	check(err)
	var resultsb strings.Builder
	for _, n := range resultInts {
		resultsb.WriteString(strconv.Itoa(n))
	}
	require.Equal(t, "62934785", resultsb.String())
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

func makeMove(ring *a23.Ring, min, max int) {
	printf("cups: %v\n", ring)

	// Pick up three cups
	pickedUp, err := ring.Remove(1, 3)
	printf("pick up: %+v\n", pickedUp)
	check(err)
	printf("after pick up: %v\n", ring)

	// Find destination cup
	targetLabel := ring.CurrentItem() - 1
	for {
		// If a destination cup was found, insert picked up cups in that location
		if idx := ring.Find(targetLabel); idx != -1 {
			println("destination:", idx+1)
			ring.Insert(pickedUp, idx+1)
			break
		}

		// If target label is smaller than the label of any cup,
		// reset to the highest value
		if targetLabel < min {
			targetLabel = max
			continue
		}

		// Reduce target and try again
		targetLabel--
	}

	// Shift to next cup
	ring.ShiftRight(1)
}
