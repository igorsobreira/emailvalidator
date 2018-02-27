package emailvalidator

import "regexp"

// Format is the default format used to validate emails
const Format = `(?i:\A\s*([^@\s]{1,64})@((?:[-a-z0-9]+\.)+[a-z]{2,})\s*\z)`

// DefaultValidator used by Valid() function using the default Format
var DefaultValidator = &Validator{Pattern: regexp.MustCompile(Format)}

// Valid returns true if email matches Format regex
func Valid(email string) bool {
	return DefaultValidator.Valid(email)
}

type Validator struct {
	Pattern *regexp.Regexp
}

func (v *Validator) Valid(email string) bool {
	return v.Pattern.MatchString(email)
}
