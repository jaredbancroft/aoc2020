package password

import (
	"regexp"
	"strconv"
	"strings"
)

//Validator holds information needed to validate a password
type Validator struct {
	min      int
	max      int
	char     string
	password string
}

//NewValidator creates a new password validator
func NewValidator(p string) *Validator {
	re := regexp.MustCompile(`([0-9]*)-([0-9]*).([a-z]):.([a-z]*)`)
	cg := re.FindStringSubmatch(p)
	min, _ := strconv.Atoi(cg[1])
	max, _ := strconv.Atoi(cg[2])
	char := cg[3]
	password := cg[4]

	return &Validator{min: min, max: max, char: char, password: password}
}

//ValidateNumber validates a password in a Validator object for the number of a given char
func (v *Validator) ValidateNumber() bool {
	count := strings.Count(v.password, v.char)
	if count >= v.min && count <= v.max {
		return true
	}
	return false
}

//ValidatePosition validates a password in a Validator object for the number of a given char
func (v *Validator) ValidatePosition() bool {
	if (string(v.password[v.min-1]) == v.char && string(v.password[v.max-1]) != v.char) ||
		(string(v.password[v.min-1]) != v.char && string(v.password[v.max-1]) == v.char) {
		return true
	}
	return false
}
