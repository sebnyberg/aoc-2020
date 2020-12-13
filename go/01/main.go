package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	f, err := os.Open("01/input")
	check(err)
	defer f.Close()
	ns := make([]int, 0, 1000)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		cur, err := strconv.Atoi(scanner.Text())
		check(err)
		ns = append(ns, cur)
		for _, n1 := range ns {
			for _, n2 := range ns {
				if n1+n2+cur == 2020 {
					fmt.Printf("%d * %d * %d = %d\n", n1, n2, cur, n1*n2*cur)
				}
			}
		}
	}
	check(scanner.Err())
}
