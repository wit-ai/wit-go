package witai

import (
	"encoding/json"
	"net/http"
)

func (c *Client) GetEntities() ([]string, error) {
	resp, err := c.request(http.MethodGet, "/entities", nil)
	if err != nil {
		return []string{}, err
	}

	defer resp.Close()

	var entities []string
	decoder := json.NewDecoder(resp)
	err = decoder.Decode(&entities)
	if err != nil {
		return []string{}, err
	}

	return entities, nil
}
