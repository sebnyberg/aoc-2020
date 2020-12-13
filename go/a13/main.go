package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	fmt.Println(findSubseq("37,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,41,x,x,x,x,x,x,x,x,x,601,x,x,x,x,x,x,x,x,x,x,x,19,x,x,x,x,17,x,x,x,x,x,23,x,x,x,x,x,29,x,443,x,x,x,x,x,x,x,x,x,x,x,x,13"))
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(ns ...int) int {
	if len(ns) == 1 {
		return ns[0]
	}
	result := ns[0] * ns[1] / gcd(ns[0], ns[1])

	for _, i := range ns[2:] {
		result = result * i / gcd(result, i)
	}

	return result
}

func findSubseq(input string) int {
	buses := make([]int, 0)
	offsets := make([]int, 0)

	// Parse bus IDs and offsets
	for i, nstr := range strings.Split(input, ",") {
		if nstr == "x" {
			continue
		}
		n, err := strconv.Atoi(nstr)
		check(err)
		buses = append(buses, n)
		offsets = append(offsets, i)
	}

	// Increment with the first bus ID until we have a number that
	// satisfies both the first and second bus ID.
	// Then change the increment to be the lowest common multiplier
	// between the first two buses and continue. Once third bus is matched,
	// change ingrement to lcm(b1,b2,b3) and so on, until done.
	incr := buses[0]
	i := 1
	t := 0
	for i < len(buses) {
		t += incr
		if (t+offsets[i])%buses[i] == 0 {
			incr = lcm(buses[0 : i+1]...)
			fmt.Printf("Bus %v joins in at an offset of %v at t=%v\n", i, offsets[i], t)
			fmt.Printf("Increment adjusted to %v\n", incr)
			i++
		}
	}

	return t
}
