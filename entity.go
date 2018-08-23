package witai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// NewEntity - https://wit.ai/docs/http/20170307#post__entities_link
type NewEntity struct {
	ID  string `json:"id"`
	Doc string `json:"doc"`
}

// UpdateEntityFields - https://wit.ai/docs/http/20170307#put__entities__entity_id_link
type UpdateEntityFields struct {
	Doc     string   `json:"doc"`
	Lookups []string `json:"lookups"`
	Values  []Value  `json:"values"`
}

// Entity - https://wit.ai/docs/http/20170307#post__entities_link
type Entity struct {
	ID      string   `json:"id"`
	Doc     string   `json:"doc"`
	Name    string   `json:"name"`
	Lang    string   `json:"lang"`
	Builtin bool     `json:"builtin"`
	Lookups []string `json:"lookups"`
	Values  []Value  `json:"values"`
}

// Value - https://wit.ai/docs/http/20170307#get__entities__entity_id_link
type Value struct {
	Value       string   `json:"value"`
	Expressions []string `json:"expressions"`
}

// GetEntities - returns list of entities. https://wit.ai/docs/http/20170307#get__entities_link
func (c *Client) GetEntities() ([]string, error) {
	resp, err := c.request(http.MethodGet, "/entities", "application/json", nil)
	if err != nil {
		return []string{}, err
	}

	defer resp.Close()

	var entities []string
	decoder := json.NewDecoder(resp)
	err = decoder.Decode(&entities)
	return entities, err
}

// CreateEntity - Creates a new entity with the given attributes. https://wit.ai/docs/http/20170307#post__entities_link
func (c *Client) CreateEntity(entity NewEntity) (*Entity, error) {
	entityJSON, err := json.Marshal(entity)
	if err != nil {
		return nil, err
	}

	resp, err := c.request(http.MethodPost, "/entities", "application/json", bytes.NewBuffer(entityJSON))
	if err != nil {
		return nil, err
	}

	defer resp.Close()

	var entityResp *Entity
	decoder := json.NewDecoder(resp)
	if err = decoder.Decode(&entityResp); err != nil {
		return nil, err
	}

	return entityResp, nil
}

// GetEntity - returns entity by ID. https://wit.ai/docs/http/20170307#get__entities__entity_id_link
func (c *Client) GetEntity(id string) (*Entity, error) {
	resp, err := c.request(http.MethodGet, fmt.Sprintf("/entities/%s", id), "application/json", nil)
	if err != nil {
		return nil, err
	}

	defer resp.Close()

	var entity *Entity
	decoder := json.NewDecoder(resp)
	if err = decoder.Decode(&entity); err != nil {
		return nil, err
	}

	return entity, nil
}

// DeleteEntity - deletes entity by ID. https://wit.ai/docs/http/20170307#delete__entities__entity_id_link
func (c *Client) DeleteEntity(id string) error {
	resp, err := c.request(http.MethodDelete, fmt.Sprintf("/entities/%s", id), "application/json", nil)
	if err != nil {
		return err
	}

	defer resp.Close()

	return nil
}

// DeleteEntityRole - deletes entity role. https://wit.ai/docs/http/20170307#delete__entities__entity_id_role_id_link
func (c *Client) DeleteEntityRole(entityID string, role string) error {
	resp, err := c.request(http.MethodDelete, fmt.Sprintf("/entities/%s:%s", entityID, role), "application/json", nil)
	if err != nil {
		return err
	}

	defer resp.Close()

	return nil
}

// UpdateEntity - Updates an entity. https://wit.ai/docs/http/20170307#put__entities__entity_id_link
func (c *Client) UpdateEntity(id string, entity UpdateEntityFields) error {
	entityJSON, err := json.Marshal(entity)
	if err != nil {
		return err
	}

	resp, err := c.request(http.MethodPut, "/entities/"+id, "application/json", bytes.NewBuffer(entityJSON))
	if err != nil {
		return err
	}

	defer resp.Close()

	decoder := json.NewDecoder(resp)
	return decoder.Decode(&entity)
}
