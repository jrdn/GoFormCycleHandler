package form

type TypeHint string

const (
	StringType     TypeHint = "string"
	StringListType TypeHint = "stringlist"
	IntType                 = "int"
	FloatType               = "float"
)

type OptionKey string

const (
	MultiOptionKey OptionKey = "multi"
)

type Field struct {
	ID    string `json:"id"`
	Label string `json:"label,omitempty"`

	Value interface{} `json:"value,omitempty"`
	Type  TypeHint    `json:"type"`

	Validators []Validator       `json:"-"`
	Errors     []ValidationError `json:"errors,omitempty"`

	Options map[OptionKey]interface{} `json:"options,omitempty"`
}

func (f *Field) Clone() *Field {
	copy := &Field{
		ID:         f.ID,
		Label:      f.Label,
		Value:      f.Value,
		Type:       f.Type,
		Validators: make([]Validator, len(f.Validators)),
		Errors:     make([]ValidationError, len(f.Errors)),
		Options:    make(map[OptionKey]interface{}),
	}

	for i, v := range f.Validators {
		copy.Validators[i] = v
	}
	for i, e := range f.Errors {
		copy.Errors[i] = e
	}
	for k, v := range f.Options {
		copy.Options[k] = v
	}
	return copy

}
