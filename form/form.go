package form

import "encoding/json"

type Form struct {
	Name string `json:"name"`
	Fields []*Field `json:"fields"`
}

type ValidationError struct {
	Msg         string `json:"msg"`
	ValidatorID string `json:"validator_id"`
}

func (ve ValidationError) Error() string {
	return ve.Msg
}

func (f *Form) Validate() bool {
	result := true
	for _, field := range f.Fields {
		field.Errors = []ValidationError{}
		for _, v := range field.Validators {
			results := v.F(field)
			for _, msg := range results.Get() {
				result = false
				field.Errors = append(field.Errors, ValidationError{
					Msg:         msg,
					ValidatorID: v.ID,
				})
			}
		}
	}
	return result
}

func (f *Form) FromJSON(data []byte) error {
	return json.Unmarshal(data, &f)
}

func (f *Form) Clone() *Form {
	copy := &Form{
		Name:   f.Name,
		Fields: make([]*Field, len(f.Fields)),
	}

	for i, field := range f.Fields {
		copy.Fields[i] = field.Clone()
	}
	return copy
}



