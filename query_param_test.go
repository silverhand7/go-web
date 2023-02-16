package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func SayHello(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(writer, "Hello")
	} else {
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Joe", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, string(body), "Hello Joe")
	fmt.Println(string(body))
}

func MultipleParameters(writer http.ResponseWriter, request *http.Request) {
	firstName := request.URL.Query().Get("first_name")
	lastName := request.URL.Query().Get("last_name")
	fmt.Fprintf(writer, "Hello %s %s", firstName, lastName)
}

func TestQueryParameters(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?first_name=Joe&last_name=Doe", nil)
	recorder := httptest.NewRecorder()

	MultipleParameters(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, string(body), "Hello Joe Doe")
	fmt.Println(string(body))
}

func MultipleParameterValues(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	names := query["name"]                        // akses map dari URL.Query()
	fmt.Fprint(writer, strings.Join(names, ", ")) // just like php implode but the last element will not have the connector (, )
}

func TestMultipleParameterValues(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Joe&name=Doe&name=Res", nil)
	recorder := httptest.NewRecorder()

	MultipleParameterValues(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, string(body), "Joe, Doe, Res")
	fmt.Println(string(body))
}
