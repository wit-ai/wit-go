package witai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// NewEntity - https://wit.ai/docs/http/20170307#post__entities_link
type NewEntity struct {
	ID  string `json:"id"`
	Doc string `json:"doc"`
}

// UpdateEntityFields - https://wit.ai/docs/http/20170307#put__entities__entity_id_link
type UpdateEntityFields struct {
	Doc     string        `json:"doc"`
	Lookups []string      `json:"lookups"`
	Values  []EntityValue `json:"values"`
}

// Entity - https://wit.ai/docs/http/20170307#post__entities_link
type Entity struct {
	ID      string        `json:"id"`
	Doc     string        `json:"doc"`
	Name    string        `json:"name"`
	Lang    string        `json:"lang"`
	Builtin bool          `json:"builtin"`
	Lookups []string      `json:"lookups"`
	Values  []EntityValue `json:"values"`
}

// EntityValue - https://wit.ai/docs/http/20170307#get__entities__entity_id_link
type EntityValue struct {
	Value       string   `json:"value"`
	Expressions []string `json:"expressions"`
	MetaData    string   `json:"metadata"`
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
	resp, err := c.request(http.MethodGet, fmt.Sprintf("/entities/%s", url.QueryEscape(id)), "application/json", nil)
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
	resp, err := c.request(http.MethodDelete, fmt.Sprintf("/entities/%s", url.QueryEscape(id)), "application/json", nil)
	if err == nil {
		resp.Close()
	}

	return err
}

// DeleteEntityRole - deletes entity role. https://wit.ai/docs/http/20170307#delete__entities__entity_id_role_id_link
func (c *Client) DeleteEntityRole(entityID string, role string) error {
	resp, err := c.request(http.MethodDelete, fmt.Sprintf("/entities/%s:%s", url.QueryEscape(entityID), url.QueryEscape(role)), "application/json", nil)
	if err == nil {
		resp.Close()
	}

	return err
}

// UpdateEntity - Updates an entity. https://wit.ai/docs/http/20170307#put__entities__entity_id_link
func (c *Client) UpdateEntity(id string, entity UpdateEntityFields) error {
	entityJSON, err := json.Marshal(entity)
	if err != nil {
		return err
	}

	resp, err := c.request(http.MethodPut, fmt.Sprintf("/entities/%s", url.QueryEscape(id)), "application/json", bytes.NewBuffer(entityJSON))
	if err != nil {
		return err
	}

	defer resp.Close()

	decoder := json.NewDecoder(resp)
	return decoder.Decode(&entity)
}

// AddEntityValue - Add a possible value into the list of values for the keyword entity. https://wit.ai/docs/http/20170307#post__entities__entity_id_values_link
func (c *Client) AddEntityValue(entityID string, value EntityValue) (*Entity, error) {
	valueJSON, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}

	resp, err := c.request(http.MethodPost, fmt.Sprintf("/entities/%s/values", url.QueryEscape(entityID)), "application/json", bytes.NewBuffer(valueJSON))
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

// DeleteEntityValue - Delete a canonical value from the entity. https://wit.ai/docs/http/20170307#delete__entities__entity_id_values_link
func (c *Client) DeleteEntityValue(entityID string, value string) error {
	resp, err := c.request(http.MethodDelete, fmt.Sprintf("/entities/%s/values/%s", url.QueryEscape(entityID), url.QueryEscape(value)), "application/json", nil)
	if err == nil {
		resp.Close()
	}

	return err
}

// AddEntityValueExpression - Create a new expression of the canonical value of the keyword entity. https://wit.ai/docs/http/20170307#post__entities__entity_id_values__value_id_expressions_link
func (c *Client) AddEntityValueExpression(entityID string, value string, expression string) (*Entity, error) {
	type expr struct {
		Expression string `json:"expression"`
	}

	exprJSON, err := json.Marshal(expr{
		Expression: expression,
	})
	if err != nil {
		return nil, err
	}

	resp, err := c.request(http.MethodPost, fmt.Sprintf("/entities/%s/values/%s/expressions", url.QueryEscape(entityID), url.QueryEscape(value)), "application/json", bytes.NewBuffer(exprJSON))
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

// DeleteEntityValueExpression - Delete an expression of the canonical value of the entity. https://wit.ai/docs/http/20170307#delete__entities__entity_id_values__value_id_expressions_link
func (c *Client) DeleteEntityValueExpression(entityID string, value string, expression string) error {
	resp, err := c.request(http.MethodDelete, fmt.Sprintf("/entities/%s/values/%s/expressions/%s", url.QueryEscape(entityID), url.QueryEscape(value), url.QueryEscape(expression)), "application/json", nil)
	if err == nil {
		resp.Close()
	}

	return err
}
