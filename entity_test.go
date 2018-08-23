package witai

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var unitTestToken = "unit_test_invalid_token"

func TestGetEntities(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`["e1", "e2"]`))
	}))
	defer func() { testServer.Close() }()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL
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

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL
	e, err := c.CreateEntity(NewEntity{
		ID:  "favorite_city",
		Doc: "A city that I like",
	})

	if err != nil || e.Lang != "en" {
		t.Fatalf("lang=en expected, got: %s", e.Lang)
	}
}

func TestGetEntity(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`{"doc": "My favorite city"}`))
	}))
	defer func() { testServer.Close() }()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL
	entity, _ := c.GetEntity("favorite_city")

	if entity == nil || entity.Doc != "My favorite city" {
		t.Fatalf("expected valid entity, got: %v", entity)
	}
}

func TestDeleteEntity(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`{}`))
	}))
	defer func() { testServer.Close() }()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL
	err := c.DeleteEntity("favorite_city")

	if err != nil {
		t.Fatalf("expected nil error, got: %v", err)
	}
}

func TestDeleteEntityRole(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`{}`))
	}))
	defer func() { testServer.Close() }()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL
	err := c.DeleteEntityRole("favorite_city", "role")

	if err != nil {
		t.Fatalf("expected nil error, got: %v", err)
	}
}

func TestUpdateEntity(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`{"doc": "new doc"}`))
	}))
	defer func() { testServer.Close() }()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL

	err := c.UpdateEntity("favorite_city", UpdateEntityFields{
		Doc: "new doc",
	})

	if err != nil {
		t.Fatalf("err=nil expected, got: %v", err)
	}
}
