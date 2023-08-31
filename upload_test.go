package go_web

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func UploadForm(writer http.ResponseWriter, request *http.Request) {
	htmlTemplate.ExecuteTemplate(writer, "upload.html", nil)
}

func UploadProcess(writer http.ResponseWriter, request *http.Request) {
	// request.ParseMultipartForm(32 << 20) set default file size (by default 32mb)

	// ambil file dari form
	file, fileHeader, err := request.FormFile("file") // file name
	if err != nil {
		panic(err)
	}

	// buat destinasi/tempat file tersimpan
	fileDestination, err := os.Create("./resources/upload/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}

	// simpan file tersebut ke destinasi yang sudah diset diatas
	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err)
	}

	name := request.PostFormValue("name")
	htmlTemplate.ExecuteTemplate(writer, "upload.success.html", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})
}

func TestUploadFormServer(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/form", UploadForm)
	mux.HandleFunc("/upload", UploadProcess)

	// uploaded file dari folder resources/upload akan diakses melalui /static/
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources/upload"))))
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources/upload/icon.png
var uploadFileTest []byte

func TestUploadFile(t *testing.T) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("name", "Name value from field")
	file, _ := writer.CreateFormFile("file", "contoh-upload.png")
	file.Write(uploadFileTest)
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/upload", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	UploadProcess(recorder, request)

	bodyResponse, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(bodyResponse))
}
