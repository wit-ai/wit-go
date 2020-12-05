// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.

package witai

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestGetApps(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`[
			{
				"id": "9890809890",
				"name": "My_Second_App",
				"lang": "en",
				"private": false,
				"created_at": "2018-01-01T00:00:01Z"
			},
			{
				"id": "9890809891",
				"name": "My_Third_App",
				"lang": "en",
				"private": false,
				"created_at": "2018-01-02T00:00:01Z"
			}
		]`))
	}))
	defer func() { testServer.Close() }()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL
	apps, err := c.GetApps(2, 0)

	wantApps := []App{
		{
			ID:        "9890809890",
			Name:      "My_Second_App",
			Lang:      "en",
			Private:   false,
			CreatedAt: Time{time.Date(2018, 1, 1, 0, 0, 1, 0, time.UTC)},
		},
		{
			ID:        "9890809891",
			Name:      "My_Third_App",
			Lang:      "en",
			Private:   false,
			CreatedAt: Time{time.Date(2018, 1, 2, 0, 0, 1, 0, time.UTC)},
		},
	}

	if err != nil {
		t.Fatalf("nil error expected, got %v", err)
	}
	if !reflect.DeepEqual(wantApps, apps) {
		t.Fatalf("expected\n\tapps: %v\n\tgot: %v", wantApps, apps)
	}
}

func TestGetApp(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`{
			"id": "2802177596527671",
			"name": "alarm-clock",
			"lang": "en",
			"private": false,
			"created_at": "2018-07-29T18:15:34-0700",
			"last_training_duration_secs": 42,
			"will_train_at": "2018-07-29T18:17:34-0700",
			"last_trained_at": "2018-07-29T18:18:34-0700",
			"training_status": "done"
		}`))
	}))
	defer func() { testServer.Close() }()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL
	app, err := c.GetApp("2802177596527671")
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

	expectedLastTrainTime, _ := time.Parse(WitTimeFormat, "2018-07-29T18:18:34-0700")
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
		res.Write([]byte(`{"success": true}`))
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
		res.Write([]byte(`{"success": true}`))
	}))
	defer func() { testServer.Close() }()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL

	err := c.UpdateApp("appid", App{
		Lang:     "fr",
		Timezone: "Europe/Paris",
	})
	if err != nil {
		t.Fatalf("err=nil expected, got: %v", err)
	}
}

func TestGetAppTags(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`[
			[
				{
					"name": "v3",
					"created_at": "2019-09-14T20:29:53-0700",
					"updated_at": "2019-09-14T20:29:53-0700",
					"desc": "third version"
				},
				{
					"name": "v2",
					"created_at": "2019-08-08T11:05:35-0700",
					"updated_at": "2019-08-08T11:09:17-0700",
					"desc": "second version, moved to v3"
				}
			],
			[
				{
					"name": "v1",
					"created_at": "2019-08-08T11:02:52-0700",
					"updated_at": "2019-09-14T15:45:22-0700",
					"desc": "legacy first version"
				}
			]
		]`))
	}))
	defer testServer.Close()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL

	wantTags := [][]AppTag{
		{
			{Name: "v3", Desc: "third version"},
			{Name: "v2", Desc: "second version, moved to v3"},
		},
		{
			{Name: "v1", Desc: "legacy first version"},
		},
	}

	tags, err := c.GetAppTags("appid")
	if err != nil {
		t.Fatalf("expected nil err, got: %v", err)
	}

	if len(tags) != len(wantTags) {
		t.Fatalf("expected\n\ttags: %v\n\tgot: %v", wantTags, tags)
	}

	for i, group := range tags {
		wantGroup := wantTags[i]

		if len(group) != len(wantGroup) {
			t.Fatalf("expected\n\ttags[%v]: %v\n\tgot: %v", i, wantTags, tags)
		}

		for j, tag := range group {
			wantTag := wantGroup[j]
			if tag.Name != wantTag.Name {
				t.Fatalf("expected\n\ttags[%v][%v].Name: %v\n\tgot: %v", i, j, wantTag.Name, tag.Name)
			}
			if tag.Desc != wantTag.Desc {
				t.Fatalf("expected\n\ttags[%v][%v].Desc: %v\n\tgot: %v", i, j, wantTag.Desc, tag.Desc)
			}
		}
	}
}

func TestGetAppTag(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{
			"name": "v1",
			"created_at": "2019-08-08T11:02:52-0700",
			"updated_at": "2019-09-14T15:45:22-0700",
			"desc": "first version"
		}`))
	}))
	defer testServer.Close()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL

	tag, err := c.GetAppTag("appid", "tagid")
	if err != nil {
		t.Fatalf("expected nil err, got: %v", err)
	}

	wantTag := &AppTag{
		Name: "v1",
		Desc: "first version",
	}

	if tag.Name != wantTag.Name {
		t.Fatalf("expected\n\ttag.Name: %v, got: %v", wantTag.Name, tag.Name)
	}
	if tag.Desc != wantTag.Desc {
		t.Fatalf("expected\n\ttag.Desc: %v, got: %v", wantTag.Desc, tag.Desc)
	}
}

func TestCreateAppTag(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"tag": "v_1"}`))
	}))
	defer testServer.Close()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL

	tag, err := c.CreateAppTag("appid", "v$1")
	if err != nil {
		t.Fatalf("expected nil err, got: %v", err)
	}

	wantTag := &AppTag{Name: "v_1"}
	if tag.Name != wantTag.Name {
		t.Fatalf("expected\n\ttag.Name: %v, got: %v", wantTag.Name, tag.Name)
	}
}

func TestUpdateAppTag(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{
			"tag": "v1.0",
			"desc": "new description"
		}`))
	}))
	defer testServer.Close()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL

	updated, err := c.UpdateAppTag("appid", "v1", AppTag{Name: "v1.0", Desc: "new description"})
	if err != nil {
		t.Fatalf("expected nil err, got: %v", err)
	}

	wantUpdated := &AppTag{Name: "v1.0", Desc: "new description"}

	if !reflect.DeepEqual(updated, wantUpdated) {
		t.Fatalf("expected\n\tupdated: %v\n\tgot: %v", wantUpdated, updated)
	}
}

func TestMoveAppTag(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{
			"tag": "v1.0",
			"desc": "1.0 version, moved to 1.1 version",
			"moved_to": "v1.1"
		}`))
	}))
	defer testServer.Close()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL

	moved, err := c.MoveAppTag("appid", "v1", "v1.1", &AppTag{Name: "v1.0", Desc: "1.0 version, moved to 1.1 version"})
	if err != nil {
		t.Fatalf("expected nil err, got: %v", err)
	}

	wantMoved := &MovedAppTag{
		Tag:     "v1.0",
		Desc:    "1.0 version, moved to 1.1 version",
		MovedTo: "v1.1",
	}

	if !reflect.DeepEqual(moved, wantMoved) {
		t.Fatalf("expected\n\tmoved: %v\n\tgot: %v", wantMoved, moved)
	}
}

func TestDeleteAppTag(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"success": true}`))
	}))
	defer testServer.Close()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL

	err := c.DeleteAppTag("appid", "tagid")
	if err != nil {
		t.Fatalf("expected nil err, got: %v", err)
	}
}
