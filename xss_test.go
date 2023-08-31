package go_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateAutoEscape(w http.ResponseWriter, r *http.Request) {
	myTemplate.ExecuteTemplate(w, "post.gohtml", map[string]interface{}{
		"Title": "Hello World",
		"Body":  `<p class="test">Tag HTML yang akan di-escape</p>`,
	})
}

func TestTemplateAutoEscape(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscape(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateDisableAutoEscape(w http.ResponseWriter, r *http.Request) {
	myTemplate.ExecuteTemplate(w, "post.gohtml", map[string]interface{}{
		"Title": "Hello World",
		"Body":  template.HTML(`<p class="test">Tag HTML yang akan di-escape</p>`),
	})
}

func TestTemplateDisableAutoEscape(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDisableAutoEscape(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
