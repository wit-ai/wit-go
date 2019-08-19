// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.

package witai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// AppTrainingStatus - Represents the status of an app
type AppTrainingStatus string

const (
	// Done status
	Done AppTrainingStatus = "done"
	// Scheduled status
	Scheduled AppTrainingStatus = "scheduled"
	// Ongoing status
	Ongoing AppTrainingStatus = "ongoing"
)

// App - https://wit.ai/docs/http/20170307#get__apps_link
type App struct {
	Name    string `json:"name"`
	Lang    string `json:"lang"`
	Private bool   `json:"private"`
	// Description presents when we get an app
	Description string `json:"description,omitempty"`
	// Use Desc when create an app
	Desc string `json:"desc,omitempty"`
	// ID presents when we get an app
	ID string `json:"id,omitempty"`
	// AppID presents when we create an app
	AppID     string `json:"app_id,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	Timezone  string `json:"timezone,omitempty"`
	// Training information
	LastTrainingDurationSecs int               `json:"last_training_duration_secs,omitempty"`
	WillTrainAt              Time              `json:"will_train_at,omitempty"`
	LastTrainedAt            Time              `json:"last_trained_at,omitempty"`
	TrainingStatus           AppTrainingStatus `json:"training_status,omitempty"`
}

// Time - Custom type to encapsulated a time.Time
type Time struct {
	time.Time
}

// UnmarshalJSON - Our unmarshal function for our custom type
func (witTime *Time) UnmarshalJSON(input []byte) error {
	strInput := string(input)
	strInput = strings.Trim(strInput, `"`)
	newTime, err := time.Parse(WitTimeFormat, strInput)
	if err != nil {
		return err
	}
	witTime.Time = newTime
	return nil
}

// CreatedApp - https://wit.ai/docs/http/20170307#post__apps_link
type CreatedApp struct {
	AccessToken string `json:"access_token"`
	AppID       string `json:"app_id"`
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
	resp, err := c.request(http.MethodGet, fmt.Sprintf("/apps/%s", url.PathEscape(id)), "application/json", nil)
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
	resp, err := c.request(http.MethodDelete, fmt.Sprintf("/apps/%s", url.PathEscape(id)), "application/json", nil)
	if err == nil {
		resp.Close()
	}

	return err
}

// CreateApp - creates new app. https://wit.ai/docs/http/20170307#post__apps_link
func (c *Client) CreateApp(app App) (*CreatedApp, error) {
	appJSON, err := json.Marshal(app)
	if err != nil {
		return nil, err
	}

	resp, err := c.request(http.MethodPost, "/apps", "application/json", bytes.NewBuffer(appJSON))
	if err != nil {
		return nil, err
	}

	defer resp.Close()

	var createdApp *CreatedApp
	decoder := json.NewDecoder(resp)
	err = decoder.Decode(&createdApp)

	return createdApp, err
}

// UpdateApp - Updates an app. https://wit.ai/docs/http/20170307#put__apps__app_id_link
func (c *Client) UpdateApp(id string, app App) (*App, error) {
	appJSON, err := json.Marshal(app)
	if err != nil {
		return nil, err
	}

	resp, err := c.request(http.MethodPut, fmt.Sprintf("/apps/%s", url.PathEscape(id)), "application/json", bytes.NewBuffer(appJSON))
	if err != nil {
		return nil, err
	}

	defer resp.Close()

	var updatedApp *App
	decoder := json.NewDecoder(resp)
	err = decoder.Decode(&updatedApp)

	return updatedApp, err
}
