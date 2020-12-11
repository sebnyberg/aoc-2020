package a08_test

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
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
	// 	in := `nop +0
	// acc +1
	// jmp +4
	// acc +3
	// jmp -3
	// acc -99
	// acc +1
	// jmp -4
	// acc +6`
	// 	f := bytes.NewBufferString(in)

	var err error
	var f io.Reader
	f, err = os.Open("input")
	check(err)

	sc := bufio.NewScanner(f)
	program := make([]instruction, 0)
	// Keep track of jmp locations
	jmpLocations := make([]int, 0)
	var i int
	for sc.Scan() {
		row := sc.Text()
		rowParts := strings.Split(row, " ")
		instr, nStr := rowParts[0], rowParts[1]
		n, err := strconv.Atoi(nStr)
		check(err)
		program = append(program, instruction{instr, n})
		if instr == INSTR_JMP {
			jmpLocations = append(jmpLocations, i)
		}
		i++
	}

	cons := console{
		pcseen:  make(map[int]struct{}),
		program: program,
	}

	// Part 1
	err = cons.Execute()
	require.Error(t, err)
	require.Equal(t, 1217, cons.acc)

	// Part 2
	for _, jmpIndex := range jmpLocations {
		cons.program[jmpIndex].op = INSTR_NOP
		err = cons.Execute()
		if err == nil {
			require.Equal(t, 501, cons.acc)
			return
		}

		// Reset
		cons.program[jmpIndex].op = INSTR_JMP
	}

	t.FailNow()
}

func (c *console) Execute() error {
	c.pc = 0
	c.pcseen = make(map[int]struct{})
	c.acc = 0
	for {
		// Check if instruction has been executed before
		if _, executed := c.pcseen[c.pc]; executed {
			return fmt.Errorf("program attempted to execute instruction at %v twice", c.pc)
		}

		// Check if we are done
		if c.pc == len(c.program) {
			return nil
		}

		// Execute instruction
		c.pcseen[c.pc] = struct{}{}
		instr := c.program[c.pc]
		fmt.Println(instr)
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
