package a04_test

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
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

type PassportField string

const (
	BirthYear      PassportField = "byr"
	IssueYear      PassportField = "iyr"
	ExpirationYear PassportField = "eyr"
	Height         PassportField = "hgt"
	HairColor      PassportField = "hcl"
	EyeColor       PassportField = "ecl"
	PassportID     PassportField = "pid"
	CountryID      PassportField = "cid"
)

type Passport struct {
	BirthYear      int
	IssueYear      int
	ExpirationYear int
	Height         string
	HairColor      string
	EyeColor       string
	PassportID     string
	CountryID      string
}

func (p *Passport) Validate() error {
	errs := make([]error, 0)
	appendErr := func(err error) {
		if err != nil {
			errs = append(errs, err)
		}
	}
	appendErr(validateBirthYear(p.BirthYear))
	appendErr(validateIssueYear(p.IssueYear))
	appendErr(validateExpirationYear(p.ExpirationYear))
	appendErr(validateHeight(p.Height))
	appendErr(validateHairColor(p.HairColor))
	appendErr(validateEyeColor(p.EyeColor))
	appendErr(validateCountryID(p.CountryID))
	appendErr(validatePassportID(p.PassportID))
	if len(errs) > 0 {
		err := errs[0]
		for _, newErr := range errs[1:] {
			err = fmt.Errorf("%v, %w", err, newErr)
		}
		return err
	}
	return nil
}

func validateBirthYear(year int) error {
	if year == 0 {
		return errors.New("birth year is required")
	}
	if year < 1920 || year > 2002 {
		return fmt.Errorf("invalid birth year: %v", year)
	}
	return nil
}

func validateIssueYear(year int) error {
	if year == 0 {
		return errors.New("issue year is required")
	}
	if year < 2010 || year > 2020 {
		return fmt.Errorf("invalid issue year: %v", year)
	}
	return nil
}

func validateExpirationYear(year int) error {
	if year == 0 {
		return errors.New("expiration year is required")
	}
	if year < 2020 || year > 2030 {
		return fmt.Errorf("invalid expiration year: %v", year)
	}
	return nil
}

var heightRegex = *regexp.MustCompile(`^(\d+)(cm|in)$`)

func validateHeight(heightEntry string) error {
	if heightEntry == "" {
		return errors.New("height is required")
	}
	parts := heightRegex.FindStringSubmatch(heightEntry)
	if len(parts) != 3 {
		return fmt.Errorf("invalid height %v", heightEntry)
	}
	heightStr, unit := parts[1], parts[2]
	height, err := strconv.Atoi(heightStr)
	if err != nil {
		return fmt.Errorf("failed to parse height %v, %v", height, err)
	}
	switch unit {
	case "cm":
		if height < 150 || height > 193 {
			return fmt.Errorf("invalid height: %v", height)
		}
	case "in":
		if height < 59 || height > 76 {
			return fmt.Errorf("invalid height: %v", height)
		}
	default:
		return fmt.Errorf("invalid height unit: %v", unit)
	}
	return nil
}

var hairColorRegex = *regexp.MustCompile(`^#[0-9a-f]{6}$`)

func validateHairColor(color string) error {
	if color == "" {
		return errors.New("hair color is required")
	}
	if !hairColorRegex.MatchString(color) {
		return fmt.Errorf("invalid hair color: %v", color)
	}
	return nil
}

func validateEyeColor(color string) error {
	if color == "" {
		return errors.New("eye color is required")
	}
	switch color {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		return nil
	}
	return fmt.Errorf("invalid eye color: %v", color)
}

var passportIDRegex = *regexp.MustCompile(`^\d{9}$`)

func validatePassportID(id string) error {
	if !passportIDRegex.MatchString(id) {
		return fmt.Errorf("passport ID %v did not meet requirements", id)
	}
	return nil
}

func validateCountryID(id string) error {
	return nil
}

var exampleInput = `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`

var examplePassports = []Passport{
	{
		BirthYear:      1937,
		CountryID:      "147",
		ExpirationYear: 2020,
		EyeColor:       "gry",
		HairColor:      "#fffffd",
		Height:         "183cm",
		IssueYear:      2017,
		PassportID:     "860033327",
	},
	{
		BirthYear:      1929,
		CountryID:      "350",
		ExpirationYear: 2023,
		EyeColor:       "amb",
		HairColor:      "#cfa07d",
		IssueYear:      2013,
		PassportID:     "028048884",
	},
	{
		BirthYear:      1931,
		ExpirationYear: 2024,
		EyeColor:       "brn",
		HairColor:      "#ae17e1",
		Height:         "179cm",
		IssueYear:      2013,
		PassportID:     "760753108",
	},
	{
		ExpirationYear: 2025,
		EyeColor:       "brn",
		HairColor:      "#cfa07d",
		Height:         "59in",
		IssueYear:      2011,
		PassportID:     "166559648",
	},
}

func Test_Validate(t *testing.T) {
	// f := bytes.NewBufferString(exampleInput)
	f, err := os.Open("input")
	check(err)

	passports := ParsePassportsFromInput(f)
	valid := 0
	for i, p := range passports {
		if err := p.Validate(); err != nil {
			t.Logf("[%d] INVALID - %v", i, err)
			continue
		}
		t.Logf("[%d] VALID", i)
		valid++
	}
	require.Equal(t, 184, valid)
}

func Test_ParseInput(t *testing.T) {
	f := bytes.NewBufferString(exampleInput)
	passports := ParsePassportsFromInput(f)
	for i := range passports {
		require.Equal(t, examplePassports[i], passports[i])
	}
}

func ParsePassportsFromInput(r io.Reader) []Passport {
	passportTexts := ParsePassportTexts(r)
	passports := make([]Passport, len(passportTexts))
	for i := range passportTexts {
		passports[i] = ParsePassportFromText(passportTexts[i])
	}
	return passports
}

func ParsePassportFromText(s string) (p Passport) {
	parts := strings.Split(s, " ")
	for _, part := range parts {
		fieldParts := strings.Split(part, ":")
		if len(fieldParts) != 2 {
			log.Fatalln(fieldParts)
		}
		k, v := fieldParts[0], fieldParts[1]
		var err error
		switch PassportField(k) {
		case BirthYear:
			p.BirthYear, err = strconv.Atoi(v)
		case IssueYear:
			p.IssueYear, err = strconv.Atoi(v)
		case ExpirationYear:
			p.ExpirationYear, err = strconv.Atoi(v)
		case Height:
			p.Height = v
		case HairColor:
			p.HairColor = v
		case EyeColor:
			p.EyeColor = v
		case PassportID:
			p.PassportID = v
		case CountryID:
			p.CountryID = v
		}
		if err != nil {
			log.Println("faied to parse year", err)
		}
	}
	return
}

func ParsePassportTexts(f io.Reader) []string {
	sc := bufio.NewScanner(f)
	passportRows := make([][]string, 0)
	passportIdx := 0
	passportRows = append(passportRows, make([]string, 0))
	for sc.Scan() {
		row := sc.Text()
		if row == "" {
			passportIdx++
			passportRows = append(passportRows, make([]string, 0))
			continue
		}
		passportRows[passportIdx] = append(passportRows[passportIdx], row)
	}
	passportTexts := make([]string, len(passportRows))
	for i, p := range passportRows {
		passportTexts[i] = strings.Join(p, " ")
	}
	return passportTexts
}
