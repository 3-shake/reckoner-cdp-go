package reckonercdp

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type Item struct {
	bar string
	foo string
}

func TestInsertError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusUnauthorized)
		},
	))
	defer ts.Close()

	client := NewClient(&ClientSettings{
		StreamingHost: ts.URL[7:len(ts.URL)],
	})
	if err := client.Insert(
		&Item{bar: "bar", foo: "foo"},
		"test_db",
		"test_table"); err == nil {
		t.Fatalf("Client#Insert raise error when response status is not 200.")
	}
}

func TestBulkInsertError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusUnauthorized)
		},
	))
	defer ts.Close()

	client := NewClient(&ClientSettings{
		StreamingHost: ts.URL[7:len(ts.URL)],
	})
	if err := client.BulkInsert(
		[]interface{}{&Item{bar: "bar", foo: "foo"}},
		"test_db",
		"test_table"); err == nil {
		t.Fatalf("Client#BulkInsert raise error when response status is not 200.")
	}
}
