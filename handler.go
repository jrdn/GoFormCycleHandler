package GoFormCycleHandler

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"path"

	"github.ol.epicgames.net/cloud-eng/goformcyclehandler/form"
)

// FormCompletionFunc is called when the form has been filled out and validated so action can be taken based on it
// TODO allow the valid form to call the completion func, but still return an error status if the completion func doesn't succeed
type FormCompletionFunc func(f *form.Form) error

type formRegistration struct {
	f *form.Form
	c FormCompletionFunc
}

func NewFormHandler() FormHandler {
	mux := http.NewServeMux()
	fh := FormHandler{
		registered: make(map[string]formRegistration),
		mux:        mux,
	}

	return fh
}

type FormHandler struct {
	registered map[string]formRegistration
	mux        *http.ServeMux
}

func (fh *FormHandler) Register(f *form.Form, c FormCompletionFunc) {
	fh.registered[f.Name] = formRegistration{
		f: f,
		c: c,
	}

	fh.mux.HandleFunc("/form/"+f.Name, func(w http.ResponseWriter, r *http.Request) {
		fh.handleForm(f.Name, w, r)
	})
}

func (fh *FormHandler) GetForm(name string) (*form.Form, error) {
	if registration, ok := fh.registered[name]; ok {
		return registration.f.Clone(), nil
	}
	return &form.Form{}, errors.New("unknown form: " + name)
}

func (fh *FormHandler) GetFormCompletionFunc(name string) (FormCompletionFunc, error) {
	if registration, ok := fh.registered[name]; ok {
		return registration.c, nil
	}
	return nil, errors.New("unknown form: " + name)
}

func (fh *FormHandler) listForms(w http.ResponseWriter, r *http.Request) {
	// list available forms
	var forms []string
	for name := range fh.registered {
		forms = append(forms, name)
	}

	formList := struct {
		Forms []string `json:"forms"`
	}{
		Forms: forms,
	}

	JSONResponse(formList, w)
}

func (fh *FormHandler) handleForm(id string, w http.ResponseWriter, r *http.Request) {
	f, err := fh.GetForm(id)
	if err != nil {
		panic(err) // TODO
	}

	if r.Body == nil {
		// if there is no request body, return the empty form
		JSONResponse(f, w)
		return
	}

	bodyData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic("TODO")
	}

	err = f.FromJSON(bodyData)
	if err != nil {
		panic(err) // TODO
	}

	f.Validate()
	if f.Validate() {
		c, err := fh.GetFormCompletionFunc(f.Name)
		if err != nil {
			panic(err) // TODO
		}

		err = c(f)
		if err != nil {
			ErrorResponse(http.StatusInternalServerError, err, w)
			return
		}

		return
	}

	// form was filled out but has errors, send it back around to be corrected
	JSONResponse(f, w)

}

func (fh FormHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, lastPart := path.Split(r.URL.Path)
	f, err := fh.GetForm(lastPart)
	if err != nil {
		fh.listForms(w, r)
		return
	}

	fh.handleForm(f.Name, w, r)
}

func JSONResponse(data interface{}, w http.ResponseWriter) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		ErrorResponse(http.StatusInternalServerError, err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(jsonData)
}

func ErrorResponse(httpCode int, err error, w http.ResponseWriter) {
	e := struct {
		Msg string `json:"msg"`
	}{
		Msg: err.Error(),
	}

	jsonData, err := json.Marshal(e)
	if err != nil {
		panic(err) // TODO
	}

	w.WriteHeader(httpCode)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(jsonData)
}
