package main

import (
	"aoc2020/a24"
	"fmt"
	"time"
)

func main() {
	defer func(start time.Time) {
		fmt.Printf("Elapsed time: %v\n", time.Since(start))
	}(time.Now())
	moves := a24.ParseAllMoves("a24/input")
	floor := a24.NewFloor()
	for _, tileMoves := range moves {
		floor.FlipTile(tileMoves)
	}
	for i := 0; i < 100; i++ {
		floor.NextDay()
	}
	fmt.Println(floor.CountBlackTiles())
}
