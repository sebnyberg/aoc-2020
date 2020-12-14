package a14_test

import (
	"bufio"
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

func Test_parseInput(t *testing.T) {
	f, err := os.Open("input")
	check(err)

	mem := make(map[int]int, 1000)
	oneMask := 0
	zeroMask := 0
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		row := sc.Text()
		parts := strings.Split(row, " = ")
		instr, valstr := parts[0], parts[1]
		if instr == "mask" {
			// Clear
			oneMask = 0
			zeroMask = 0
			for i, ch := range valstr {
				switch ch {
				case 'X':
					continue
				case '1':
					oneMask |= 1 << (35 - i)
				case '0':
					zeroMask |= 1 << (35 - i)
				default:
					panic(ch)
				}
			}
			continue
		}

		// "mem" instruction
		addrstr := strings.TrimRight(strings.TrimLeft(instr, "me["), "]")
		addr, err := strconv.Atoi(addrstr)
		check(err)
		val, err := strconv.Atoi(valstr)
		check(err)

		// Apply mask
		val |= oneMask
		val &= ^zeroMask
		mem[addr] = val
	}

	sum := 0
	for _, val := range mem {
		sum += val
	}
	require.Equal(t, 0, sum)

	t.FailNow()
}
