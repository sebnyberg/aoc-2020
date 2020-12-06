package main

import (
	"aoc2020/a04"
	"fmt"
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	runDay := func(inputPath string, handler func(f *os.File)) {
		f, err := os.Open(inputPath)
		check(err)
		defer f.Close()
		handler(f)
	}

	if len(os.Args) <= 1 {
		log.Fatalln("provide the AOC day as an argument to get its solution, e.g. 'go run main.go 04'")
	}

	switch os.Args[1] {
	case "04":
		runDay("a04/input", func(f *os.File) {
			fmt.Printf("valid passports: %v\n", a04.Run(f))
		})
	}
}
