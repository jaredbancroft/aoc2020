package passport

import (
	"regexp"
	"strconv"
)

//Passport object
type Passport struct {
	Byr, Iyr, Eyr, Hgt, Hcl, Ecl, Pid, Cid string
}

//NewPassport creates a new passport object
func NewPassport(fields map[string]string) *Passport {
	byr := fields["byr"]
	iyr := fields["iyr"]
	eyr := fields["eyr"]
	hgt := fields["hgt"]
	hcl := fields["hcl"]
	ecl := fields["ecl"]
	pid := fields["pid"]
	cid := fields["cid"]

	return &Passport{Byr: byr, Iyr: iyr, Eyr: eyr, Hgt: hgt, Hcl: hcl, Ecl: ecl, Pid: pid, Cid: cid}
}

//Validate a passport and return true or false
func (p *Passport) Validate() bool {
	if p.Byr != "" && p.Iyr != "" && p.Eyr != "" && p.Hgt != "" && p.Hcl != "" && p.Ecl != "" && p.Pid != "" {
		return true
	}
	return false
}

//ValidateMore validates harder!
func (p *Passport) ValidateMore() bool {
	if p.validateByr() && p.validateIyr() && p.validateEyr() &&
		p.validateHgt() && p.validateHcl() && p.validateEcl() && p.validatePid() {
		return true
	}
	return false
}

func (p *Passport) validateByr() bool {
	b, err := strconv.Atoi(p.Byr)
	if err != nil {
		return false
	}

	if b >= 1920 && b <= 2002 {
		return true
	}

	return false
}

func (p *Passport) validateIyr() bool {
	b, err := strconv.Atoi(p.Iyr)
	if err != nil {
		return false
	}

	if b >= 2010 && b <= 2020 {
		return true
	}

	return false
}

func (p *Passport) validateEyr() bool {
	b, err := strconv.Atoi(p.Eyr)
	if err != nil {
		return false
	}

	if b >= 2020 && b <= 2030 {
		return true
	}

	return false
}

func (p *Passport) validateHgt() bool {
	re := regexp.MustCompile(`(\d+)(in|cm)`)
	height := re.FindStringSubmatch(p.Hgt)
	var numString string
	var unit string
	if len(height) > 0 {
		numString = height[1]
		unit = height[2]
	} else {
		return false
	}

	num, err := strconv.Atoi(numString)
	if err != nil {
		return false
	}

	switch unit {
	case "in":
		if num >= 59 && num <= 76 {
			return true
		}
	case "cm":
		if num >= 150 && num <= 193 {
			return true
		}
	}
	return false
}

func (p *Passport) validateHcl() bool {
	re := regexp.MustCompile(`#[0-9a-f]{6}`)
	hair := re.FindString(p.Hcl)
	if len(hair) > 0 {
		return true
	}
	return false
}

func (p *Passport) validateEcl() bool {
	eyes := map[string]bool{
		"amb": true,
		"blu": true,
		"brn": true,
		"gry": true,
		"grn": true,
		"hzl": true,
		"oth": true,
	}

	if eyes[p.Ecl] {
		return true
	}
	return false
}

func (p *Passport) validatePid() bool {
	if len(p.Pid) == 9 {
		return true
	}
	return false
}
