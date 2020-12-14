package a10_test

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"sort"
// 	"strconv"
// 	"testing"
// )

// func check(err error) {
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// }

// func Test_Day10(t *testing.T) {
// 	f, err := os.Open("testinput4")
// 	check(err)
// 	sc := bufio.NewScanner(f)
// 	adapters := make([]int, 0)
// 	for sc.Scan() {
// 		n, err := strconv.Atoi(sc.Text())
// 		check(err)
// 		adapters = append(adapters, n)
// 	}
// 	// Append initial adapter
// 	adapters = append(adapters, 0)
// 	sort.Ints(adapters)

// 	printArr(adapters)

// 	arrs := 1
// 	// Count arrangements for each section
// 	j := 0
// 	for i := 1; i < len(adapters); i++ {
// 		if adapters[i]-adapters[i-1] == 3 {
// 			printArr(adapters[j:i])
// 			sectionArrs := getArrs(adapters[j:i])
// 			fmt.Println("section arrs: ", sectionArrs)
// 			arrs *= sectionArrs
// 			j = i - 1
// 		}
// 	}

// 	fmt.Println(arrs)

// 	t.FailNow()
// }

// var depth = 0
// var iter = 0

// func getArrs(adapters []int) int {
// 	arrs := 1

// 	// fmt.Println("depth", depth)
// 	depth++
// 	iter++
// 	if depth > 100 || iter%1000 == 0 {
// 		fmt.Println(depth, iter)
// 	}

// 	// Plan: when faced with a removal option,
// 	// run and return getArrs with the item removed
// 	for i := 2; i < len(adapters); i++ {
// 		if adapters[i]-adapters[i-2] <= 3 {
// 			// printArr(adapters)
// 			// fmt.Printf("removing item %v at %v\n", adapters[i-1], i)
// 			// Try to remove the middle adapter and run getArrs again
// 			withoutElem := make([]int, len(adapters))
// 			copy(withoutElem, adapters)
// 			withoutElem = append(withoutElem[i-2:i-1], withoutElem[i:]...)
// 			arrs += getArrs(withoutElem)
// 		}
// 	}

// 	depth--

// 	return arrs
// }

// func printArr(adapters []int) {
// 	fmt.Print("(", adapters[0], "), ")
// 	for _, a := range adapters[1:] {
// 		fmt.Print(a, ", ")
// 	}
// 	fmt.Print("(", adapters[len(adapters)-1]+3, ")")
// 	fmt.Printf("\n")
// }
