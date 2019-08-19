// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.

package witai

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetSamples(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`[{"text": "s1"}, {"text":"s2"}]`))
	}))
	defer func() { testServer.Close() }()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL
	samples, _ := c.GetSamples(1, 0)

	if len(samples) != 2 {
		t.Fatalf("expected 2 samples, got: %v", samples)
	}
}

func TestValidateSamples(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`{"sent": true, "n": 2}`))
	}))
	defer func() { testServer.Close() }()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL
	r, _ := c.ValidateSamples([]Sample{
		{
			Text: "hello",
		},
	})

	if r.N != 2 || !r.Sent {
		t.Fatalf("expected N=2 and Sent=true, got: %v", r)
	}
}

func TestDeleteSamples(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`{"sent": true, "n": 2}`))
	}))
	defer func() { testServer.Close() }()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL
	r, _ := c.DeleteSamples([]Sample{
		{
			Text: "hello",
		},
	})

	if r.N != 2 || !r.Sent {
		t.Fatalf("expected N=2 and Sent=true, got: %v", r)
	}
}
