package a19_test

import (
	"bufio"
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
	require.Equal(t, 0, ntot)
	t.FailNow()
}

type ruleFn func(s string) (matched bool, width int)

type ruleMatcher struct {
	rules map[int]ruleFn
}

func (m ruleMatcher) match(s string) bool {
	matched, nmatched := m.rules[0](s)
	if nmatched < len(s) {
		return false
	}
	return matched
}

func parse(ruleRows []string) ruleMatcher {
	m := ruleMatcher{
		rules: make(map[int]ruleFn),
	}
	for id, row := range ruleRows {
		assignParts := strings.Split(row, ":")

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
		m.rules[id] = func(s string) (matched bool, nmatched int) {
			orParts := strings.Split(assignParts[1], "|")
			var newmatched int
			for _, orPart := range orParts {
				matched = false
				nmatched = 0

				andParts := strings.Split(strings.Trim(orPart, " "), " ")
				for _, andPart := range andParts {
					newmatched = 0
					ruleID, err := strconv.Atoi(andPart)
					check(err)
					if nmatched >= len(s) {
						return false, nmatched
					}
					matched, newmatched = m.rules[ruleID](s[nmatched:])
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
	return m
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
