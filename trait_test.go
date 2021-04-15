// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.

package witai

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetTraits(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(`[
			{
				"id": "2690212494559269",
				"name": "wit$sentiment"
			},
			{
				"id": "254954985556896",
				"name": "faq"
			}
		]`))
	}))
	defer testServer.Close()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL
	traits, _ := c.GetTraits()

	wantTraits := []Trait{
		{ID: "2690212494559269", Name: "wit$sentiment"},
		{ID: "254954985556896", Name: "faq"},
	}

	if !reflect.DeepEqual(traits, wantTraits) {
		t.Fatalf("expected\n\ttraits: %v\n\tgot: %v", wantTraits, traits)
	}
}

func TestCreateTrait(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(`{
			"id": "13989798788",
			"name": "politeness",
			"values": [
				{
					"id": "97873388",
					"value": "polite"
				},
				{
					"id": "54493392772",
					"value": "rude"
				} 
			]
		}`))
	}))
	defer testServer.Close()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL
	trait, err := c.CreateTrait("politeness", []string{"polite", "rude"})

	wantTrait := &Trait{
		ID:   "13989798788",
		Name: "politeness",
		Values: []TraitValue{
			{ID: "97873388", Value: "polite"},
			{ID: "54493392772", Value: "rude"},
		},
	}

	if err != nil {
		t.Fatalf("nil error expected, got %v", err)
	}
	if !reflect.DeepEqual(wantTrait, trait) {
		t.Fatalf("expected\n\ttrait: %v\n\tgot: %v", wantTrait, trait)
	}
}

func TestGetTrait(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(`{
			"id": "13989798788",
			"name": "politeness",
			"values": [
				{
					"id": "97873388",
					"value": "polite"
				},
				{
					"id": "54493392772",
					"value": "rude"
				} 
			]
		}`))
	}))
	defer testServer.Close()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL
	trait, err := c.GetTrait("politeness")

	wantTrait := &Trait{
		ID:   "13989798788",
		Name: "politeness",
		Values: []TraitValue{
			{ID: "97873388", Value: "polite"},
			{ID: "54493392772", Value: "rude"},
		},
	}

	if err != nil {
		t.Fatalf("nil error expected, got %v", err)
	}
	if !reflect.DeepEqual(wantTrait, trait) {
		t.Fatalf("expected\n\ttrait: %v\n\tgot: %v", wantTrait, trait)
	}
}

func TestDeleteTrait(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`{"deleted": "politeness"}`))
	}))
	defer func() { testServer.Close() }()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL
	if err := c.DeleteTrait("politeness"); err != nil {
		t.Fatalf("expected nil error, got: %v", err)
	}
}

func TestAddTraitValue(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(`{
			"id": "13989798788",
			"name": "politeness",
			"values": [
				{
					"id": "97873388",
					"value": "polite"
				},
				{
					"id": "54493392772",
					"value": "rude"
				},
				{
					"id": "828283932",
					"value": "neutral"
				}
			]
		}`))
	}))
	defer func() { testServer.Close() }()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL
	trait, err := c.AddTraitValue("politeness", "neutral")

	wantTrait := &Trait{
		ID:   "13989798788",
		Name: "politeness",
		Values: []TraitValue{
			{ID: "97873388", Value: "polite"},
			{ID: "54493392772", Value: "rude"},
			{ID: "828283932", Value: "neutral"},
		},
	}

	if err != nil {
		t.Fatalf("nil error expected, got %v", err)
	}

	if !reflect.DeepEqual(wantTrait, trait) {
		t.Fatalf("expected\n\ttrait: %v\n\tgot: %v", wantTrait, trait)
	}
}

func TestDeleteTraitValue(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`{"deleted": "neutral"}`))
	}))
	defer func() { testServer.Close() }()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL
	if err := c.DeleteTraitValue("politeness", "neutral"); err != nil {
		t.Fatalf("expected nil error, got: %v", err)
	}
}
