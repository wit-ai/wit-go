// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.

package witai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Utterance - https://wit.ai/docs/http/20200513/#get__utterances_link
type Utterance struct {
	Text     string            `json:"text"`
	Intent   UtteranceIntent   `json:"intent"`
	Entities []UtteranceEntity `json:"entities"`
	Traits   []UtteranceTrait  `json:"traits"`
}

// UtteranceIntent - https://wit.ai/docs/http/20200513/#get__utterances_link
type UtteranceIntent struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// UtteranceEntity - https://wit.ai/docs/http/20200513/#get__utterances_link
type UtteranceEntity struct {
	ID       string            `json:"id"`
	Name     string            `json:"name"`
	Role     string            `json:"role"`
	Start    int               `json:"start"`
	End      int               `json:"end"`
	Body     string            `json:"body"`
	Entities []UtteranceEntity `json:"entities"`
}

// UtteranceTrait - https://wit.ai/docs/http/20200513/#get__utterances_link
type UtteranceTrait struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

// GetUtterances - Returns an array of utterances.
//
// https://wit.ai/docs/http/20200513/#get__utterances_link
func (c *Client) GetUtterances(limit int, offset int) ([]Utterance, error) {
	if limit <= 0 {
		limit = 0
	}
	if offset <= 0 {
		offset = 0
	}

	resp, err := c.request(http.MethodGet, fmt.Sprintf("/utterances?limit=%d&offset=%d", limit, offset), "application/json", nil)
	if err != nil {
		return []Utterance{}, err
	}

	defer resp.Close()

	var utterances []Utterance
	decoder := json.NewDecoder(resp)
	err = decoder.Decode(&utterances)
	return utterances, err
}

// DeleteUtterances - Delete validated utterances from your app.
//
// https://wit.ai/docs/http/20200513/#delete__utterances_link
func (c *Client) DeleteUtterances(texts []string) (*TrainingResponse, error) {
	type text struct {
		Text string `json:"text"`
	}
	reqTexts := make([]text, len(texts))
	for i, t := range texts {
		reqTexts[i] = text{Text: t}
	}

	utterancesJSON, err := json.Marshal(reqTexts)
	if err != nil {
		return nil, err
	}

	resp, err := c.request(http.MethodDelete, "/utterances", "application/json", bytes.NewBuffer(utterancesJSON))
	if err != nil {
		return nil, err
	}

	defer resp.Close()

	var r *TrainingResponse
	decoder := json.NewDecoder(resp)
	err = decoder.Decode(&r)
	return r, err
}

// TrainingResponse - https://wit.ai/docs/http/20200513/#post__utterances_link
type Training struct {
	Text     string           `json:"text"`
	Intent   string           `json:"intent,omitempty"`
	Entities []TrainingEntity `json:"entities"`
	Traits   []TrainingTrait  `json:"traits"`
}

// TrainingResponse - https://wit.ai/docs/http/20200513/#post__utterances_link
type TrainingEntity struct {
	Entity   string           `json:"entity"`
	Start    int              `json:"start"`
	End      int              `json:"end"`
	Body     string           `json:"body"`
	Entities []TrainingEntity `json:"entities"`
}

// TrainingResponse - https://wit.ai/docs/http/20200513/#post__utterances_link
type TrainingTrait struct {
	Trait string `json:"trait"`
	Value string `json:"value"`
}

// TrainingResponse - https://wit.ai/docs/http/20200513/#post__utterances_link
type TrainingResponse struct {
	Sent bool `json:"sent"`
	N    int  `json:"n"`
}

// TrainUtterances - Add utterances (sentence + entities annotations) to train your app programmatically.
//
// https://wit.ai/docs/http/20200513/#post__utterances_link
func (c *Client) TrainUtterances(trainings []Training) (*TrainingResponse, error) {
	utterancesJSON, err := json.Marshal(trainings)
	if err != nil {
		return nil, err
	}

	resp, err := c.request(http.MethodPost, "/utterances", "application/json", bytes.NewBuffer(utterancesJSON))
	if err != nil {
		return nil, err
	}

	defer resp.Close()

	var r *TrainingResponse
	decoder := json.NewDecoder(resp)
	err = decoder.Decode(&r)
	return r, err
}
