package a05_test

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

var testInput = `abc

a
b
c

ab
ac

a
a
a
a

b`

func Test_ParseAnswers(t *testing.T) {
	// # characters per line = number of answers
	// # rows per group = number of people

	f, err := os.Open("input")
	check(err)
	// f := bytes.NewBufferString(testInput)

	groupAnswers := parseAnswers(f)
	total := 0
	for _, group := range groupAnswers {
		anyAnswered := make(map[string]struct{})
		personAnswers := make(map[int]map[string]struct{})
		// collect answers into map structure
		for personIdx, answer := range group {
			personAnswers[personIdx] = make(map[string]struct{})
			for _, ch := range answer {
				anyAnswered[string(ch)] = struct{}{}
				personAnswers[personIdx][string(ch)] = struct{}{}
			}
		}
		// fmt.Printf("%+v\n", personAnswers)
		// each question that anyone answered should be answered by all
		for anyAnswer := range anyAnswered {
			answered := true
			for personIdx := range group {
				_, personAnswered := personAnswers[personIdx][anyAnswer]
				if !personAnswered {
					answered = false
				}
			}
			if answered == true {
				total++
			}
		}
	}
	fmt.Println(total)

	t.FailNow()
}

func parseAnswers(r io.Reader) [][]string {
	sc := bufio.NewScanner(r)
	rows := make([][]string, 1)
	rows[0] = make([]string, 0)
	groupIdx := 0

	for sc.Scan() {
		row := strings.Trim(sc.Text(), " ")
		if row == "" {
			groupIdx++
			rows = append(rows, make([]string, 0))
			continue
		}
		rows[groupIdx] = append(rows[groupIdx], row)
	}
	return rows
}
