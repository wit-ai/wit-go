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
		ID:  "integration_entity_id",
		Doc: "integration_entity_doc",
	}
	integrationApp = App{
		Name:        "integration_app_id",
		Private:     false,
		Description: "integration_app_desc",
		Lang:        "en",
	}
	integrationEntityUpdateFields = Entity{
		Name:    "integration_entity_id",
		Lookups: []string{"keywords"},
		Doc:     "integration_entity_doc_updated",
	}
)

func TestIntegrationInvalidToken(t *testing.T) {
	c := NewClient("invalid_token")
	_, err := c.GetEntity(integrationEntity.ID)
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
	c.DeleteEntity(integrationEntity.ID)

	// create entity
	entity, err := c.CreateEntity(integrationEntity)
	if err != nil {
		t.Fatalf("expected nil error got: %v", err)
	}
	if entity == nil {
		t.Fatalf("expected non nil entity")
	}
	if entity.Lang != "en" {
		t.Fatalf("expected lang=en, got: %s", entity.Lang)
	}

	// create may take some time
	time.Sleep(2 * time.Second)
}

func TestIntegrationUpdateEntity(t *testing.T) {
	c := getIntegrationClient()

	// update entity
	err := c.UpdateEntity(integrationEntity.ID, integrationEntityUpdateFields)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	time.Sleep(time.Second)
}

func TestIntegrationAddEntityValue(t *testing.T) {
	c := getIntegrationClient()

	var err error

	// add entity value 1
	if _, err = c.AddEntityValue(integrationEntity.ID, EntityValue{
		Value:       "London",
		Expressions: []string{"London"},
	}); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	// add entity value 2
	if _, err = c.AddEntityValue(integrationEntity.ID, EntityValue{
		Value:       "Ho Chi Minh City",
		Expressions: []string{"Ho Chi Minh City"},
	}); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	// add entity value expression
	if _, err = c.AddEntityValueExpression(integrationEntity.ID, "Ho Chi Minh City", "HoChiMinh"); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if _, err = c.AddEntityValueExpression(integrationEntity.ID, "Ho Chi Minh City", "hochiminhcity"); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if err = c.DeleteEntityValueExpression(integrationEntity.ID, "Ho Chi Minh City", "HoChiMinh"); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	// delete entity value 1
	if err = c.DeleteEntityValue(integrationEntity.ID, "Ho Chi Minh City"); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
}

func TestIntegrationGetEntity(t *testing.T) {
	c := getIntegrationClient()

	// check entity
	e, err := c.GetEntity(integrationEntity.ID)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if len(e.Values) != 1 {
		t.Fatalf("expected 1 value, got %v", e.Values)
	}

	if e.Doc != integrationEntityUpdateFields.Doc {
		t.Fatalf("expected doc=%s, got %s", integrationEntityUpdateFields.Doc, e.Doc)
	}

	entities, err := c.GetEntities()
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if len(entities) == 0 {
		t.Fatalf("expected >0 entities, got %v", entities)
	}

	err = c.DeleteEntity(integrationEntity.ID)
	if err != nil {
		t.Fatalf("expected nil error got err=%v", err)
	}
}

func TestIntegrationSamples(t *testing.T) {
	c := getIntegrationClient()

	// cleanup
	c.DeleteSamples([]Sample{
		{
			Text: "I want to fly SFO",
		},
	})

	// Deletion takes time
	time.Sleep(time.Second * 5)

	// samples test
	_, validateErr := c.ValidateSamples([]Sample{
		{
			Text: "I want to fly SFO",
			Entities: []SampleEntity{
				{
					Entity: "intent",
					Value:  "flight_request",
				},
				{
					Entity: "wit$location",
					Value:  "SFO",
					Start:  17,
					End:    20,
				},
			},
		},
	})
	if validateErr != nil {
		t.Fatalf("expected nil error, got %v", validateErr)
	}

	// Training takes time
	time.Sleep(time.Second * 20)

	// get samples
	samples, samplesErr := c.GetSamples(1, 0)
	if samplesErr != nil {
		t.Fatalf("expected nil error, got %v", samplesErr)
	}
	if len(samples) != 1 {
		t.Fatalf("expected 1 sample, got %v", samples)
	}

	// delete samples
	_, delSamplesErr := c.DeleteSamples([]Sample{
		{
			Text: "I want to fly SFO",
		},
	})
	if delSamplesErr != nil {
		t.Fatalf("expected nil error, got %v", delSamplesErr)
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
