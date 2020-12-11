package a07_test

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
	"unicode"

	"github.com/stretchr/testify/require"
)

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func Test_countBagsInShinyGold(t *testing.T) {
	f, err := os.Open("input")
	check(err)

	sc := bufio.NewScanner(f)

	// Read list of bag -> contained bags from input file
	bagsList := make([]ColoredBag, 0)
	for sc.Scan() {
		bag, err := ParseBag(sc.Text())
		check(err)
		bagsList = append(bagsList, bag)
	}

	// Create map of bag id -> bags inside + count
	bags := make(map[string]map[string]int)
	for _, bag := range bagsList {
		bags[bag.id] = bag.bagsInside
	}

	// Bags to add
	require.Equal(t, 34988, countBagsInside(bags, "shiny gold"))
}

func countBagsInside(bagsMap map[string]map[string]int, bagID string) int {
	var sum int
	for insideBagID, insideBagCount := range bagsMap[bagID] {
		sum += insideBagCount
		sum += insideBagCount * countBagsInside(bagsMap, insideBagID)
	}
	return sum
}

func Test_findShinyGold(t *testing.T) {
	f, err := os.Open("input")
	check(err)

	sc := bufio.NewScanner(f)

	// Read list of bag -> contained bags from input file
	bags := make([]ColoredBag, 0)
	for sc.Scan() {
		bag, err := ParseBag(sc.Text())
		check(err)
		bags = append(bags, bag)
	}

	// List of BagIDs which contain the sought-after bag
	require.Equal(t, 177, len(findShinyGold(bags)))
}

func findShinyGold(rawBags []ColoredBag) map[string]struct{} {
	// Read list of bag -> contained bags from input file
	bags := map[string]map[string]struct{}{}
	for _, bag := range rawBags {
		if _, exists := bags[bag.id]; exists {
			log.Fatalln("same bag parsed twice... possible error")
		}
		bags[bag.id] = make(map[string]struct{})
		for containedBagID := range bag.bagsInside {
			bags[bag.id][containedBagID] = struct{}{}
		}
	}

	containsBag := map[string]struct{}{}

	// How many bag colors can eventually contain at least one shiny gold bag?
	// Going with brute-force
	for bagID := range bags {
		// Keep track of which bags we've seen so far to avoid circular dependencies
		seen := make(map[string]struct{})
		seen[bagID] = struct{}{}

		check := make(map[string]struct{})
		for k := range bags[bagID] {
			check[k] = struct{}{}
		}

		var cur string
		for len(check) > 0 {
			// Pick first bag
			for cur = range check {
				break
			}

			// Check if current is the sought-after bag
			if cur == "shiny gold" {
				containsBag[bagID] = struct{}{}
				break
			}

			// Mark bag as seen
			seen[cur] = struct{}{}

			// Add any unseen bags contained in this bag to list of bags to check
			for containedBagID := range bags[cur] {
				if _, exists := seen[containedBagID]; !exists {
					check[containedBagID] = struct{}{}
				}
			}

			// Remove current from list of bags to check
			delete(check, cur)
		}
	}

	return containsBag
}

func Test_BagParser(t *testing.T) {
	for _, tc := range []struct {
		in      string
		want    ColoredBag
		wantErr error
	}{
		{
			"light red bags contain 1 bright white bag, 2 muted yellow bags.",
			ColoredBag{
				"light red",
				map[string]int{
					"bright white": 1,
					"muted yellow": 2,
				},
			},
			nil,
		},
		{
			"faded blue bags contain no other bags.",
			ColoredBag{
				"faded blue",
				map[string]int{},
			},
			nil,
		},
	} {
		t.Run(fmt.Sprintf("%v", tc.in), func(t *testing.T) {
			res, err := ParseBag(tc.in)
			require.NoError(t, err)
			require.Equal(t, tc.want, res)
		})
	}
}

type BagParser struct {
	s      []rune
	curIdx int
}

