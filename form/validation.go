package form

import "fmt"



type ValidationResults struct {
	r []string
}

func (vr *ValidationResults) Add(format string, args ...interface{}) {
	vr.r = append(vr.r, fmt.Sprintf(format, args...))
}

func (vr *ValidationResults) IsEmpty() bool {
	return len(vr.r) == 0
}

func (vr *ValidationResults) Get() []string {
	return vr.r
}

type Validator struct {
	ID string
	F  func(*Field) ValidationResults
}

var NotNullValidator = Validator{
	ID: "notNull",
	F: func(field *Field) (results ValidationResults) {
		if field.Value == nil {
			results.Add("field is null")
		}
		return
	},
}
