// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.

package witai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Sample - https://wit.ai/docs/http/20170307#get__samples_link
type Sample struct {
	Text     string         `json:"text"`
	Entities []SampleEntity `json:"entities"`
}

// SampleEntity - https://wit.ai/docs/http/20170307#get__samples_link
type SampleEntity struct {
	Entity       string         `json:"entity"`
	Value        string         `json:"value"`
	Role         string         `json:"role"`
	Start        int            `json:"start,omitempty"`
	End          int            `json:"end,omitempty"`
	Subentitites []SampleEntity `json:"subentities"`
}

// ValidateSampleResponse - https://wit.ai/docs/http/20170307#post__samples_link
type ValidateSampleResponse struct {
	Sent bool `json:"sent"`
	N    int  `json:"n"`
}

// GetSamples - Returns an array of samples. https://wit.ai/docs/http/20170307#get__samples_link
func (c *Client) GetSamples(limit int, offset int) ([]Sample, error) {
	if limit <= 0 {
		limit = 0
	}
	if offset <= 0 {
		offset = 0
	}

	resp, err := c.request(http.MethodGet, fmt.Sprintf("/samples?limit=%d&offset=%d", limit, offset), "application/json", nil)
	if err != nil {
		return []Sample{}, err
	}

	defer resp.Close()

	var samples []Sample
	decoder := json.NewDecoder(resp)
	err = decoder.Decode(&samples)
	return samples, err
}

// ValidateSamples - Validate samples (sentence + entities annotations) to train your app programmatically. https://wit.ai/docs/http/20170307#post__samples_link
func (c *Client) ValidateSamples(samples []Sample) (*ValidateSampleResponse, error) {
	samplesJSON, err := json.Marshal(samples)
	if err != nil {
		return nil, err
	}

	resp, err := c.request(http.MethodPost, "/samples", "application/json", bytes.NewBuffer(samplesJSON))
	if err != nil {
		return nil, err
	}

	defer resp.Close()

	var r *ValidateSampleResponse
	decoder := json.NewDecoder(resp)
	err = decoder.Decode(&r)
	return r, err
}

// DeleteSamples - Delete validated samples from your app. https://wit.ai/docs/http/20170307#delete__samples_link
// Only text property is required
func (c *Client) DeleteSamples(samples []Sample) (*ValidateSampleResponse, error) {
	samplesJSON, err := json.Marshal(samples)
	if err != nil {
		return nil, err
	}

	resp, err := c.request(http.MethodDelete, "/samples", "application/json", bytes.NewBuffer(samplesJSON))
	if err != nil {
		return nil, err
	}

	defer resp.Close()

	var r *ValidateSampleResponse
	decoder := json.NewDecoder(resp)
	err = decoder.Decode(&r)
	return r, err
}
