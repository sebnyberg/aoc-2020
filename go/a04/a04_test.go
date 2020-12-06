package a04_test

import (
	"aoc2020/a04"
	"bufio"
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

var input = `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in
`

func Test_ScanPassport(t *testing.T) {
	f := bytes.NewBufferString(input)
	s := bufio.NewScanner(f)
	s.Split(a04.ScanPassport)
	s.Scan()
	require.Equal(t, `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm`, s.Text())
	s.Scan()
	require.Equal(t, `iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884 hcl:#cfa07d byr:1929`, s.Text())
	s.Scan()
	require.Equal(t, `hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:179cm`, s.Text())
	s.Scan()
	require.Equal(t, `hcl:#cfa07d eyr:2025 pid:166559648 iyr:2011 ecl:brn hgt:59in`, s.Text())
}
