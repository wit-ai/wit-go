// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.

package witai

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

var unitTestToken = "unit_test_invalid_token"

func TestGetEntities(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(`[
			{
				"id": "2690212494559269",
				"name": "car"
			}, {
				"id": "254954985556896",
				"name": "color"
			}
		]`))
	}))
	defer func() { testServer.Close() }()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL
	entities, _ := c.GetEntities()

	wantEntities := []Entity{
		{ID: "2690212494559269", Name: "car"},
		{ID: "254954985556896", Name: "color"},
	}

	if !reflect.DeepEqual(entities, wantEntities) {
		t.Fatalf("expected\n\tentities: %v\n\tgot: %v", wantEntities, entities)
	}
}

func TestCreateEntity(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(`{
			"id": "5418abc7-cc68-4073-ae9e-3a5c3c81d965",
			"name": "favorite_city",
			"roles": [{"id": "1", "name": "favorite_city"}],
			"lookups": ["free-text", "keywords"],
			"keywords": []	
		}`))
	}))
	defer func() { testServer.Close() }()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL
	e, err := c.CreateEntity(Entity{
		Name:  "favorite_city",
		Roles: []string{},
	})

	wantEntity := &CreateEntityResponse{
		ID:       "5418abc7-cc68-4073-ae9e-3a5c3c81d965",
		Name:     "favorite_city",
		Roles:    []EntityRole{{ID: "1", Name: "favorite_city"}},
		Lookups:  []string{"free-text", "keywords"},
		Keywords: []EntityKeyword{},
	}

	if err != nil {
		t.Fatalf("nil error expected, got %v", err)
	}
	if !reflect.DeepEqual(wantEntity, e) {
		t.Fatalf("expected\n\tentity: %v\n\tgot: %v", wantEntity, e)
	}
}

func TestGetEntity(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(`{
			"id": "571979db-f6ac-4820-bc28-a1e0787b98fc",
			"name": "first_name",
			"lookups": ["free-text", "keywords"],
			"roles": [{"id": "1", "name": "first_name"}],
			"keywords": [ {
				"keyword": "Willy",
				"synonyms": ["Willy"]
			}, {
				"keyword": "Laurent",
				"synonyms": ["Laurent"]
			}, {
				"keyword": "Julien",
				"synonyms": ["Julien"]
			} ]
		}`))
	}))
	defer func() { testServer.Close() }()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL
	entity, err := c.GetEntity("first_name")

	wantEntity := &CreateEntityResponse{
		ID:      "571979db-f6ac-4820-bc28-a1e0787b98fc",
		Name:    "first_name",
		Roles:   []EntityRole{{ID: "1", Name: "first_name"}},
		Lookups: []string{"free-text", "keywords"},
		Keywords: []EntityKeyword{
			{Keyword: "Willy", Synonyms: []string{"Willy"}},
			{Keyword: "Laurent", Synonyms: []string{"Laurent"}},
			{Keyword: "Julien", Synonyms: []string{"Julien"}},
		},
	}

	if err != nil {
		t.Fatalf("nil error expected, got %v", err)
	}
	if !reflect.DeepEqual(wantEntity, entity) {
		t.Fatalf("expected\n\tentity: %v\n\tgot: %v", wantEntity, entity)
	}
}

func TestUpdateEntity(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`{
			"id": "5418abc7-cc68-4073-ae9e-3a5c3c81d965",
			"name": "favorite_city",
			"roles": [{"id": "1", "name": "favorite_city"}],
			"lookups": ["free-text", "keywords"],
			"keywords": [
				{
					"keyword": "Paris",
					"synonyms": ["Paris", "City of Light", "Capital of France"]
				},
				{
					"keyword": "Seoul",
					"synonyms": ["Seoul", "서울", "Kimchi paradise"]
				}
			]	
		}`))
	}))
	defer func() { testServer.Close() }()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL

	if _, err := c.UpdateEntity("favorite_city", Entity{
		Name:    "favorite_city",
		Roles:   []string{"favorite_city"},
		Lookups: []string{"free-text", "keywords"},
		Keywords: []EntityKeyword{
			{Keyword: "Paris", Synonyms: []string{"Paris", "City of Light", "Capital of France"}},
			{Keyword: "Seoul", Synonyms: []string{"Seoul", "서울", "Kimchi paradise"}},
		},
	}); err != nil {
		t.Fatalf("err=nil expected, got: %v", err)
	}
}

