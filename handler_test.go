package GoFormCycleHandler

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"github.ol.epicgames.net/cloud-eng/goformcyclehandler/form"
	"github.ol.epicgames.net/cloud-eng/goformcyclehandler/tests"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupTestFormHandler(t *testing.T) FormHandler {
	t.Helper()
	handler := NewFormHandler()
	for _, f := range tests.TestForms {
		handler.Register(f, func(f *form.Form) error {
			t.Log("accepted form " + f.Name)
			return nil
		})
	}
	return handler
}

func TestListForms(t *testing.T) {
	rec := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	require.NoError(t, err)

	handler := setupTestFormHandler(t)
	handler.ServeHTTP(rec, req)
	require.Equal(t, http.StatusOK, rec.Code)

	body, err := ioutil.ReadAll(rec.Body)
	require.NoError(t, err)

	t.Log(string(body))

	l := &struct {
		Forms []string `json:"forms"`
	}{}
	err = json.Unmarshal(body, l)
	require.NoError(t, err)

	require.Len(t, l.Forms, len(tests.TestForms))
	for _, name := range tests.GetTestFormNames() {
		require.Contains(t, l.Forms, name)
	}
}

func TestGetForm(t *testing.T) {
	rec := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/test1", nil)
	require.NoError(t, err)

	handler := setupTestFormHandler(t)
	handler.ServeHTTP(rec, req)
	require.Equal(t, http.StatusOK, rec.Code)

	body, err := ioutil.ReadAll(rec.Body)
	require.NoError(t, err)
	t.Log(string(body))

	// TODO check form content is correct
}

func TestSubmitIncompleteForm(t *testing.T) {
	rec := httptest.NewRecorder()
	requestBody := `
{
	"name":"test1",
	"fields":[
		{"id":"field1","type":"string","ui_hint":""},
		{"id":"field2","type":"string","ui_hint":"","value":"foo"}
	]
}
`

	req, err := http.NewRequest(http.MethodPost, "/test1", bytes.NewBufferString(requestBody))
	require.NoError(t, err)

	handler := setupTestFormHandler(t)
	handler.ServeHTTP(rec, req)
	require.Equal(t, http.StatusOK, rec.Code)

	body, err := ioutil.ReadAll(rec.Body)
	require.NoError(t, err)

	t.Log(string(body))
	// TODO
}

func TestSubmitCompleteForm(t *testing.T) {
	rec := httptest.NewRecorder()
	requestBody := `
{
	"name":"test1",
	"fields":[
		{"id":"field1","type":"string","ui_hint":"","value":"bar"},
		{"id":"field2","type":"string","ui_hint":"","value":"foo"}
	]
}
`

	req, err := http.NewRequest(http.MethodPost, "/test1", bytes.NewBufferString(requestBody))
	require.NoError(t, err)

	handler := setupTestFormHandler(t)
	handler.ServeHTTP(rec, req)
	require.Equal(t, http.StatusOK, rec.Code)

	body, err := ioutil.ReadAll(rec.Body)
	require.NoError(t, err)

	t.Log(string(body))
	// TODO
}