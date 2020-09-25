package form

// IsFloatValidator checks that the field value can be cast into a Go int
var IsFloatValidator = Validator{
	ID: "isFloat",
	F: func(field *Field) (results ValidationResults) {
		if _, ok := field.Value.(int); !ok {
			results.Add("cannot parse value as float")
		}
		return
	},
}