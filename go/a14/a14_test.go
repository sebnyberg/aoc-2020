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

func Test_day14part1(t *testing.T) {
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
	require.Equal(t, 14722016054794, sum)
}

type floatFunc func(bits []int) []int

func Test_day14part2(t *testing.T) {
	f, err := os.Open("input")
	check(err)

	createPerm := func(i int) func([]int) []int {
		return func(bits []int) []int {
			res := make([]int, len(bits)*2)
			for j := range bits {
				res[2*j] = bits[j] & ^(1 << (35 - i))
				res[2*j+1] = bits[j] | 1<<(35-i)
			}
			return res
		}
	}

	mem := make(map[int]int, 1000)
	oneMask := 0
	floatMask := 0
	sc := bufio.NewScanner(f)
	floatFuncs := make([]floatFunc, 0)
	for sc.Scan() {
		row := sc.Text()
		parts := strings.Split(row, " = ")
		instr, valstr := parts[0], parts[1]
		if instr == "mask" {
			// Clear
			oneMask = 0
			floatMask = 0
			floatFuncs = floatFuncs[:0]
			for i, ch := range valstr {
				switch ch {
				case 'X':
					floatFuncs = append(floatFuncs, createPerm(i))
					floatMask |= 1 << (35 - i)
				case '1':
					oneMask |= 1 << (35 - i)
				case '0':
					continue
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
		addr |= oneMask

		// Clear floating bits
		addr &= ^floatMask

		// Pass through float permutations
		addrs := []int{addr}
		for _, perm := range floatFuncs {
			addrs = perm(addrs)
		}

		for _, addr := range addrs {
			mem[addr] = val
		}
	}

	sum := 0
	for _, val := range mem {
		sum += val
	}
	require.Equal(t, 3618217244644, sum)
}
