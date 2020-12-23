package a23_test

import (
	"log"
	"testing"
)

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func Test_a23(t *testing.T) {
	// input := []int{3, 8, 9, 1, 2, 5, 4, 6, 7}
	// ring := a23.Ring{
	// 	Items: input,
	// }
	// // Pick up three cups
	// pickedUp, err := ring.Remove(1, 3)
	// check(err)

	// // Find destination cup
	// targetLabel := ring.CurrentItem() - 1
	// for {
	// 	if idx := ring.Find(targetLabel); idx != -1 {
	// 		ring.Insert(pickedUp, idx)
	// 	}
	// if targetLabel < lowestValue {
	// 	targetLabel = highestValue
	// }
	// if !ring.HasItem(targetLabel) {
	// 	targetLabel--
	// }
	// }
}
