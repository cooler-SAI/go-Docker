package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8080/testpath", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	if status := w.Result().StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Hello, you've hit /testpath\n"
	if body := w.Body.String(); body != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", body, expected)
	}
}

func TestMainFunction(t *testing.T) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	go func() {
		main()
	}()

	time.Sleep(1 * time.Second)

	resp, err := http.Get("http://localhost:8080/")
	if err != nil {
		t.Fatalf("Could not start server: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Unexpected status code: got %v, want %v", resp.StatusCode, http.StatusOK)
	}
}
