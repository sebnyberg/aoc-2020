package a19_test

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/require"
)

func Test_ruleMatcher(t *testing.T) {
	f, err := os.Open("input")
	check(err)
	sc := bufio.NewScanner(f)
	rows := make([]string, 0)
	for sc.Scan() {
		if sc.Text() == "" {
			break
		}
		rows = append(rows, sc.Text())
	}
	m := parse(rows)

	var ntot int
	for sc.Scan() {
		input := sc.Text()
		if m.match(input) {
			ntot++
		}
	}
	require.Equal(t, 309, ntot)
}

type ruleFn func(s string) (matched bool, width int)

type ruleMatcher struct {
	rules map[int]ruleFn
}

func (m ruleMatcher) match(s string) bool {
	// Custom rule 0... just want to finish this day
	rule8alts := strings.Split(genRecursions("8", "42 | 42 8", len(s)-2), "|")
	rule11alts := strings.Split(genRecursions("11", "42 31 | 42 11 31", len(s)-1), "|")
	rule0alts := make([]string, 0)
	for _, rule8alt := range rule8alts {
		for _, rule11alt := range rule11alts {
			rule0 := strings.Trim(rule8alt, " ") + " " + strings.Trim(rule11alt, " ")
			if len(rule0) > len(s) {
				continue
			}
			rule0alts = append(rule0alts, rule0)
		}
	}
	for _, rule0alt := range rule0alts {
		m.rules[0] = newOrRule(rule0alt, m.rules)
		match, nmatched := m.rules[0](s)
		if match && nmatched >= len(s) {
			return true
		}
	}
	return false
}

func parse(ruleRows []string) ruleMatcher {
	m := ruleMatcher{
		rules: make(map[int]ruleFn),
	}
	for _, row := range ruleRows {
		assignParts := strings.Split(row, ":")
		id, err := strconv.Atoi(assignParts[0])
		check(err)

		// Parse literal
		if strings.ContainsRune(assignParts[1], '"') {
			litParts := strings.Split(assignParts[1], "\"")
			m.rules[id] = func(s string) (bool, int) {
				ch, _ := utf8.DecodeRuneInString(s[:1])
				return strings.ContainsRune(litParts[1], ch), 1
			}
			continue
		}

		// Parse rule referencer (and/or)
		m.rules[id] = newOrRule(assignParts[1], m.rules)
	}
	return m
}

func newOrRule(rule string, otherRules map[int]ruleFn) ruleFn {
	return func(s string) (matched bool, nmatched int) {
		orParts := strings.Split(rule, "|")
		var newmatched int
		for i := 0; i < len(orParts); i++ {
			matched = false
			nmatched = 0
			orPart := orParts[i]

			andParts := strings.Split(strings.Trim(orPart, " "), " ")
			for _, andPart := range andParts {
				newmatched = 0
				ruleID, err := strconv.Atoi(andPart)
				check(err)
				if nmatched >= len(s) {
					return false, nmatched
				}
				matched, newmatched = otherRules[ruleID](s[nmatched:])
				if !matched {
					break
				}
				nmatched += newmatched
			}

			if matched {
				return true, nmatched
			}

		}
		return false, 0
	}
}
func Test_genRecursions(t *testing.T) {
	for _, tc := range []struct {
		id     string
		rule   string
		maxlen int
		want   string
	}{
		{"8", "9 10", 10, "9 10"},
		{"8", "9 10 | 11 | 12", 10, "9 10 | 11 | 12"},
		{"8", "42 | 42 8", 3, "42 | 42 42 | 42 42 42"},
		{"8", "42 | 42 8", 4, "42 | 42 42 | 42 42 42 | 42 42 42 42"},
		{"11", "42 31 | 42 11 31", 3, "42 31"},
		{"11", "42 31 | 42 11 31", 4, "42 31 | 42 42 31 31"},
		{"11", "42 31 | 42 11 31", 6, "42 31 | 42 42 31 31 | 42 42 42 31 31 31"},
	} {
		t.Run(fmt.Sprintf("%v/%v/%v", tc.id, tc.rule, tc.maxlen), func(t *testing.T) {
			got := genRecursions(tc.id, tc.rule, tc.maxlen)
			require.Equal(t, tc.want, got)
		})
	}
}

func genRecursions(ruleID string, rule string, maxlen int) string {
	if !strings.Contains(rule, ruleID) ||
		!strings.ContainsRune(rule, '|') {
		return rule
	}
	orParts := strings.Split(rule, " | ")
	// The terminator ends the recurrence pattern
	terminator := strings.Trim(orParts[0], " ")
	recurser := strings.Trim(orParts[1], " ")
	terminatorLen := len(strings.Split(terminator, " "))
	curLen := len(strings.Split(strings.Trim(orParts[1], " "), " ")) - 1
	if maxlen < curLen+terminatorLen {
		return orParts[0]
	}

	for i := 1; ; i++ {
		// Copy and add last part to the end of the list
		orParts = append(orParts, orParts[i])

		// Replace ruleID with last part within new last part
		orParts[i+1] = strings.Replace(orParts[i], ruleID, recurser, 1)

		// Replace ruleID in current
		orParts[i] = strings.Replace(orParts[i], ruleID, terminator, 1)

		curLen += terminatorLen

		if curLen >= maxlen {
			return strings.Join(orParts[:len(orParts)-1], " | ")
		}
	}
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
