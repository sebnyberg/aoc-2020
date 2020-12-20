package a18

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_day18(t *testing.T) {
	f, err := os.Open("input")
	check(err)
	sc := bufio.NewScanner(f)
	var sum int
	for sc.Scan() {
		row := sc.Text()
		sum += parse(row).value()
	}
	require.Equal(t, 88534268715686, sum)
}

func Test_parser(t *testing.T) {
	for _, tc := range []struct {
		in   string
		want int
	}{
		{"1 + (2 * 3) + (4 * (5 + 6))", 51},
		{"2 * 3 + (4 * 5)", 46},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 1445},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 669060},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 23340},
	} {
		t.Run(fmt.Sprintf("%v", tc.in), func(t *testing.T) {
			e := parse(tc.in)
			require.Equal(t, tc.want, e.value())
		})
	}
}
