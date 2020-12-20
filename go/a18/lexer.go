package a18

import (
	"fmt"
	"log"
	"strings"
	"unicode/utf8"
)

type tokenType int

const (
	tokenErr tokenType = iota
	tokenEOF
	tokenAdd
	tokenSub
	tokenMul
	tokenDiv
	tokenLParen
	tokenRParen
	tokenNum
)

type token struct {
	typ tokenType
	val string
}

func (i token) String() string {
	switch i.typ {
	case tokenEOF:
		return "EOF"
	case tokenErr:
		return i.val
	case tokenAdd:
		return "+"
	case tokenSub:
		return "-"
	case tokenMul:
		return "*"
	case tokenDiv:
		return "/"
	case tokenLParen:
		return "("
	case tokenRParen:
		return ")"
	case tokenNum:
		return fmt.Sprintf("%v", i.val)
	default:
		panic(fmt.Sprintf("unknown token type %v", i.typ))
	}
}

type lexer struct {
	tokens []token // read tokens
	in     string  // string to lex
	start  int     // start position (since last emit / ignore)
	pos    int     // current position
	width  int     // width of last read rune
}

const eof = rune(0)

type stateFn func(l *lexer) stateFn

func lex(in string) []token {
	l := &lexer{
		in:     in,
		tokens: make([]token, 0),
	}
	for state := lexText(l); state != nil; {
		state = state(l)
	}
	return l.tokens
}

func lexText(l *lexer) stateFn {
	for {
		switch ch := l.next(); {
		case ch == eof:
			l.emit(tokenEOF)
			return nil
		case ch == ' ':
			l.ignore()
		case ch == '(':
			l.emit(tokenLParen)
		case ch == ')':
			l.emit(tokenRParen)
		case ch == '+':
			l.emit(tokenAdd)
		case ch == '-':
			l.emit(tokenSub)
		case ch == '*':
			l.emit(tokenMul)
		case ch == '/':
			l.emit(tokenDiv)
		case ch >= '0' && ch <= '9':
			l.acceptRun("0123456789")
			l.emit(tokenNum)
		default:
			return l.errorf("invalid character: %v", string(ch))
		}
	}

}

// next reads and returns the next rune, returning
// eof if no more runes can be read.
func (l *lexer) next() (ch rune) {
	if l.pos >= len(l.in) {
		l.width = 0
		return eof
	}

	ch, l.width = utf8.DecodeRuneInString(l.in[l.pos:])
	l.pos += l.width
	return ch
}

// ignore read characters.
func (l *lexer) ignore() {
	l.start = l.pos
}

// backup from last read rune.
func (l *lexer) backup() {
	l.pos -= l.width
}

// acceptRun accepts runes as long as they match the provided string.
func (l *lexer) acceptRun(valid string) {
	for strings.ContainsRune(valid, l.next()) {
	}
	l.backup()
}

// emit the currently read block to the list of tokens
func (l *lexer) emit(typ tokenType) {
	l.tokens = append(l.tokens, token{
		typ: typ,
		val: l.in[l.start:l.pos],
	})
	l.start = l.pos
}

// errorf emits an error to the list of tokens
func (l *lexer) errorf(format string, args ...interface{}) stateFn {
	l.tokens = append(l.tokens, token{
		tokenErr,
		fmt.Sprintf(format, args...),
	})
	return nil
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