func TestDeleteEntity(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`{"deleted": "favorite_city"}`))
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
		res.Write([]byte(`{"deleted": "flight:destination"}`))
	}))
	defer func() { testServer.Close() }()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL
	if err := c.DeleteEntityRole("flight", "destination"); err != nil {
		t.Fatalf("expected nil error, got: %v", err)
	}
}

func TestAddEntityKeyword(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`{
			"id": "5418abc7-cc68-4073-ae9e-3a5c3c81d965",
			"name": "favorite_city",
			"roles": ["favorite_city"],
			"lookups": ["free-text", "keywords"],
			"keywords": [
				{
					"keyword": "Brussels",
					"synonyms": ["Brussels", "Capital of Belgium"]
				},
				{
					"keyword": "Paris",
					"synonyms": ["Paris", "City of Light", "Capital of France"]
				},
				{
					"keyword": "Seoul",
					"synonyms": ["Seoul", "서울", "Kimchi paradise"]
				}
			]	
		}`))
	}))
	defer func() { testServer.Close() }()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL
	entity, err := c.AddEntityKeyword("favorite_city", EntityKeyword{
		Keyword:  "Brussels",
		Synonyms: []string{"Brussels", "Capital of Belgium"},
	})

	wantEntity := &Entity{
		ID:      "5418abc7-cc68-4073-ae9e-3a5c3c81d965",
		Name:    "favorite_city",
		Roles:   []string{"favorite_city"},
		Lookups: []string{"free-text", "keywords"},
		Keywords: []EntityKeyword{
			{Keyword: "Brussels", Synonyms: []string{"Brussels", "Capital of Belgium"}},
			{Keyword: "Paris", Synonyms: []string{"Paris", "City of Light", "Capital of France"}},
			{Keyword: "Seoul", Synonyms: []string{"Seoul", "서울", "Kimchi paradise"}},
		},
	}

	if err != nil {
		t.Fatalf("nil error expected, got %v", err)
	}

	if !reflect.DeepEqual(wantEntity, entity) {
		t.Fatalf("expected\n\tentity: %v\n\tgot: %v", wantEntity, entity)
	}
}

func TestDeleteEntityKeyword(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`{"deleted": "Paris"}`))
	}))
	defer func() { testServer.Close() }()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL
	if err := c.DeleteEntityKeyword("favorite_city", "Paris"); err != nil {
		t.Fatalf("expected nil error, got: %v", err)
	}
}

func TestAddEntityKeywordSynonym(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`{
			"id": "5418abc7-cc68-4073-ae9e-3a5c3c81d965",
			"name": "favorite_city",
			"roles": ["favorite_city"],
			"lookups": ["free-text", "keywords"],
			"keywords": [
				{
					"keyword": "Brussels",
					"synonyms": ["Brussels", "Capital of Belgium"]
				},
				{
					"keyword": "Paris",
					"synonyms": ["Paris", "City of Light", "Capital of France", "Camembert city"]
				},
				{
					"keyword": "Seoul",
					"synonyms": ["Seoul", "서울", "Kimchi paradise"]
				}
			]	
		}`))
	}))
	defer func() { testServer.Close() }()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL
	entity, err := c.AddEntityKeywordSynonym("favorite_city", "Paris", "Camembert city")

	wantEntity := &Entity{
		ID:      "5418abc7-cc68-4073-ae9e-3a5c3c81d965",
		Name:    "favorite_city",
		Roles:   []string{"favorite_city"},
		Lookups: []string{"free-text", "keywords"},
		Keywords: []EntityKeyword{
			{Keyword: "Brussels", Synonyms: []string{"Brussels", "Capital of Belgium"}},
			{Keyword: "Paris", Synonyms: []string{"Paris", "City of Light", "Capital of France", "Camembert city"}},
			{Keyword: "Seoul", Synonyms: []string{"Seoul", "서울", "Kimchi paradise"}},
		},
	}

	if err != nil {
		t.Fatalf("nil error expected, got %v", err)
	}

	if !reflect.DeepEqual(wantEntity, entity) {
		t.Fatalf("expected\n\tentity: %v\n\tgot: %v", wantEntity, entity)
	}
}

func TestDeleteEntityKeywordSynonym(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`{"deleted": "Camembert City"}`))
	}))
	defer func() { testServer.Close() }()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL
	if err := c.DeleteEntityKeywordSynonym("favorite_city", "Paris", "Camembert City"); err != nil {
		t.Fatalf("expected nil error, got: %v", err)
	}
}
