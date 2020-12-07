package a06_test

import (
	"fmt"
	"log"
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

var testInput = `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`

func Test_BagParser(t *testing.T) {
	for _, tc := range []struct {
		in      string
		want    ContainerBag
		wantErr error
	}{
		{
			"light red bags contain 1 bright white bag, 2 muted yellow bags.",
			ContainerBag{
				"light red bags",
				[]BagCount{
					{1, "bright white bag"},
					{2, "muted yellow bags"},
				},
			},
			nil,
		},
		{
			"faded blue bags contain no other bags.",
			ContainerBag{
				"faded blue bags",
				[]BagCount{},
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

type BagCount struct {
	count int
	bagID string
}

type ContainerBag struct {
	bagID         string
	containedBags []BagCount
}

func ParseBag(row string) (ContainerBag, error) {
	bp := BagParser{s: []rune(row)}
	return bp.ParseContainerBag()
}

func (p *BagParser) ParseContainerBag() (res ContainerBag, err error) {
	res.bagID, err = p.ParseBagID()
	if err != nil {
		return
	}
	if err = p.ParseWhitespace(); err != nil {
		return
	}
	if err = p.ParseContainSeparator(); err != nil {
		return
	}
	res.containedBags, err = p.ParseContainedBags()
	return
}

// Parse the BagID. Expected to be in this format:
// "some bag name"
func (p *BagParser) ParseBagID() (s string, err error) {
	bagWords := make([]string, 3)

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

	// Parse final word (do not parse the trailing ',' or '.')
	bagWords[2], err = p.ParseWord()
	if err != nil {
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

func (p *BagParser) ParseContainedBags() ([]BagCount, error) {
	bagCount := make([]BagCount, 0)

	for {
		// Parse BagCount
		switch p.s[p.curIdx] {
		case 'n':
			return bagCount, nil
		default:
			newBag, err := p.ParseBagCount()
			if err != nil {
				return nil, err
			}
			bagCount = append(bagCount, newBag)
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

func (p *BagParser) ParseBagCount() (bagCount BagCount, err error) {
	bagCount.count, err = p.ParseDigit()
	if err != nil {
		return
	}
	if err = p.ParseWhitespace(); err != nil {
		return
	}
	bagCount.bagID, err = p.ParseBagID()
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
