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

	// add entity value 1
	if _, err = c.AddEntityValue(integrationEntity.ID, EntityValue{
		Value:       "London",
		Expressions: []string{"London"},
	}); err != nil {
		t.Fatalf("expected non nil entity")
	}
	// add entity value 2
	if _, err = c.AddEntityValue(integrationEntity.ID, EntityValue{
		Value:       "HCMC",
		Expressions: []string{"Ho Chi Minh", "HCMC"},
	}); err != nil {
		t.Fatalf("expected non nil entity")
	}

	// delete entity 1
	if err = c.DeleteEntityValue(integrationEntity.ID, "HCMC"); err != nil {
		t.Fatalf("expected non nil entity")
	}

	err = c.UpdateEntity(integrationEntity.ID, integrationEntityUpdateFields)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

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

func getIntegrationClient() *Client {
	return NewClient(os.Getenv("WITAI_INTEGRATION_TOKEN"))
}
