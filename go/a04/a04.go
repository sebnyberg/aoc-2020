package a04

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

type Passport struct {
	byr string // byr (Birth Year)
	cid string // cid (Country ID)
	ecl string // ecl (Eye Color)
	eyr string // eyr (Expiration Year)
	hcl string // hcl (Hair Color)
	hgt string // hgt (Height)
	iyr string // iyr (Issue Year)
	pid string // pid (Passport ID)
}

func (p *Passport) valid() bool {
	if p.byr == "" {
		return false
	}
	if p.ecl == "" {
		return false
	}
	if p.eyr == "" {
		return false
	}
	if p.hcl == "" {
		return false
	}
	if p.hgt == "" {
		return false
	}
	if p.iyr == "" {
		return false
	}
	if p.pid == "" {
		return false
	}
	if p.cid == "" {
		return true
	}
	return true
}

func ParsePassportRow(row string) (p Passport) {
	passportParts := strings.Split(row, " ")
	for _, part := range passportParts {
		k, v := part[:3], part[3:]
		switch k {
		case "byr":
			p.byr = v
		case "cid":
			p.cid = v
		case "ecl":
			p.ecl = v
		case "eyr":
			p.eyr = v
		case "hcl":
			p.hcl = v
		case "hgt":
			p.hgt = v
		case "iyr":
			p.iyr = v
		case "pid":
			p.pid = v
		default:
			panic("unknown field " + v)
		}
	}
	return p
}

func Run(r io.Reader) (nvalid int) {
	sc := bufio.NewScanner(r)
	sc.Split(ScanPassport)
	for sc.Scan() {
		p := ParsePassportRow(sc.Text())
		if p.valid() {
			nvalid++
		}
	}
	return nvalid
}

func ScanPassport(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	i := bytes.IndexByte(data, '\n')
	for {
		j := bytes.IndexByte(data[i+1:], '\n')
		// fmt.Println("----")
		// fmt.Println(string(data[i+1:]))
		// fmt.Println("----")
		// fmt.Println("i", i, "j", j)
		if j == -1 {
			break
		}
		if j == 0 {
			return i + 2, bytes.ReplaceAll(data[:i], []byte{'\n'}, []byte{' '}), nil
		}
		i = i + j + 1
	}

	if atEOF {
		return len(data), bytes.ReplaceAll(data[:len(data)-1], []byte{'\n'}, []byte{' '}), nil
	}

	return 0, nil, nil
}
