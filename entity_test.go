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
	if err != nil {
		t.Fatalf("nil error expected, got %v", err)
	}
	if e.Lang != "en" {
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
	entity, err := c.GetEntity("favorite_city")
	if err != nil {
		t.Fatalf("nil error expected, got %v", err)
	}
	if entity.Doc != "My favorite city" {
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
	if err := c.DeleteEntity("favorite_city"); err != nil {
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
	if err := c.DeleteEntityRole("favorite_city", "role"); err != nil {
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

	if err := c.UpdateEntity("favorite_city", UpdateEntityFields{
		Doc: "new doc",
	}); err != nil {
		t.Fatalf("err=nil expected, got: %v", err)
	}
}

func TestAddEntityValue(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`{"lang": "de"}`))
	}))
	defer func() { testServer.Close() }()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL
	e, err := c.AddEntityValue("favorite_city", EntityValue{
		Value: "Minsk",
	})
	if err != nil {
		t.Fatalf("nil error expected, got %v", err)
	}
	if e.Lang != "de" {
		t.Fatalf("lang=de expected, got: %s", e.Lang)
	}
}

func TestDeleteEntityValue(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`{}`))
	}))
	defer func() { testServer.Close() }()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL
	if err := c.DeleteEntityValue("favorite_city", "London"); err != nil {
		t.Fatalf("expected nil error, got: %v", err)
	}
}

func TestAddEntityValueExpression(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`{"lang": "de"}`))
	}))
	defer func() { testServer.Close() }()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL
	e, err := c.AddEntityValueExpression("favorite_city", "Minsk", "Minsk")
	if err != nil {
		t.Fatalf("nil error expected, got %v", err)
	}
	if e.Lang != "de" {
		t.Fatalf("lang=de expected, got: %s", e.Lang)
	}
}

func TestDeleteEntityValueExpression(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`{}`))
	}))
	defer func() { testServer.Close() }()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL
	if err := c.DeleteEntityValueExpression("favorite_city", "Minsk", "Minsk"); err != nil {
		t.Fatalf("expected nil error, got: %v", err)
	}
}
