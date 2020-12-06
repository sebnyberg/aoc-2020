package a04

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	BirthYear      int // byr (Birth Year)
	IssueYear      int // iyr (Issue Year)
	ExpirationYear int // eyr (Expiration Year)

	Height string // hgt (Height)

	EyeColor  string // ecl (Eye Color)
	HairColor string // hcl (Hair Color)

	PassportID string // pid (Passport ID)
	CountryID  string // cid (Country ID)

	PrintErr bool
}

var (
	hgtRegex = *regexp.MustCompile(`(\d+)(\w+)`)
	hclRegex = *regexp.MustCompile(`^#[a-f0-9]{6}$`)
	pidRegex = *regexp.MustCompile(`[0-9]{9}`)
)

func (p *Passport) Validate() error {
	errs := make([]error, 0)
	appendErr := func(err error) {
		if err != nil {
			errs = append(errs, err)
		}
	}
	if p.BirthYear < 1920 || p.BirthYear > 2020 {
		appendErr(fmt.Errorf("invalid birth year: %v", p.BirthYear))
	}
	if p.IssueYear < 2010 || p.IssueYear > 2030 {
		appendErr(fmt.Errorf("invalid issue year: %v", p.IssueYear))
	}
	if p.ExpirationYear < 2020 || p.ExpirationYear > 2030 {
		appendErr(fmt.Errorf("invalid expiration year: %v", p.ExpirationYear))
	}

	// validate height
	validateHgt := func(hgt string) error {
		if hgt == "" {
			return errors.New("height is required")
		}
		hgtparts := hgtRegex.FindStringSubmatch(hgt)
		if hgtparts == nil || len(hgtparts) != 3 {
			return fmt.Errorf("failed to parse height: %v resulted in %v", p.Height, hgtparts)
		}
		height, err := strconv.Atoi(hgtparts[1])
		if err != nil {
			return fmt.Errorf("failed to convert height %v to an int: %v", hgtparts[1], err)
		}
		unit := hgtparts[2]
		switch unit {
		case "cm":
			if height < 150 || height > 193 {
				return errors.New("height must be 150cm <= height <= 193cm")
			}
		case "in":
			if height < 59 || height > 76 {
				return errors.New("height must be 59in <= height <= 76in")
			}
		default:
			return fmt.Errorf("unknown height unit %v", unit)
		}
		return nil
	}
	appendErr(validateHgt(p.Height))

	// Hair color
	validateHcl := func(hcl string) error {
		if hcl == "" {
			return errors.New("hair color is required")
		}
		if !hclRegex.MatchString(hcl) {
			return errors.New("hair color did not match color regex")
		}
		return nil
	}
	appendErr(validateHcl(p.HairColor))

	// Eye color
	switch p.EyeColor {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		break
	default:
		appendErr(fmt.Errorf("invalid eye color: %v", p.EyeColor))
	}

	// Passport ID
	if !pidRegex.MatchString(p.PassportID) {
		appendErr(fmt.Errorf("passport id %v did not match regex", p.PassportID))
	}
	if len(errs) > 0 {
		compositeErr := errs[0]
		for _, err := range errs[1:] {
			compositeErr = fmt.Errorf("%v\n%w", err, compositeErr)
		}
		return compositeErr
	}
	return nil
}

func ParsePassportRow(row string) (p Passport) {
	passportParts := strings.Split(row, " ")
	for _, part := range passportParts {
		fieldParts := strings.Split(part, ":")
		k, v := fieldParts[0], fieldParts[1]
		switch k {
		case "byr":
			p.BirthYear, _ = strconv.Atoi(v)
		case "iyr":
			p.IssueYear, _ = strconv.Atoi(v)
		case "eyr":
			p.ExpirationYear, _ = strconv.Atoi(v)
		case "cid":
			p.CountryID = v
		case "ecl":
			p.EyeColor = v
		case "hcl":
			p.HairColor = v
		case "hgt":
			p.Height = v
		case "pid":
			p.PassportID = v
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
		fmt.Printf("%+v\n", p)
		if err := p.Validate(); err != nil {
			fmt.Printf("INVALID! err: %v", err)
			continue
		}
		nvalid++
	}
	return nvalid
}

func ScanPassport(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	i := bytes.IndexByte(data, '\n')
	if i != -1 {
		for {
			j := bytes.IndexByte(data[i+1:], '\n')
			if j == -1 {
				break
			}
			if j == 0 {
				return i + 2, bytes.ReplaceAll(data[:i], []byte{'\n'}, []byte{' '}), nil
			}
			i = i + j + 1
		}
	}

	if atEOF {
		return len(data), bytes.ReplaceAll(data, []byte{'\n'}, []byte{' '}), nil
	}

	return 0, nil, nil
}
