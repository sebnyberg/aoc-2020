package a16_test

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type RangeRule struct {
	start int
	end   int
}

func Test_day16(t *testing.T) {
	f, err := os.Open("input")
	check(err)
	sc := bufio.NewScanner(f)
	validations := make(map[string][]RangeRule)

	// scan rules
	for sc.Scan() {
		row := sc.Text()
		if row == "" {
			break
		}
		ruleparts := strings.Split(row, ":")
		ruleName := ruleparts[0]
		ruleranges := strings.Split(ruleparts[1], "or")
		for _, rulerange := range ruleranges {
			rangeparts := strings.Split(strings.Trim(rulerange, " "), "-")
			start, err := strconv.Atoi(rangeparts[0])
			check(err)
			end, err := strconv.Atoi(rangeparts[1])
			check(err)
			if _, exists := validations[ruleName]; exists {
				validations[ruleName] = append(validations[ruleName], RangeRule{start, end})
				continue
			}
			validations[ruleName] = []RangeRule{{start, end}}
		}
	}

	// Scan your ticket
	// var ticket []int
	sc.Scan() // Skip header
	for sc.Scan() {
		row := sc.Text()
		if row == "" {
			break
		}
		// ticket = parseTicket(row)
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

	// Validate nearby tickets
	invalidFields := make([]int, 0)
	validTickets := make([][]int, 0)
	for _, nearbyTicket := range nearbyTickets {
		invalidFieldsForTicket := validateTicket(nearbyTicket, validations)
		if len(invalidFieldsForTicket) == 0 {
			validTickets = append(validTickets, nearbyTicket)
			continue
		}
		invalidFields = append(invalidFields, invalidFieldsForTicket...)
	}

	// Part 1
	var sum int
	for _, invalidField := range invalidFields {
		sum += invalidField
	}
	require.Equal(t, 20058, sum)

	// Part 2
	ruleNameValidForField := make(map[string]int)

	// For each valid ticket and field, check which rule applies
	for ticketEntryIdx := 0; ticketEntryIdx < len(validTickets[0]); ticketEntryIdx++ {
		// For each rule
		for fieldName, fieldRules := range validations {
			validForAll := true
			// For each ticket
			for _, validTicket := range validTickets {
				// If none of the range rules apply, continue to next field
				validForTicket := false
				for _, rangeRules := range fieldRules {
					if validTicket[ticketEntryIdx] >= rangeRules.start && validTicket[ticketEntryIdx] <= rangeRules.end {
						validForTicket = true
						break
					}
				}
				if !validForTicket {
					validForAll = false
					break
				}
			}
			if validForAll {
				ruleNameValidForField[fieldName] = ticketEntryIdx
			}
		}
	}

	fmt.Println(ruleNameValidForField)
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

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
