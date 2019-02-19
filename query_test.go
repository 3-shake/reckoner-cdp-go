package reckonercdp

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var dummyResponse = `{"history_id":103,"uuid":"37d88b12-856f-4e9c-8426-b867ca84cde1","has_next":true,
"records":[{"bar":"bar1","foo":"foo1"}, {"bar":"bar2","foo":"foo2"}]}`

func TestQuery(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(dummyResponse))
		},
	))

	defer ts.Close()
	client := NewClient(&ClientSettings{
		Host: ts.URL[7:len(ts.URL)],
	})

	res, err := client.Query("SELECT * FROM \\`reckoner-dev.load_test\\`.\\`load_test\\` LIMIT 2")
	if err != nil {
		t.Fatalf("Client#Query does not work.")
		return
	}
	log.Println(res)
}
