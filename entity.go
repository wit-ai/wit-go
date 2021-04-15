// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.

package witai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Entity represents a wit-ai Entity.
//
// https://wit.ai/docs/http/20200513/#post__entities_link
type Entity struct {
	ID       string          `json:"id"`
	Name     string          `json:"name"`
	Lookups  []string        `json:"lookups,omitempty"`
	Roles    []string        `json:"roles,omitempty"`
	Keywords []EntityKeyword `json:"keywords,omitempty"`
}

// EntityKeyword is a keyword lookup for an Entity.
//
// https://wit.ai/docs/http/20200513/#post__entities__entity_keywords_link
type EntityKeyword struct {
	Keyword  string   `json:"keyword"`
	Synonyms []string `json:"synonyms"`
}

// GetEntities - returns list of entities.
//
// https://wit.ai/docs/http/20200513/#get__entities_link
func (c *Client) GetEntities() ([]Entity, error) {
	resp, err := c.request(http.MethodGet, "/entities", "application/json", nil)
	if err != nil {
		return []Entity{}, err
	}

	defer resp.Close()

	var entities []Entity
	decoder := json.NewDecoder(resp)
	err = decoder.Decode(&entities)
	return entities, err
}

// CreateEntity - Creates a new entity with the given attributes
//
// https://wit.ai/docs/http/20200513/#post__entities_link
func (c *Client) CreateEntity(entity Entity) (*Entity, error) {
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
	err = decoder.Decode(&entityResp)
	return entityResp, err
}

// GetEntity - returns entity by ID or name.
//
// https://wit.ai/docs/http/20200513/#get__entities__entity_link
func (c *Client) GetEntity(entityID string) (*Entity, error) {
	resp, err := c.request(http.MethodGet, fmt.Sprintf("/entities/%s", url.PathEscape(entityID)), "application/json", nil)
	if err != nil {
		return nil, err
	}

	defer resp.Close()

	var entity *Entity
	decoder := json.NewDecoder(resp)
	err = decoder.Decode(&entity)
	return entity, err
}

// UpdateEntity - Updates an entity by ID or name.
//
// https://wit.ai/docs/http/20200513/#put__entities__entity_link
func (c *Client) UpdateEntity(entityID string, entity Entity) error {
	entityJSON, err := json.Marshal(entity)
	if err != nil {
		return err
	}

	resp, err := c.request(http.MethodPut, fmt.Sprintf("/entities/%s", url.PathEscape(entityID)), "application/json", bytes.NewBuffer(entityJSON))
	if err != nil {
		return err
	}

	defer resp.Close()

	decoder := json.NewDecoder(resp)
	return decoder.Decode(&entity)
}

// DeleteEntity - deletes entity by ID or name.
//
// https://wit.ai/docs/http/20200513/#delete__entities__entity_link
func (c *Client) DeleteEntity(entityID string) error {
	resp, err := c.request(http.MethodDelete, fmt.Sprintf("/entities/%s", url.PathEscape(entityID)), "application/json", nil)
	if err == nil {
		resp.Close()
	}

	return err
}

// DeleteEntityRole - deletes entity role.
//
// https://wit.ai/docs/http/20200513/#delete__entities__entity_role_link
func (c *Client) DeleteEntityRole(entityID string, role string) error {
	resp, err := c.request(http.MethodDelete, fmt.Sprintf("/entities/%s:%s", url.PathEscape(entityID), url.PathEscape(role)), "application/json", nil)
	if err == nil {
		resp.Close()
	}

	return err
}

// AddEntityKeyword - Add a possible value into the list of values for the keyword entity.
//
// https://wit.ai/docs/http/20200513/#post__entities__entity_keywords_link
func (c *Client) AddEntityKeyword(entityID string, keyword EntityKeyword) (*Entity, error) {
	valueJSON, err := json.Marshal(keyword)
	if err != nil {
		return nil, err
	}

	resp, err := c.request(http.MethodPost, fmt.Sprintf("/entities/%s/keywords", url.PathEscape(entityID)), "application/json", bytes.NewBuffer(valueJSON))
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

// DeleteEntityKeyword - Delete a keyword from the keywords entity.
//
// https://wit.ai/docs/http/20200513/#delete__entities__entity_keywords__keyword_link
func (c *Client) DeleteEntityKeyword(entityID string, keyword string) error {
	resp, err := c.request(http.MethodDelete, fmt.Sprintf("/entities/%s/keywords/%s", url.PathEscape(entityID), url.PathEscape(keyword)), "application/json", nil)
	if err == nil {
		resp.Close()
	}

	return err
}

// AddEntityKeywordSynonym - Create a new synonym of the canonical value of the keywords entity.
//
// https://wit.ai/docs/http/20200513/#post__entities__entity_keywords__keyword_synonyms_link
func (c *Client) AddEntityKeywordSynonym(entityID string, keyword string, synonym string) (*Entity, error) {
	type syn struct {
		Synonym string `json:"synonym"`
	}

	exprJSON, err := json.Marshal(syn{
		Synonym: synonym,
	})
	if err != nil {
		return nil, err
	}

	resp, err := c.request(http.MethodPost, fmt.Sprintf("/entities/%s/keywords/%s/synonyms", url.PathEscape(entityID), url.PathEscape(keyword)), "application/json", bytes.NewBuffer(exprJSON))
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

// DeleteEntityKeywordSynonym - Delete a synonym of the keyword of the entity.
//
// https://wit.ai/docs/http/20200513/#delete__entities__entity_keywords__keyword_synonyms__synonym_link
func (c *Client) DeleteEntityKeywordSynonym(entityID string, keyword string, expression string) error {
	resp, err := c.request(http.MethodDelete, fmt.Sprintf("/entities/%s/keywords/%s/synonyms/%s", url.PathEscape(entityID), url.PathEscape(keyword), url.PathEscape(expression)), "application/json", nil)
	if err == nil {
		resp.Close()
	}

	return err
}
