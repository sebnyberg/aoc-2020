package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode/utf8"
)

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	parseValue("(1)")
}

type parser struct {
	input string
	pos   int
	start int
	width int
}

type op int

const (
	opNoop op = iota
	opMul
	opDiv
	opAdd
	opSub
)

func (o op) String() string {
	switch o {
	case opNoop:
		return ""
	case opMul:
		return "+"
	case opDiv:
		return "/"
	case opAdd:
		return "+"
	case opSub:
		return "-"
	default:
		return "INVALID_OP"
	}
}

type expr struct {
	left  *term
	op    op
	right *term
}

func (e *expr) String() string {
	if e.op == opNoop {
		return e.left.String()
	}
	return fmt.Sprintf("%v %v %v", e.left, e.op, e.right)
}

type term struct {
	left  *factor
	op    op
	right *factor
}

func (t *term) String() string {
	if t.op == opNoop {
		return t.left.String()
	}
	return fmt.Sprintf("%v %v %v", t.left, t.op, t.right)
}

type factor struct {
	num  int
	expr *expr
}

func (f *factor) String() string {
	if f.expr != nil {
		return fmt.Sprintf("( %v )", f.expr)
	}
	return fmt.Sprintf("%v", f.num)
}

// Grammar:
// num := [0-9]+
// factor := num | "(" expr ")"
// term := factor [ (add|sub) term ]
// expr := term [ (mul|div) term ]
func parseValue(in string) int {
	p := &parser{input: in}
	e := p.parseExpr()
	fmt.Printf("%+v\n", e)
	return 0
}

func (p *parser) parseExpr() *expr {
	var e expr
	e.left = p.parseTerm()
	return &e
}

func (p *parser) parseTerm() *term {
	var t term
	t.left = p.parseFactor()
	return &t
}

func (p *parser) parseFactor() *factor {
	var f factor
	if p.accept("0123456789") {
		n, err := strconv.Atoi(p.input[p.start:p.pos])
		check(err)
		f.num = n
		return &f
	}

	if p.skip("(") {
		f.expr = p.parseExpr()
		if !p.skip(")") {
			panic("found no matching right parenthesis after expression")
		}
		return &f
	}
	panic("encountered non-numeric factor and no lparen")
}

// skip past the provided string, returning true
// if any characters were skipped
func (p *parser) skip(chset string) (skipped bool) {
	for strings.ContainsRune(chset, p.next()) {
		skipped = true
		p.ignore()
	}
	p.backup()
	return skipped
}

func (p *parser) ignore() {
	p.start = p.pos
}

func (p *parser) accept(valid string) bool {
	p.skip(" ")
	fmt.Println("checking ", valid)
	if strings.ContainsRune(valid, p.next()) {
		return true
	}
	p.backup()
	return false
}

func (p *parser) backup() {
	p.pos -= p.width
}

const eof = rune(0)

func (p *parser) next() (ch rune) {
	if p.pos >= len(p.input) {
		return eof
	}
	ch, p.width = utf8.DecodeRuneInString(p.input[p.pos:])
	p.pos += p.width
	return ch
}
