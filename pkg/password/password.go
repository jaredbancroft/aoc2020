package password

import (
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
func NewValidator(min int, max int, char string, password string) *Validator {

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
