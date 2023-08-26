package go_web

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed resources/*.gohtml
var webTemplate embed.FS

var myTemplate = template.Must(template.ParseFS(webTemplate, "resources/*.gohtml"))

func TemplateCaching(w http.ResponseWriter, r *http.Request) {
	myTemplate.ExecuteTemplate(w, "simple.gohtml", "Hello template caching")
}

func TestTemplateCaching(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()
	TemplateCaching(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
