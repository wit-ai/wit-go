package witai

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// NewEntity - https://wit.ai/docs/http/20170307#post__entities_link
type NewEntity struct {
	ID  string `json:"id"`
	Doc string `json:"doc"`
}

// Entity - https://wit.ai/docs/http/20170307#post__entities_link
type Entity struct {
	ID      string `json:"id"`
	Doc     string `json:"doc"`
	Name    string `json:"name"`
	Lang    string `json:"lang"`
	BuiltIn bool   `json:"builtin"`
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
func (c *Client) CreateEntity(e NewEntity) (*Entity, error) {
	entityJSON, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}

	resp, err := c.request(http.MethodPost, "/entities", "application/json", bytes.NewBuffer(entityJSON))
	if err != nil {
		return nil, err
	}

	defer resp.Close()

	var entity *Entity
	decoder := json.NewDecoder(resp)
	err = decoder.Decode(&entity)
	if err != nil {
		return nil, err
	}

	return entity, nil
}
