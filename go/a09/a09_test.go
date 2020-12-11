package a09_test

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func Test_day9(t *testing.T) {
	// 	input := `35
	// 20
	// 15
	// 25
	// 47
	// 40
	// 62
	// 55
	// 65
	// 95
	// 102
	// 117
	// 150
	// 182
	// 127
	// 219
	// 299
	// 277
	// 309
	// 576`
	// 	f := bytes.NewBufferString(input)
	f, err := os.Open("input")
	check(err)
	sc := bufio.NewScanner(f)
	// The plan here is to make a triangle-ish array of arrays
	// Where sums are added as new numbers are added, like this:
	// n1
	//
	// n1 n1+n2
	// n2
	//
	// n1 n1+n2 n1+n3
	// n2 n2+n3
	// n3
	//
	// ...and so on
	// To check if a new number satisfies the requirements, simply
	// iterate over all arrays from index 1:end and check if there is a match
	npreamble := 25
	ns := make([][]int, 0)
	ns2 := make([]int, 0)
	var i int
	offender := 0
	for sc.Scan() {
		// Parse n
		n, err := strconv.Atoi(sc.Text())
		check(err)
		ns2 = append(ns2, n)

		// if we are past the preamble, check if the number has a sum
		if i >= npreamble {
			for j := range ns {
				for _, s := range ns[j][1:] {
					if s == n {
						// remove unused parts of the array
						ns = ns[1:]
						goto EndCheck
					}
				}
			}
			// No match found, exit!
			offender = n
			// printArr(ns)
			break
		}
	EndCheck:

		// add sum of this number and all previous numbers to their arrays
		for j := range ns {
			ns[j] = append(ns[j], ns[j][0]+n)
		}
		ns = append(ns, []int{n})
		i++
	}

	require.Equal(t, 31161678, offender)

	if offender == 0 {
		panic("did not find offender!")
	}
	// Day 2
	cont := findContiguousSet(ns2, offender)
	// sort set, so that min is first and max last
	sort.Ints(cont)
	require.Equal(t, 5453868, cont[0]+cont[len(cont)-1])
}

func printArr(arr [][]int) {
	for _, a := range arr {
		fmt.Print(a[0])
		for _, aa := range a[1:] {
			fmt.Print(",\t", aa)
		}
		fmt.Print("\n")
	}
}

func Test_contiguousSet(t *testing.T) {
	for _, tc := range []struct {
		in   []int
		s    int
		want []int
	}{
		{nil, 3, nil},
		{[]int{}, 3, nil},
		{[]int{1, 2}, 3, []int{1, 2}},
		{[]int{1, 2, 3, 4, 5}, 12, []int{3, 4, 5}},
		{[]int{35, 20, 15, 25, 47, 40, 62, 55}, 127, []int{15, 25, 47, 40}},
	} {
		t.Run(fmt.Sprintf("%v", tc.in), func(t *testing.T) {
			got := findContiguousSet(tc.in, tc.s)
			require.Equal(t, tc.want, got)
		})
	}
}

func findContiguousSet(in []int, s int) []int {
	if len(in) < 2 {
		return nil
	}
	var i, j int
	// starting point
	var cursum = in[0]
	// edge case
	// Re-slice the array based on i, j until j = len(s)-1 and sum < s
	// Then there is no "moves" left that could find the target anymore.
	for {
		// While the current sum is too small, move ahead, adding numbers
		for cursum < s && j < len(in) {
			j++
			cursum += in[j]
		}
		if cursum == s {
			break
		}
		// The sum is now too great, or j == len(in)
		// Subtract from the tail until we are below the sum, or at stop condition
		for cursum > s && i < len(in)-1 {
			cursum -= in[i]
			i++
		}
		if cursum == s {
			break
		}

		// Stop condition
		if i == len(in)-1 {
			return nil
		}
	}
	res := make([]int, j-i+1)
	copy(res, in[i:j+1])
	return res
}
