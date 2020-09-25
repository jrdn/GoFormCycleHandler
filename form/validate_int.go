package form

// IsIntValidator checks that the field value can be cast into a Go int
var IsIntValidator = Validator{
	ID: "isInt",
	F: func(field *Field) (results ValidationResults) {
		if _, ok := field.Value.(int); !ok {
			results.Add("cannot parse value as int")
		}
		return
	},
}