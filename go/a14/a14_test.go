package a14_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_parseInput(t *testing.T) {
	testCases := []struct {
		input string
		want  string
	}{
		{"mem[4616] = 8311689", "mem"},
		{"mem[4616] = 8311689", "mem"},
		{"mem[4616] = 8311689", "mem"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.input, func(t *testing.T) {
			got := parseInput(testCase.input)
			if got != testCase.want {
				require.Equal(t, testCase.want, got)
			}
		})
	}
}

func parseInput(input string) string {
	return "mem"
}
