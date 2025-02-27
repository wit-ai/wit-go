//go:build integration
// +build integration

// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.

package witai

import (
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
)

var (
	integrationEntity = Entity{
		Name:  "integration_entity_name",
		Roles: []string{"favorite_city"},
	}
	integrationApp = App{
		Name:    "integration_app_id",
		Private: false,
	}
	integrationEntityUpdateFields = Entity{
		Name:    "integration_entity_name",
		Roles:   []string{"favorite_city"},
		Lookups: []string{"keywords"},
	}
)

func TestIntegrationInvalidToken(t *testing.T) {
	c := NewClient("invalid_token")
	_, err := c.GetEntity(integrationEntity.Name)
	if err == nil {
		t.Fatalf("expected error, got: nil")
	}
}

func TestIntegrationGetUnknownEntity(t *testing.T) {
	c := getIntegrationClient()
	_, err := c.GetEntity("unknown_id")
	if err == nil {
		t.Fatalf("expected error, got: nil")
	}
}

func TestIntegrationDeleteUnknownEntity(t *testing.T) {
	c := getIntegrationClient()
	err := c.DeleteEntity("unknown_id")
	if err == nil {
		t.Fatalf("expected error, got: nil")
	}
}

func TestIntegrationUnknownEntity(t *testing.T) {
	c := getIntegrationClient()
	_, err := c.GetEntity("unknown_id")
	if err == nil {
		t.Fatalf("expected error, got: nil")
	}
}

func TestIntegrationCreateEntity(t *testing.T) {
	c := getIntegrationClient()

	// just to make sure we don't create duplicates
	c.DeleteEntity(integrationEntity.Name)

	// delete may take some time
	time.Sleep(2 * time.Second)

	// create entity
	entity, err := c.CreateEntity(integrationEntity)
	if err != nil {
		t.Fatalf("expected nil error got: %v", err)
	}
	if entity == nil {
		t.Fatalf("expected non nil entity")
	}
	if entity.Name != integrationEntity.Name {
		t.Fatalf("expected entity name %s, got %s", integrationEntity.Name, entity.Name)
	}

	// create may take some time
	time.Sleep(2 * time.Second)
}

func TestIntegrationUpdateEntity(t *testing.T) {
	c := getIntegrationClient()

	// update entity
	entity, err := c.UpdateEntity(integrationEntity.Name, integrationEntityUpdateFields)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if entity == nil {
		t.Fatalf("expected non nil entity")
	}
	if entity.Name != integrationEntity.Name {
		t.Fatalf("expected entity name %s, got %s", integrationEntity.Name, entity.Name)
	}

	time.Sleep(time.Second)
}

func TestIntegrationGetEntity(t *testing.T) {
	c := getIntegrationClient()

	e, err := c.GetEntity(integrationEntity.Name)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if e.Name != integrationEntity.Name {
		t.Fatalf("expected entity name %s, got %s", integrationEntity.Name, e.Name)
	}

	entities, err := c.GetEntities()
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if len(entities) == 0 {
		t.Fatalf("expected >0 entities, got %v", entities)
	}

	err = c.DeleteEntity(integrationEntity.Name)
	if err != nil {
		t.Fatalf("expected nil error got err=%v", err)
	}
}

func TestIntegrationExport(t *testing.T) {
	c := getIntegrationClient()

	uri, _ := c.Export()

	if !strings.Contains(uri, "fbcdn.net") {
		t.Fatalf("uri should contain fbcdn.net, got: %s", uri)
	}
}

func getIntegrationClient() *Client {
	c := NewClient(os.Getenv("WITAI_INTEGRATION_TOKEN"))
	c.SetHTTPClient(&http.Client{
		Timeout: time.Second * 20,
	})
	return c
}

func TestIntegrationDictation(t *testing.T) {
	c := getIntegrationClient()

	f, err := os.Open("./testdata/test.mp3")
	if err != nil {
		t.Fatalf("unable to open test file, err: %v", err)
	}
	defer f.Close()

	req := DictationRequest{
		File:        f,
		ContentType: "audio/mpeg3",
	}

	resp, err := c.Dictation(req)
	if err != nil {
		t.Fatalf("unexpected err, got %v", err)
	}

	if resp.Text != "Hi this is a test file" {
		t.Fatalf("unexpected transcript, got %s", resp.Text)
	}
}
