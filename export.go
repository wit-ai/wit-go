// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.

package witai

import (
	"encoding/json"
	"net/http"
)

type exportResponse struct {
	URI string `json:"uri"`
}

// Export - Returns download URI. https://wit.ai/docs/http/20170307#get__export_link
func (c *Client) Export() (string, error) {
	resp, err := c.request(http.MethodGet, "/export", "application/json", nil)
	if err != nil {
		return "", err
	}

	defer resp.Close()

	var r exportResponse
	decoder := json.NewDecoder(resp)
	err = decoder.Decode(&r)
	return r.URI, err
}