type ColoredBag struct {
	id         string
	bagsInside map[string]int
}

func ParseBag(row string) (ColoredBag, error) {
	bp := BagParser{s: []rune(row)}
	return bp.ParseContainerBag()
}

func (p *BagParser) ParseContainerBag() (res ColoredBag, err error) {
	res.id, err = p.ParseBagID()
	if err != nil {
		return
	}
	if err = p.ParseWhitespace(); err != nil {
		return
	}
	if err = p.ParseContainSeparator(); err != nil {
		return
	}
	res.bagsInside, err = p.ParseContainedBags()
	return
}

// Parse the BagID. Expected to be in this format:
// "some bag name"
func (p *BagParser) ParseBagID() (s string, err error) {
	bagWords := make([]string, 2)

	for i := 0; i < 2; i++ {
		// Parse word
		bagWords[i], err = p.ParseWord()
		if err != nil {
			return
		}

		// Parse whitespace
		if err = p.ParseWhitespace(); err != nil {
			return s, fmt.Errorf("failed to parse bag id, %v", err)
		}
	}

	// Parse "bag[s]" - throw it away
	if _, err = p.ParseWord(); err != nil {
		return
	}

	return strings.Join(bagWords, " "), nil
}

// ParseWord parses a single word.
// The first character must be a letter.
// Breaks on ' ', '.' and ','.
func (p *BagParser) ParseWord() (string, error) {
	if !unicode.IsLetter(p.s[p.curIdx]) {
		return "", fmt.Errorf("non-character at beginning of a word: %q, s: %s", p.s[p.curIdx], string(p.s[p.curIdx:]))
	}
	// First character is a letter
	// Parse until ' ', '.' or ','
	// Result will be p.s[startIdx:p.curIdx]
	startIdx := p.curIdx
	for {
		switch p.s[p.curIdx] {
		case ' ', '.', ',':
			return string(p.s[startIdx:p.curIdx]), nil
		}
		if !unicode.IsLetter(p.s[p.curIdx]) {
			return "", fmt.Errorf("non-letter within a word: %v, %v", p.s[p.curIdx], p.s[p.curIdx:])
		}
		p.curIdx++
	}
}

func (p *BagParser) ParseContainSeparator() error {
	word, err := p.ParseWord()
	if err != nil {
		return err
	}
	if word != "contain" {
		return fmt.Errorf("expected 'contain', got %v", word)
	}
	// skip whitespace
	p.curIdx++
	return nil
}

func (p *BagParser) ParseWhitespace() error {
	if p.s[p.curIdx] != ' ' {
		return fmt.Errorf("failed to parse whitespace: %v, %v", p.curIdx, p.s[p.curIdx:])
	}
	p.curIdx++
	return nil
}

func (p *BagParser) ParseContainedBags() (bagAndCount map[string]int, err error) {
	bagCount := make(map[string]int)

	for {
		// Parse BagCount
		switch p.s[p.curIdx] {
		case 'n':
			return bagCount, nil
		default:
			count, id, err := p.ParseBagCount()
			if err != nil {
				return nil, err
			}
			bagCount[id] = count
		}

		// Parse separator
		switch p.s[p.curIdx] {
		case ',':
			p.curIdx += 2
			continue
		case '.':
			return bagCount, nil
		}
	}
}

func (p *BagParser) ParseBagCount() (count int, id string, err error) {
	count, err = p.ParseDigit()
	if err != nil {
		return
	}
	if err = p.ParseWhitespace(); err != nil {
		return
	}
	id, err = p.ParseBagID()
	return
}

func (p *BagParser) ParseDigit() (int, error) {
	if !unicode.IsDigit(p.s[p.curIdx]) {
		return 0, fmt.Errorf("non-digit found when parsing count: %q, %v", p.s[p.curIdx], string(p.s[p.curIdx:]))
	}
	d := int(p.s[p.curIdx] - '0')
	p.curIdx++
	return d, nil
}
