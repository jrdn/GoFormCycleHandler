package main

import (
	"net/http"

	GoFormCycleHandler "github.ol.epicgames.net/cloud-eng/goformcyclehandler"
	"github.ol.epicgames.net/cloud-eng/goformcyclehandler/form"
)

var f = form.Form{
	Name: "ExampleForm",
	Fields: []*form.Field{
		{
			ID:    "name",
			Type:  form.StringType,
			Value: "b",
			Validators: []form.Validator{
				form.NotNullValidator,
				form.IsStringValidator,
				form.StringMinLengthValidator(5),
			},
		},
		{
			ID:   "age",
			Type: form.IntType,
			Validators: []form.Validator{
				form.IsIntValidator,
			},
		},
		{
			ID:   "price",
			Type: form.FloatType,
			Validators: []form.Validator{
				form.IsFloatValidator,
			},
		},
		{
			ID:    "select",
			Type:  form.StringListType,
			Value: []string{"asdf"},
			Validators: []form.Validator{
				form.IsStringListValidator,
				form.StringListValuesInMultiValidator,
			},
			Options: map[form.OptionKey]interface{}{
				form.MultiOptionKey: []string{"foo", "bar", "baz"},
			},
		},
	},
}

func main() {
	var completion GoFormCycleHandler.FormCompletionFunc = func(f *form.Form) error {
		return nil
	}

	formHandler := GoFormCycleHandler.NewFormHandler()
	formHandler.Register(&f, completion)

	mux := http.NewServeMux()
	mux.Handle("/form/", formHandler)
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	_ = server.ListenAndServe()
}
