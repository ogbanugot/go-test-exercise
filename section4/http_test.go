// Section 4: Testing

// HTTP Testing for HTTP Server
package section4

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"go-test-exercise/section3"
)

func TestHelloWorldHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(section3.HelloWorldHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Hello, World!\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}