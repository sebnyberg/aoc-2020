package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// Count number of valid passwords found in the file
func main() {
	fmt.Printf("valid for part 1: %v\n", part1())
	fmt.Printf("valid for part 2: %v\n", part2())
}

func part1() int {
	f, err := os.Open("02/input")
	check(err)
	defer f.Close()

	nvalid := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		rowparts := strings.Split(scanner.Text(), " ")
		rangeparts := strings.Split(rowparts[0], "-")
		min, err := strconv.Atoi(rangeparts[0])
		check(err)
		max, err := strconv.Atoi(rangeparts[1])
		check(err)
		letter := rowparts[1][:1]
		password := rowparts[2]
		occurrences := 0
		for _, ch := range password {
			if string(ch) == letter {
				occurrences++
			}
		}
		if occurrences <= max && occurrences >= min {
			nvalid++
		}
	}
	return nvalid
}

func part2() int {
	f, err := os.Open("02/input")
	check(err)
	defer f.Close()

	nvalid := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		rowparts := strings.Split(scanner.Text(), " ")
		positionparts := strings.Split(rowparts[0], "-")
		firstPos, err := strconv.Atoi(positionparts[0])
		firstPos--
		check(err)
		secondPos, err := strconv.Atoi(positionparts[1])
		secondPos--
		check(err)
		letter := []rune(rowparts[1][:1])[0]
		password := rowparts[2]
		var found bool
		for i, ch := range password {
			if i == firstPos {
				found = ch == letter
				continue
			}
			if i == secondPos {
				if ch == letter && !found || ch != letter && found {
					nvalid++
				}
			}
		}
	}
	return nvalid
}
