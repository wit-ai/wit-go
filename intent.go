// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.

package witai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Intent - represents a wit-ai intent.
//
// https://wit.ai/docs/http/20200513/#get__intents_link
type Intent struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Entities []Entity `json:"entities,omitempty"`
}

// GetIntents - returns the list of intents.
//
// https://wit.ai/docs/http/20200513/#get__intents_link
func (c *Client) GetIntents() ([]Intent, error) {
	resp, err := c.request(http.MethodGet, "/intents", "application/json", nil)
	if err != nil {
		return []Intent{}, err
	}

	defer resp.Close()

	var intents []Intent
	err = json.NewDecoder(resp).Decode(&intents)
	return intents, err
}

// CreateIntent - creates a new intent with the given name.
//
// https://wit.ai/docs/http/20200513/#post__intents_link
func (c *Client) CreateIntent(name string) (*Intent, error) {
	intentJSON, err := json.Marshal(Intent{Name: name})
	if err != nil {
		return nil, err
	}

	resp, err := c.request(http.MethodPost, "/intents", "application/json", bytes.NewBuffer(intentJSON))
	if err != nil {
		return nil, err
	}

	defer resp.Close()

	var intentResp *Intent
	err = json.NewDecoder(resp).Decode(&intentResp)
	return intentResp, err
}

// GetIntent - returns intent by name.
//
// https://wit.ai/docs/http/20200513/#get__intents__intent_link
func (c *Client) GetIntent(name string) (*Intent, error) {
	resp, err := c.request(http.MethodGet, fmt.Sprintf("/intents/%s", url.PathEscape(name)), "application/json", nil)
	if err != nil {
		return nil, err
	}

	defer resp.Close()

	var intent *Intent
	err = json.NewDecoder(resp).Decode(&intent)
	return intent, err
}

// DeleteIntent - permanently deletes an intent by name.
//
// https://wit.ai/docs/http/20200513/#delete__intents__intent_link
func (c *Client) DeleteIntent(name string) error {
	resp, err := c.request(http.MethodDelete, fmt.Sprintf("/intents/%s", url.PathEscape(name)), "application/json", nil)
	if err == nil {
		resp.Close()
	}

	return err
}
