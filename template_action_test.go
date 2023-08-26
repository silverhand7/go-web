package go_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateIfElse(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./resources/if_else.html"))
	t.ExecuteTemplate(w, "if_else.html", map[string]interface{}{
		"Title": "Ada Title",
	})
}

func TestTemplateActionIfElse(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()
	TemplateIfElse(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateOperatorComparator(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./resources/comparator.html"))
	t.ExecuteTemplate(w, "comparator.html", map[string]interface{}{
		"Value": 11,
	})
}

func TestTemplateComparator(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()
	TemplateOperatorComparator(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateRangeLoop(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./resources/range_loop.html"))
	t.ExecuteTemplate(w, "range_loop.html", map[string]interface{}{
		"Title": "Hello World",
		"Hobbies": []string{
			"Game", "Code", "Cook", "Eat",
		},
	})
}

func TestTemplateLoopRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()
	TemplateRangeLoop(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateWith(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./resources/with.html"))
	t.ExecuteTemplate(w, "with.html", map[string]interface{}{
		"Title": "Hello World",
		"Name":  "Joe",
		"Hobbies": []string{
			"Game", "Code", "Cook", "Eat",
		},
		"Address": map[string]interface{}{
			"Street": "Jalan Denpasar",
			"City":   "Denpasar",
		},
	})
}

func TestTemplateWith(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()
	TemplateWith(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
