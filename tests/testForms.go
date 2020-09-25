package tests

import "github.ol.epicgames.net/cloud-eng/goformcyclehandler/form"

var TestForms = []*form.Form{
	{Name: "test1",
		Fields: []*form.Field{
			{ID: "field1", Type: form.StringType, Validators: []form.Validator{form.NotNullValidator, form.IsStringValidator}},
			{ID: "field2", Type: form.StringType, Validators: []form.Validator{form.NotNullValidator, form.IsStringValidator}},
		},
	},
}

func GetTestFormNames() []string {
	names := make([]string, len(TestForms))
	for i, t := range TestForms {
		names[i] = t.Name
	}
	return names
}

