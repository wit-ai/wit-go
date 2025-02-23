// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.

package witai

import (
	"encoding/json"
	"io"
	"net/http"
)

type DictationRequest struct {
	File        io.Reader
	ContentType string
}

type DictationToken struct {
	End   int    `json:"end"`
	Start int    `json:"start"`
	Token string `json:"token"`
}

type DictationSpeech struct {
	Confidence float64          `json:"confidence"`
	Tokens     []DictationToken `json:"tokens"`
}

type DictationResponse struct {
	Speech DictationSpeech `json:"speech"`
	Text   string          `json:"text"`
	Type   string          `json:"type"`
}

// Dictation - Returns the text transcription from an audio file or stream.
func (c *Client) Dictation(req DictationRequest) (*DictationResponse, error) {
	resp, err := c.request(http.MethodPost, "/dictation", req.ContentType, req.File)
	if err != nil {
		return nil, err
	}

	defer resp.Close()

	var msgResp *DictationResponse
	decoder := json.NewDecoder(resp)
	for {
		err := decoder.Decode(&msgResp)
		if err != nil {
			break
		}
	}
	return msgResp, err
}
