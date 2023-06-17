package http

import (
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestBadVar(t *testing.T) {
	get = func(url string) (*http.Response, error) {
		json := `{"id": 1, "name": "AnuchitO", "description": "gopher"}`
		return &http.Response{
			Body: io.NopCloser(strings.NewReader(json)),
		}, nil
	}

	resp, err := MakeHTTPCall("http://localhost:3333")
	if err != nil {
		t.Fatal(err)
	}

	want := &Response{
		ID:          1,
		Name:        "AnuchitO",
		Description: "gopher",
	}

	if !reflect.DeepEqual(resp, want) {
		t.Errorf("expected (%v), got (%v)", want, resp)
	}
}
