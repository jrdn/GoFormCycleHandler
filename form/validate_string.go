package form

import (
	"regexp"
)

var rxEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// IsStringValidator checks that the field value can be cast into a Go string
var IsStringValidator = Validator{
	ID: "isString",
	F: func(field *Field) (results ValidationResults) {
		if _, ok := field.Value.(string); !ok {
			results.Add("cannot parse value as string")
		}
		return
	},
}

var IsStringListValidator = Validator{
	ID: "isStringList",
	F: func(field *Field) (results ValidationResults) {
		if _, ok := field.Value.([]string); !ok {
			results.Add("cannot parse value as string list")
		}

		return
	},
}

var StringListValuesInMultiValidator = Validator{
	ID: "stringListValuesInMulti",
	F: func(field *Field) (results ValidationResults) {
		for _, v := range field.Value.([]string) {
			found := false
			for _, m := range field.Options[MultiOptionKey].([]string)  {
				if v == m {
					found = true
					break
				}
			}
			if !found {
				results.Add("value %v not allowed", v)
			}
		}
		return
	},
}

// StringMinLengthValidator returns a validator which checks that the form value is a string & is longer than the
// specified value.
func StringMinLengthValidator(min int) Validator {
	return Validator{
		ID: "stringMinLength",
		F: func(field *Field) (results ValidationResults) {
			r := IsStringValidator.F(field)
			if !r.IsEmpty() {
				return r
			}

			s := field.Value.(string)
			if len(s) < min {
				results.Add("string shorter than minimum length: %d", min)
			}

			return
		},
	}
}

// EmailValidator checks that the field value is an email address
var EmailValidator = Validator{
	ID: "email",
	F: func(field *Field) (results ValidationResults) {

		r := IsStringValidator.F(field)
		if !r.IsEmpty() { return r }

		s := field.Value.(string)
		if len(s) > 254 {
			results.Add("email address is too long")
			return
		}

		if !rxEmail.MatchString(s) {
			results.Add("email address is invalid")
		}

		return
	},
}
