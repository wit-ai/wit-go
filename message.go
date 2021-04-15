// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.

package witai

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// MessageResponse - https://wit.ai/docs/http/20200513/#get__message_link
type MessageResponse struct {
	ID       string                     `json:"msg_id"`
	Text     string                     `json:"text"`
	Intents  []MessageIntent            `json:"intents"`
	Entities map[string][]MessageEntity `json:"entities"`
	Traits   map[string][]MessageTrait  `json:"traits"`
}

// MessageEntity - https://wit.ai/docs/http/20200513/#get__message_link
type MessageEntity struct {
	ID         string                 `json:"id"`
	Name       string                 `json:"name"`
	Role       string                 `json:"role"`
	Start      int                    `json:"start"`
	End        int                    `json:"end"`
	Body       string                 `json:"body"`
	Value      string                 `json:"value"`
	Confidence float64                `json:"confidence"`
	Entities   []MessageEntity        `json:"entities"`
	Extra      map[string]interface{} `json:"-"`
}

// MessageTrait - https://wit.ai/docs/http/20200513/#get__message_link
type MessageTrait struct {
	ID         string                 `json:"id"`
	Value      string                 `json:"value"`
	Confidence float64                `json:"confidence"`
	Extra      map[string]interface{} `json:"-"`
}

// MessageIntent - https://wit.ai/docs/http/20200513/#get__message_link
type MessageIntent struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Confidence float64 `json:"confidence"`
}

// MessageRequest - https://wit.ai/docs/http/20200513/#get__message_link
type MessageRequest struct {
	Query   string          `json:"q"`
	Tag     string          `json:"tag"`
	N       int             `json:"n"`
	Context *MessageContext `json:"context"`
	Speech  *Speech         `json:"-"`
}

// Speech - https://wit.ai/docs/http/20170307#post__speech_link
type Speech struct {
	File        io.Reader `json:"file"`
	ContentType string    `json:"content_type"` // Example: audio/raw;encoding=unsigned-integer;bits=16;rate=8000;endian=big
}

// MessageContext - https://wit.ai/docs/http/20170307#context_link
type MessageContext struct {
	ReferenceTime string        `json:"reference_time"` // "2014-10-30T12:18:45-07:00"
	Timezone      string        `json:"timezone"`
	Locale        string        `json:"locale"`
	Coords        MessageCoords `json:"coords"`
}

// MessageCoords - https://wit.ai/docs/http/20170307#context_link
type MessageCoords struct {
	Lat  float32 `json:"lat"`
	Long float32 `json:"long"`
}

// Parse - parses text and returns entities
func (c *Client) Parse(req *MessageRequest) (*MessageResponse, error) {
	if req == nil {
		return nil, errors.New("invalid request")
	}

	q := buildParseQuery(req)

	resp, err := c.request(http.MethodGet, "/message"+q, "application/json", nil)
	if err != nil {
		return nil, err
	}

	defer resp.Close()

	var msgResp *MessageResponse
	decoder := json.NewDecoder(resp)
	err = decoder.Decode(&msgResp)
	return msgResp, err
}

// Speech - sends audio file for parsing
func (c *Client) Speech(req *MessageRequest) (*MessageResponse, error) {
	if req == nil || req.Speech == nil {
		return nil, errors.New("invalid request")
	}

	q := buildParseQuery(req)

	resp, err := c.request(http.MethodPost, "/speech"+q, req.Speech.ContentType, req.Speech.File)
	if err != nil {
		return nil, err
	}

	defer resp.Close()

	var msgResp *MessageResponse
	decoder := json.NewDecoder(resp)
	err = decoder.Decode(&msgResp)
	return msgResp, err
}

func buildParseQuery(req *MessageRequest) string {
	q := fmt.Sprintf("?q=%s", url.PathEscape(req.Query))
	if req.N != 0 {
		q += fmt.Sprintf("&n=%d", req.N)
	}
	if req.Tag != "" {
		q += fmt.Sprintf("&tag=%s", req.Tag)
	}
	if req.Context != nil {
		b, _ := json.Marshal(req.Context)
		if b != nil {
			q += fmt.Sprintf("&context=%s", url.PathEscape(string(b)))
		}
	}

	return q
}
