package a16_test

import (
	"bufio"
	"errors"
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
	myTicket := make([]int, 0)
	for sc.Scan() {
		row := sc.Text()
		if row == "" {
			break
		}
		myTicket = parseTicket(row)
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

	// Reorder tickets by field
	ticketByField := make([][]int, len(validTickets[0]))
	for _, validTicket := range validTickets {
		for i := 0; i < len(validTickets[0]); i++ {
			ticketByField[i] = append(ticketByField[i], validTicket[i])
		}
	}
	require.Equal(t, validTickets[0][0], ticketByField[0][0])
	require.Equal(t, validTickets[0][1], ticketByField[1][0])
	require.Equal(t, validTickets[len(validTickets)-1][0], ticketByField[0][len(validTickets)-1])

	fieldRules := make(map[int]map[string]bool)
	for colIdx, ticketCol := range ticketByField {
		fieldRules[colIdx] = make(map[string]bool)
		for rulename, rules := range validations {
			if fieldsValid(rules, ticketCol) {
				fieldRules[colIdx][rulename] = true
			}
		}
	}

	logicisize(fieldRules)
	sum = 1
	for fieldIdx, ruleNames := range fieldRules {
		for k := range ruleNames {
			if strings.HasPrefix(k, "departure") {
				sum *= myTicket[fieldIdx]
			}
		}
	}

	require.Equal(t, 366871907221, sum)
}

func Test_logicisize(t *testing.T) {
	for _, tc := range []struct {
		in   map[int]map[string]bool
		want map[int]map[string]bool
	}{
		{
			map[int]map[string]bool{
				0: {"a": true, "c": true},
				1: {"a": true, "b": true, "c": true},
				2: {"c": true},
			},
			map[int]map[string]bool{
				0: {"a": true},
				1: {"b": true},
				2: {"c": true},
			},
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.in), func(t *testing.T) {
			logicisize(tc.in)
			require.Equal(t, nil, logicisize(tc.in))
			require.Equal(t, tc.want, tc.in)
		})
	}
}

func logicisize(in map[int]map[string]bool) (err error) {
	check := func(in map[int]map[string]bool) bool {
		for _, v := range in {
			if len(v) != 1 {
				return false
			}
		}
		return true
	}

	for !check(in) {
		singleItemIndex := -1
		singleItem := ""
		for k, v := range in {
			if len(v) == 1 {
				singleItemIndex = k
				for vv := range in[k] {
					singleItem = vv
				}
				break
			}
		}

		if singleItemIndex == -1 {
			return errors.New("needs at least one row with one item to reduce")
		}

		for k := range in {
			if k == singleItemIndex {
				continue
			}
			delete(in[k], singleItem)
		}
	}
	return nil
}

func fieldsValid(rules []RangeRule, fieldValues []int) bool {
	// if no rule matches, return false
	for _, fieldValue := range fieldValues {
		ok := false
		for _, rule := range rules {
			// rule matches, continue outer loop
			if fieldValue >= rule.start && fieldValue <= rule.end {
				ok = true
			}
		}
		if !ok {
			return false
		}
	}
	return true
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
