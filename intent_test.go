package witai

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetIntents(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`[
			{
				"id": "2690212494559269",
				"name": "buy_car"
			},
			{
				"id": "254954985556896",
				"name": "get_weather"
			},
			{
				"id": "233273197778131",
				"name": "make_call"
			}
		]`))
	}))
	defer testServer.Close()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL
	intents, _ := c.GetIntents()

	wantIntents := []Intent{
		{ID: "2690212494559269", Name: "buy_car"},
		{ID: "254954985556896", Name: "get_weather"},
		{ID: "233273197778131", Name: "make_call"},
	}

	if !reflect.DeepEqual(intents, wantIntents) {
		t.Fatalf("expected\n\tintents: %v\n\tgot: %v", wantIntents, intents)
	}
}

func TestCreateIntent(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{
			"id": "13989798788",
			"name": "buy_flowers"
		}`))
	}))
	defer testServer.Close()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL
	intent, err := c.CreateIntent("buy_flowers")

	wantIntent := &Intent{
		ID:   "13989798788",
		Name: "buy_flowers",
	}

	if err != nil {
		t.Fatalf("nil error expected, got %v", err)
	}
	if !reflect.DeepEqual(wantIntent, intent) {
		t.Fatalf("expected\n\tentity: %v\n\tgot: %v", wantIntent, intent)
	}
}

func TestGetIntent(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{
			"id": "13989798788",
			"name": "buy_flowers",
			"entities": [{
				"id": "9078938883",
				"name": "flower:flower"
			},{
				"id": "11223229984",
				"name": "wit$contact:contact"
			}]
		}`))
	}))
	defer testServer.Close()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL
	intent, err := c.GetIntent("buy_flowers")

	wantIntent := &Intent{
		ID:   "13989798788",
		Name: "buy_flowers",
		Entities: []Entity{
			{ID: "9078938883", Name: "flower:flower"},
			{ID: "11223229984", Name: "wit$contact:contact"},
		},
	}

	if err != nil {
		t.Fatalf("nil error expected, got %v", err)
	}
	if !reflect.DeepEqual(wantIntent, intent) {
		t.Fatalf("expected\n\tentity: %v\n\tgot: %v", wantIntent, intent)
	}
}

func TestDeleteIntent(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"deleted": "buy_flowers"}`))
	}))
	defer testServer.Close()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL
	if err := c.DeleteIntent("buy_flowers"); err != nil {
		t.Fatalf("nil error expected, got %v", err)
	}
}
