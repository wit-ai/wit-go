package witai

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetEntities(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`["e1", "e2"]`))
	}))
	defer func() { testServer.Close() }()

	apiBase = testServer.URL

	c := NewClient("token")
	entities, _ := c.GetEntities()

	if len(entities) != 2 {
		t.Fatalf("expected 2 entities, got: %d", len(entities))
	}
}

func TestCreateEntity(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`{"lang": "en"}`))
	}))
	defer func() { testServer.Close() }()

	apiBase = testServer.URL

	c := NewClient("token")
	e, err := c.CreateEntity(NewEntity{
		ID:  "favorite_city",
		Doc: "A city that I like",
	})

	if err != nil || e.Lang != "en" {
		t.Fatalf("lang=en expected, got: %s", e.Lang)
	}
}
