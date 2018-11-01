// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.

package witai

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetApps(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`[{"name": "app1"}, {"name":"app2"}]`))
	}))
	defer func() { testServer.Close() }()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL
	apps, _ := c.GetApps(1, 0)

	if len(apps) != 2 {
		t.Fatalf("expected 2 apps, got: %v", apps)
	}
}

func TestGetApp(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`{"name": "alarm-clock"}`))
	}))
	defer func() { testServer.Close() }()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL
	app, err := c.GetApp("my-id")
	if err != nil {
		t.Fatalf("not expected err, got: %s", err.Error())
	}

	if app.Name != "alarm-clock" {
		t.Fatalf("expected alarm-clock, got: %v", app)
	}
}

func TestCreateApp(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`{"id": "ai-dee"}`))
	}))
	defer func() { testServer.Close() }()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL
	app, err := c.CreateApp(App{
		Name: "app",
	})
	if err != nil {
		t.Fatalf("nil error expected, got %v", err)
	}
	if app.ID != "ai-dee" {
		t.Fatalf("id=ai-dee expected, got: %s", app.ID)
	}
}

func TestDeleteApp(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`{}`))
	}))
	defer func() { testServer.Close() }()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL
	if err := c.DeleteApp("appid"); err != nil {
		t.Fatalf("expected nil error, got: %v", err)
	}
}

func TestUpdateApp(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`{"description": "updated"}`))
	}))
	defer func() { testServer.Close() }()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL

	app, err := c.UpdateApp("appid", App{
		Description: "new desc",
	})
	if err != nil {
		t.Fatalf("err=nil expected, got: %v", err)
	}
	if app.Description != "updated" {
		t.Fatalf("description=updated expected, got: %s", app.Lang)
	}
}
