package a16_test

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

type RangeRule struct {
	start int
	end   int
}

func Test_day16(t *testing.T) {
	f, err := os.Open("testinput")
	check(err)
	sc := bufio.NewScanner(f)
	rules := make(map[string][]RangeRule)

	// scan rules
	for sc.Scan() {
		row := sc.Text()
		if row == "" {
			break
		}
		ruleparts := strings.Split(row, ":")
		rulename := ruleparts[0]
		ruleranges := strings.Split(ruleparts[1], "or")
		for _, rulerange := range ruleranges {
			rangeparts := strings.Split(strings.Trim(rulerange, " "), "-")
			start, err := strconv.Atoi(rangeparts[0])
			check(err)
			end, err := strconv.Atoi(rangeparts[1])
			check(err)
			if _, exists := rules[rulename]; exists {
				rules[rulename] = append(rules[rulename], RangeRule{start, end})
				continue
			}
			rules[rulename] = []RangeRule{{start, end}}
		}
	}

	// Scan your ticket
	var ticket []int
	sc.Scan() // Skip header
	for sc.Scan() {
		row := sc.Text()
		if row == "" {
			break
		}
		ticket = parseTicket(row)
	}

	// Scan nearby tickets
	nearbyTickets := make([][]int, 0)
	sc.Scan() // Skip header
	for sc.Scan() {
		row := sc.Text()
		if row == "" {
			break
		}
		nearbyTickets = append(nearbyTickets, parseTicket(row))
	}

	fmt.Println(ticket)
	fmt.Println(nearbyTickets)
	fmt.Println(rules)

	// Validate nearby tickets
	invalidFields := make([]int, 0)
	for _, nearbyTicket := range nearbyTickets {
		invalidFields = append(invalidFields, validateTicket(nearbyTicket, rules)...)
	}

	fmt.Println(invalidFields)

	t.FailNow()
}

func validateTicket(ticket []int, validations map[string][]RangeRule) []int {
	invalidFields := make([]int, 0)
	for _, ticketEntry := range ticket {
		isvalid := false
		for _, validation := range validations {
			for _, rule := range validation {
				if ticketEntry >= rule.start && ticketEntry <= rule.end {
					isvalid = true
					break
				}
			}
		}
		if !isvalid {
			invalidFields = append(invalidFields, ticketEntry)
		}
	}
	return invalidFields
}

func parseTicket(row string) []int {
	ticket := make([]int, 0)
	rownums := strings.Split(row, ",")
	for _, rownumstr := range rownums {
		n, err := strconv.Atoi(rownumstr)
		check(err)
		ticket = append(ticket, n)
	}
	return ticket
}
