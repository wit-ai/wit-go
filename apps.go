package witai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// App - https://wit.ai/docs/http/20170307#get__apps_link
type App struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Lang        string `json:"lang"`
	Private     bool   `json:"private"`
	ID          string `json:"id,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	Timezone    string `json:"timezone,omitempty"`
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

// GetApp - returns map by ID. https://wit.ai/docs/http/20170307#get__apps__app_id_link
func (c *Client) GetApp(id string) (*App, error) {
	resp, err := c.request(http.MethodGet, fmt.Sprintf("/apps/%s", url.QueryEscape(id)), "application/json", nil)
	if err != nil {
		return nil, err
	}

	defer resp.Close()

	var app *App
	decoder := json.NewDecoder(resp)
	if err = decoder.Decode(&app); err != nil {
		return nil, err
	}

	return app, nil
}

// DeleteApp - deletes app by ID. https://wit.ai/docs/http/20170307#delete__apps__app_id_link
func (c *Client) DeleteApp(id string) error {
	resp, err := c.request(http.MethodDelete, fmt.Sprintf("/apps/%s", url.QueryEscape(id)), "application/json", nil)
	if err == nil {
		resp.Close()
	}

	return err
}

// CreateApp - creates new app. https://wit.ai/docs/http/20170307#post__apps_link
func (c *Client) CreateApp(app *App) error {
	appJSON, err := json.Marshal(app)
	if err != nil {
		return err
	}

	resp, err := c.request(http.MethodPost, "/apps", "application/json", bytes.NewBuffer(appJSON))
	if err != nil {
		return err
	}

	defer resp.Close()

	decoder := json.NewDecoder(resp)
	return decoder.Decode(&app)
}

// UpdateApp - Updates an app. https://wit.ai/docs/http/20170307#put__apps__app_id_link
func (c *Client) UpdateApp(id string, app *App) error {
	appJSON, err := json.Marshal(app)
	if err != nil {
		return err
	}

	resp, err := c.request(http.MethodPut, fmt.Sprintf("/apps/%s", url.QueryEscape(id)), "application/json", bytes.NewBuffer(appJSON))
	if err != nil {
		return err
	}

	defer resp.Close()

	decoder := json.NewDecoder(resp)
	return decoder.Decode(&app)
}
