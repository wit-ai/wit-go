// +build integration

package witai

import (
	"os"
	"testing"
)

var (
	integrationEntity = NewEntity{
		ID:  "integration_entity_id",
		Doc: "integration_entity_doc",
	}
	integrationEntityUpdateFields = UpdateEntityFields{
		Doc: "integration_entity_doc_updated",
	}
)

func TestIntegrationInvalidToken(t *testing.T) {
	c := NewClient("invalid_token")
	_, err := c.GetEntity(integrationEntity.ID)
	if err == nil {
		t.Fatalf("expected error, got: nil")
	}
}

func TestIntegrationEntities(t *testing.T) {
	c := getIntegrationClient()
	// just to make sure we don't create suplicates
	c.DeleteEntity(integrationEntity.ID)

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

	err = c.UpdateEntity(integrationEntity.ID, integrationEntityUpdateFields)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	e, err := c.GetEntity(integrationEntity.ID)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
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

func getIntegrationClient() *Client {
	return NewClient(os.Getenv("WITAI_INTEGRATION_TOKEN"))
}
