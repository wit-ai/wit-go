package witai

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestExport(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`{"uri": "https://download"}`))
	}))
	defer func() { testServer.Close() }()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL
	uri, _ := c.Export()

	if uri != "https://download" {
		t.Fatalf("wrong download uri, got: %s", uri)
	}
}
