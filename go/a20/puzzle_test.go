package a20_test

import (
	"aoc2020/a20"
	"os"
	"testing"
)

func Test_Puzzle_Solve(t *testing.T) {
	f, err := os.Open("input")
	check(err)
	puzzle := a20.ParsePuzzle(f)
	puzzle.Solve()
	t.FailNow()
}
