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
	if err != nil || entity == nil || entity.Lang != "en" {
		t.Fatalf("expected nil error and lang=en, got: %v, err=%v", entity, err)
	}

	entities, err := c.GetEntities()
	if err != nil || len(entities) == 0 {
		t.Fatalf("expected >0 entities, got %v, err=%v", entities, err)
	}

	err = c.DeleteEntity(integrationEntity.ID)
	if err != nil {
		t.Fatalf("expected nil error got err=%v", err)
	}
}

func getIntegrationClient() *Client {
	return NewClient(os.Getenv("WITAI_INTEGRATION_TOKEN"))
}
