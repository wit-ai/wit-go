// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.

package witai

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Locales - https://wit.ai/docs/http/20170307#get__language_link
type Locales struct {
	DetectedLocales []Locale `json:"detected_locales"`
}

// Locale - https://wit.ai/docs/http/20170307#get__language_link
type Locale struct {
	Locale     string  `json:"locale"`
	Confidence float64 `json:"confidence"`
}

// Detect - returns the detected languages from query - https://wit.ai/docs/http/20170307#get__language_link
func (c *Client) Detect(text string) (*Locales, error) {
	resp, err := c.request(http.MethodGet, fmt.Sprintf("/language?q=%s", url.PathEscape(text)), "application/json", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Close()

	var locales *Locales
	decoder := json.NewDecoder(resp)
	err = decoder.Decode(&locales)
	if err != nil {
		return nil, err
	}

	return locales, nil
}
