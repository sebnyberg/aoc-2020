package a18

import (
	"fmt"
	"strconv"
)

type op int

const (
	opNoop op = iota
	opAdd
	opSub
	opMul
	opDiv
)

func (o op) String() string {
	switch o {
	case opNoop:
		return "NOOP"
	case opAdd:
		return "+"
	case opSub:
		return "-"
	case opMul:
		return "*"
	case opDiv:
		return "/"
	default:
		panic("unknown op")
	}
}

// factor = num | "(" expr ")"
type factor struct {
	n int
	e *expr
}

func (f *factor) String() string {
	if f.e != nil {
		return f.e.String()
	}
	return fmt.Sprintf("%v", f.n)
}

func (f *factor) value() int {
	if f.e != nil {
		return f.e.value()
	}
	return f.n
}

// term = factor | factor mulOp term
type term struct {
	lhs *factor
	op  op
	rhs *term
}

func (t *term) String() string {
	if t.op == opNoop {
		return t.lhs.String()
	}
	return fmt.Sprintf("%v %v %v", t.lhs, t.op, t.rhs)
}

func (t *term) value() int {
	switch t.op {
	case opNoop:
		return t.lhs.value()
	case opAdd:
		return t.lhs.value() + t.rhs.value()
	case opSub:
		return t.lhs.value() - t.rhs.value()
	default:
		panic("failed to evaluate expression value")
	}
}

type expr struct {
	lhs *term
	op  op
	rhs *expr
}

func (e *expr) String() string {
	if e.op == opNoop {
		return e.lhs.String()
	}
	return fmt.Sprintf("%v %v %v", e.lhs, e.op, e.rhs)
}

func (e *expr) value() int {
	switch e.op {
	case opNoop:
		return e.lhs.value()
	case opMul:
		return e.lhs.value() * e.rhs.value()
	case opDiv:
		return e.lhs.value() / e.rhs.value()
	default:
		panic("failed to evaluate expression value")
	}
}

// Grammar:
// num = "0" | "1" | ... | "9" ;
// mulOp = "*" | "/"
// addOp = "+" | "-"
// factor = num | "(" expr ")" ;
// term = factor | factor addOp term
// expr = term | term mulOp expr ;
func parse(in string) *expr {
	p := parser{
		tokens: lex(in),
	}
	return p.parseExpr()
}

type parser struct {
	tokens []token
	pos    int
}

// expr = term | term mulOp expr
func (p *parser) parseExpr() *expr {
	var e expr
	e.lhs = p.parseTerm()
	t := p.next()
	switch t.typ {
	case tokenEOF:
		return &e
	case tokenMul:
		e.op = opMul
	case tokenDiv:
		e.op = opDiv
	default:
		p.backup()
		return &e
	}
	e.rhs = p.parseExpr()
	return &e
}

// term = factor | factor addOp term
func (p *parser) parseTerm() *term {
	var te term
	te.lhs = p.parseFactor()
	t := p.next()
	switch t.typ {
	case tokenEOF:
		return &te
	case tokenAdd:
		te.op = opAdd
	case tokenSub:
		te.op = opSub
	default: // bubble up when mul / div
		p.backup()
		return &te
	}
	te.rhs = p.parseTerm()
	return &te
}

// factor = num | "(" expr ")"
func (p *parser) parseFactor() *factor {
	var f factor

	t := p.next()
	switch t.typ {
	case tokenEOF:
		panic("got eofToken when parsing factor")
	case tokenNum:
		var err error
		f.n, err = strconv.Atoi(t.val)
		check(err)
	case tokenLParen:
		f.e = p.parseExpr()
		if t = p.next(); t.typ != tokenRParen {
			panic(fmt.Sprintf("factor was not closed by right parenthesis, got: %v", t))
		}
	default:
		panic(fmt.Sprintf("invalid token when parsing factor: %v", t))
	}

	return &f
}

func (p *parser) next() (t token) {
	if p.pos >= len(p.tokens) {
		return token{typ: tokenEOF}
	}
	t = p.tokens[p.pos]
	p.pos++
	return t
}

func (p *parser) backup() {
	p.pos--
}
