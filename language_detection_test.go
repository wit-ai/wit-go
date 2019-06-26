// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.

package witai

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLocales(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		_, _ = res.Write(
			[]byte(`{
				"detected_locales": [
					{
						"locale": "en_XX",
						"confidence": 1
					}
				]
			}`),
		)
	}))
	defer func() { testServer.Close() }()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL
	l, err := c.Detect("Hello world")

	if err != nil {
		t.Fatalf("nil error expected, got %v", err)
	}

	if l.DetectedLocales[0].Locale != "en_XX" {
		t.Fatalf("lang=en expected, got %v", l.DetectedLocales[0].Locale)
	}
}
