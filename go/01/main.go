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
			for j := i + 1; j < len(ns)-1; j++ {
				n1, n2 := ns[i], ns[j]
				if n1+n2+cur == 2020 {
					fmt.Printf("%d * %d * %d = %d\n", ns[i], ns[j], cur, ns[i]*ns[j]*cur)
				}
			}
		}
	}
	check(scanner.Err())
}
