package a10_test

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"testing"
)

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func Test_day10(t *testing.T) {
	var err error
	var f io.Reader
	// Read inputs, each row is an adapter with an output joltage
	// Each adapter can take an input outputJoltage-3 <= jolt <= outputJoltage
	f, err = os.Open("input")
	check(err)

	sc := bufio.NewScanner(f)
	ns := make([]int, 0)
	var n int
	for sc.Scan() {
		n, err = strconv.Atoi(sc.Text())
		check(err)
		ns = append(ns, n)
	}
	sort.Ints(ns)

	prevN := ns[0]
	n1jolt := 0
	n3jolt := 1

	if ns[0] == 1 {
		n1jolt++
	}
	if ns[0] == 3 {
		n3jolt++
	}
	for _, n := range ns[1:] {
		if n == prevN {
			panic("prevN == n, should not happen")
		}
		// End condition
		if n-prevN > 3 {
			break
		}

		if n-prevN == 1 {
			n1jolt++
		}

		if n-prevN == 3 {
			n3jolt++
		}
		prevN = n
	}

	fmt.Println(n1jolt)
	fmt.Println(n3jolt)
	fmt.Println(n1jolt * n3jolt)
	t.FailNow()
}
