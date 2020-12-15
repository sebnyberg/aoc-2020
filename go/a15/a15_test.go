package a15_test

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"
)

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func Test_day15(t *testing.T) {
	// Parse input
	input := "2,0,1,9,5,19"
	// input := "0,3,6"
	nstrs := strings.Split(input, ",")
	ns := make([]int, 0)
	for _, nstr := range nstrs {
		n, err := strconv.Atoi(nstr)
		check(err)
		ns = append(ns, n)
	}

	spokenAtTurn := make(map[int]int)
	for i, n := range ns[:len(ns)-1] {
		spokenAtTurn[n] = i + 1
	}

	// Add -1 to the start to make indexing easier
	lastSpoken := ns[len(ns)-1]

	for turn := len(ns) + 1; turn <= 30000000; turn++ {
		// If the last number was new speak 0 and continue
		lastSpokenAt, exists := spokenAtTurn[lastSpoken]
		if !exists {
			// spoken = append(spoken, 0)
			spokenAtTurn[lastSpoken] = turn - 1
			lastSpoken = 0
			continue
		}

		// Last number has been spoken before
		spokenAtTurn[lastSpoken] = turn - 1
		lastSpoken = turn - 1 - lastSpokenAt
	}
	// // NOTE! Turn
	fmt.Println(lastSpoken)
	t.FailNow()
}
