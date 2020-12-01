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
		for i := 0; i < len(ns)-1; i++ {
			if ns[i]+cur == 2020 {
				fmt.Printf("%d * %d = %d\n", ns[i], cur, ns[i]*cur)
			}
		}
	}
	check(scanner.Err())
}
