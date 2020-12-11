package a08_test

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

type instruction struct {
	op string
	n  int
}

const (
	INSTR_NOP string = "nop"
	INSTR_ACC string = "acc"
	INSTR_JMP string = "jmp"
)

type console struct {
	acc     int
	pc      int
	pcseen  map[int]struct{}
	program []instruction
}

func Test_console(t *testing.T) {
	input := `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`
	buf := bytes.NewBufferString(input)

	sc := bufio.NewScanner(buf)
	program := make([]instruction, 0)
	for sc.Scan() {
		row := sc.Text()
		rowParts := strings.Split(row, " ")
		instr, nStr := rowParts[0], rowParts[1]
		n, err := strconv.Atoi(nStr)
		check(err)
		program = append(program, instruction{instr, n})
	}

	cons := console{
		pcseen:  make(map[int]struct{}),
		program: program,
	}

	err := cons.Execute()
	require.Error(t, err)
	require.Equal(t, 5, cons.acc)
}

func (c *console) Execute() error {
	for {
		// Check if instruction has been executed before
		if _, executed := c.pcseen[c.pc]; executed {
			return fmt.Errorf("program attempted to execute instruction at %v twice", c.pc)
		}

		// Execute instruction
		c.pcseen[c.pc] = struct{}{}
		instr := c.program[c.pc]
		switch instr.op {
		case INSTR_NOP:
			c.pc++
		case INSTR_ACC:
			c.acc += instr.n
			c.pc++
		case INSTR_JMP:
			c.pc += instr.n
		default:
			log.Fatalln("invalid instruction", instr)
		}
	}
}
