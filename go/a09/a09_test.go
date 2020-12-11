package a09_test

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
	var i int
	for sc.Scan() {
		// Parse n
		n, err := strconv.Atoi(sc.Text())
		check(err)

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
			fmt.Println("failed to find match for number", n)
			printArr(ns)
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

	require.Equal(t, ns, "")
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
