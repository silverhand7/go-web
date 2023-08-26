package go_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateLayout(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"./resources/layouts/header.html",
		"./resources/layouts/footer.html",
		"./resources/layouts/main.html",
	))
	t.ExecuteTemplate(w, "main", map[string]interface{}{
		"Title": "Template Layouts",
		"Name":  "Joe",
	})
}

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()
	TemplateLayout(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))

	// This is how to render the template as a real website.
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	t := template.Must(template.ParseFiles(
	// 		"./resources/layouts/header.html",
	// 		"./resources/layouts/footer.html",
	// 		"./resources/layouts/main.html",
	// 	))
	// 	t.ExecuteTemplate(w, "main.html", map[string]interface{}{
	// 		"Title": "Template Layouts",
	// 		"Name":  "Joe",
	// 	})
	// })

	// http.ListenAndServe(":8080", nil)
}
