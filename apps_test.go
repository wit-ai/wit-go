// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.

package witai

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
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
		res.Write([]byte(`{"name": "alarm-clock","training_status":"done","will_train_at":"2018-07-29T18:17:34-0700","last_training_duration_secs":42,"last_trained_at":"2018-07-29T18:16:34-0700"}`))
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
	if app.TrainingStatus != Done {
		t.Fatalf("expected Done, got: %v", app)
	}

	expectedNextTrainTime, _ := time.Parse(WitTimeFormat, "2018-07-29T18:17:34-0700")
	if !app.WillTrainAt.Time.Equal(expectedNextTrainTime) {
		t.Fatalf("expected %v got: %v", app.WillTrainAt, expectedNextTrainTime)
	}

	if app.LastTrainingDurationSecs != 42 {
		t.Fatalf("Expected 42 got %v", app.LastTrainingDurationSecs)
	}

	expectedLastTrainTime, _ := time.Parse(WitTimeFormat, "2018-07-29T18:16:34-0700")
	if !app.LastTrainedAt.Time.Equal(expectedLastTrainTime) {
		t.Fatalf("expected %v got: %v", app.WillTrainAt, expectedLastTrainTime)
	}
}

func TestCreateApp(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`{"app_id": "ai-dee"}`))
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
	if app.AppID != "ai-dee" {
		t.Fatalf("id=ai-dee expected, got: %s", app.AppID)
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
