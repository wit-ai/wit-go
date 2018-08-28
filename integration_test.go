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
	integrationApp = App{
		Name:        "integration_app_id",
		Private:     false,
		Description: "integration_app_desc",
		Lang:        "en",
	}
	integrationEntityUpdateFields = UpdateEntityFields{
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

func TestIntegrationApps(t *testing.T) {
	c := getIntegrationClient()

	// delete app
	apps, err := c.GetApps(10, 0)
	for _, a := range apps {
		if a.Name == integrationApp.Name {
			c.DeleteApp(a.ID)
		}
	}

	app, err := c.CreateApp(integrationApp)
	if err != nil {
		t.Fatalf("not expected error, got %v", err)
	}

	getApp, err := c.GetApp(app.AppID)
	if err != nil {
		t.Fatalf("not expected error, got %v", err)
	}
	if getApp.Name != integrationApp.Name {
		t.Fatalf("expected app name %s, got %s", integrationApp.Name, getApp.Name)
	}

	err = c.DeleteApp(app.AppID)
	if err != nil {
		t.Fatalf("not expected error, got %v", err)
	}
}

func TestIntegrationFullFlow(t *testing.T) {
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

	// update entity
	err = c.UpdateEntity(integrationEntity.ID, integrationEntityUpdateFields)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	// add entity value 1
	if _, err = c.AddEntityValue(integrationEntity.ID, EntityValue{
		Value:       "London",
		Expressions: []string{"London"},
	}); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	// add entity value 2
	if _, err = c.AddEntityValue(integrationEntity.ID, EntityValue{
		Value:       "HCMC",
		Expressions: []string{"HCMC"},
	}); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	// add entity value expression
	if _, err = c.AddEntityValueExpression(integrationEntity.ID, "HCMC", "HoChiMinh"); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if _, err = c.AddEntityValueExpression(integrationEntity.ID, "HCMC", "hochiminhcity"); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if err = c.DeleteEntityValueExpression(integrationEntity.ID, "HCMC", "HoChiMinh"); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	// delete entity 1
	if err = c.DeleteEntityValue(integrationEntity.ID, "HCMC"); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

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

func getIntegrationClient() *Client {
	return NewClient(os.Getenv("WITAI_INTEGRATION_TOKEN"))
}
