package witai

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// App - https://wit.ai/docs/http/20170307#get__apps_link
type App struct {
	Name        string `json:"name"`
	ID          string `json:"id"`
	Description string `json:"description"`
	Lang        string `json:"lang"`
	Private     bool   `json:"private"`
	CreatedAt   string `json:"created_at"`
}

// GetApps - Returns an array of all apps that you own. https://wit.ai/docs/http/20170307#get__apps_link
func (c *Client) GetApps(limit int, offset int) ([]App, error) {
	if limit <= 0 {
		limit = 0
	}
	if offset <= 0 {
		offset = 0
	}

	resp, err := c.request(http.MethodGet, fmt.Sprintf("/apps?limit=%d&offset=%d", limit, offset), "application/json", nil)
	if err != nil {
		return []App{}, err
	}

	defer resp.Close()

	var apps []App
	decoder := json.NewDecoder(resp)
	err = decoder.Decode(&apps)
	return apps, err
}
