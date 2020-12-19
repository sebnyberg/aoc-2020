package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Println(lex("1 + (2 * 3 ) * 3"))
}

type item struct {
	typ itemType
	val string
}

func (i item) String() string {
	switch i.typ {
	case itemTypeEOF:
		return "EOF"
	case itemTypeErr:
		return i.val
	}
	if len(i.val) > 10 {
		return fmt.Sprintf("%.10q...", i.val)
	}
	return fmt.Sprintf("%q", i.val)
}

type itemType int

const (
	itemTypeEOF itemType = iota
	itemTypeErr
	itemTypeMul
	itemTypeDiv
	itemTypeAdd
	itemTypeSub
	itemTypeLParen
	itemTypeRParen
	itemTypeNum
)

func lex(in string) []item {
	l := &lexer{
		input: in,
		items: make([]item, 0),
	}
	for state := lexText; state != nil; {
		state = state(l)
	}
	return l.items
}

type lexer struct {
	items []item
	input string
	width int
	pos   int
	start int
}

type stateFn func(l *lexer) stateFn

const eof = rune(0)

func lexText(l *lexer) stateFn {
	for {
		r := l.next()
		if r == eof {
			fmt.Println("EOF!!")
		}
		switch r := l.next(); {
		case r == eof || r == '\n':
			fmt.Println("hehe!!")
			return nil
		case r == ' ':
			l.ignore()
		case r == '(':
			l.emit(itemTypeLParen)
		case r == ')':
			l.emit(itemTypeRParen)
		case r == '*':
			l.emit(itemTypeMul)
		case r == '/':
			l.emit(itemTypeDiv)
		case r >= '0' && r <= '9':
			for l.accept("0123456789") {
			}
			l.emit(itemTypeNum)
		}
	}
}

// Put an item into the list of items int he lexer
func (l *lexer) emit(typ itemType) {
	l.items = append(l.items, item{
		typ: typ,
		val: l.input[l.start:l.pos],
	})
	l.start = l.pos
}

func (l *lexer) ignore() {
	l.start = l.pos
}

func (l *lexer) accept(valid string) bool {
	if strings.IndexRune(valid, l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}

func (l *lexer) backup() {
	l.pos -= l.width
}

func (l *lexer) next() (ch rune) {
	if l.pos >= len(l.input) {
		fmt.Println("returning eof")
		return eof
	}

	ch, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return ch
}
