package a10_test

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

func Test_day10(t *testing.T) {
	f, err := os.Open("input")
	check(err)
	sc := bufio.NewScanner(f)
	ns := make([]int, 0, 50)
	ns = append(ns, 0)
	for sc.Scan() {
		n, err := strconv.Atoi(sc.Text())
		check(err)
		ns = append(ns, n)
	}
	sort.Ints(ns)
	segs := findSegments(ns)
	narr := 1
	for _, seg := range segs {
		fmt.Println(len(seg))
		printArr(seg)
		switch len(seg) {
		case 3:
			narr *= 2
		case 4:
			narr *= 4
		case 5:
			narr *= 7
		}
	}

	require.Equal(t, 84627647627264, narr)
}

func Test_findSegments(t *testing.T) {
	for _, tc := range []struct {
		in   []int
		want [][]int
	}{
		{[]int{1, 2, 3, 4}, [][]int{{1, 2, 3, 4}}},
		{[]int{1, 2, 3, 4, 7, 8}, [][]int{{1, 2, 3, 4}, {7, 8}}},
		{[]int{1, 4, 7, 8, 11, 14, 17}, [][]int{{1}, {4}, {7, 8}, {11}, {14}, {17}}},
	} {
		t.Run(fmt.Sprintf("%v", tc.in), func(t *testing.T) {
			require.Equal(t, tc.want, findSegments(tc.in))
		})
	}
}

func findSegments(in []int) [][]int {
	segment := 0
	res := [][]int{
		{in[0]},
	}
	for i := 1; i < len(in); i++ {
		if in[i]-in[i-1] == 2 {
			panic("wut a 2")
		}
		if in[i]-in[i-1] == 3 {
			segment++
			res = append(res, []int{in[i]})
			continue
		}
		res[segment] = append(res[segment], in[i])
	}
	return res
}

func printArr(in []int) {
	fmt.Print(in[0])
	for _, i := range in[1:] {
		fmt.Printf(", %v", i)
	}
	fmt.Print("\n")
}
