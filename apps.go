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

// App - https://wit.ai/docs/http/20200513/#get__apps_link
type App struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name"`
	Lang    string `json:"lang"`
	Private bool   `json:"private"`

	// Timezone is only used when creating/updating an app; it's
	// not available when getting the details of an app.
	Timezone string `json:"timezone,omitempty"`

	CreatedAt Time `json:"created_at,omitempty"`

	WillTrainAt              Time              `json:"will_train_at,omitempty"`
	LastTrainedAt            Time              `json:"last_trained_at,omitempty"`
	LastTrainingDurationSecs int               `json:"last_training_duration_secs,omitempty"`
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

// CreatedApp - https://wit.ai/docs/http/20200513/#post__apps_link
type CreatedApp struct {
	AccessToken string `json:"access_token"`
	AppID       string `json:"app_id"`
}

// GetApps - Returns an array of all apps that you own.
//
// https://wit.ai/docs/http/20200513/#get__apps_link
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

// GetApp - Returns an object representation of the specified app.
//
// https://wit.ai/docs/http/20200513/#get__apps__app_link
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

// CreateApp - creates new app.
//
// https://wit.ai/docs/http/20200513/#post__apps_link
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

// UpdateApp - Updates an app.
//
// https://wit.ai/docs/http/20200513/#put__apps__app_link
func (c *Client) UpdateApp(id string, app App) error {
	appJSON, err := json.Marshal(app)
	if err != nil {
		return err
	}

	resp, err := c.request(http.MethodPut, fmt.Sprintf("/apps/%s", url.PathEscape(id)), "application/json", bytes.NewBuffer(appJSON))
	if err == nil {
		resp.Close()
	}

	return err
}

// DeleteApp - deletes app by ID.
//
// https://wit.ai/docs/http/20200513/#delete__apps__app_link
func (c *Client) DeleteApp(id string) error {
	resp, err := c.request(http.MethodDelete, fmt.Sprintf("/apps/%s", url.PathEscape(id)), "application/json", nil)
	if err == nil {
		resp.Close()
	}

	return err
}

// AppTag - https://wit.ai/docs/http/20200513/#get__apps__app_tags__tag_link
type AppTag struct {
	Name string `json:"name,omitempty"`
	Desc string `json:"desc,omitempty"`

	CreatedAt Time `json:"created_at,omitempty"`
	UpdatedAt Time `json:"updated_at,omitempty"`
}

// GetAppTags - Returns an array of all tag groups for an app.
// Within a group, all tags point to the same app state (as a result of moving tags).
//
// https://wit.ai/docs/http/20200513/#get__apps__app_tags_link
func (c *Client) GetAppTags(appID string) ([][]AppTag, error) {
	resp, err := c.request(http.MethodGet, fmt.Sprintf("/apps/%s/tags", url.PathEscape(appID)), "application/json", nil)
	if err != nil {
		return nil, err
	}

	defer resp.Close()

	var tags [][]AppTag
	decoder := json.NewDecoder(resp)
	err = decoder.Decode(&tags)
	return tags, err
}

// GetAppTag - returns the detail of the specified tag.
//
// https://wit.ai/docs/http/20200513/#get__apps__app_tags__tag_link
func (c *Client) GetAppTag(appID, tagID string) (*AppTag, error) {
	resp, err := c.request(http.MethodGet, fmt.Sprintf("/apps/%s/tags/%s", url.PathEscape(appID), url.PathEscape(tagID)), "application/json", nil)
	if err != nil {
		return nil, err
	}

	defer resp.Close()

	var tag *AppTag
	decoder := json.NewDecoder(resp)
	err = decoder.Decode(&tag)
	return tag, err
}

// CreateAppTag - Take a snapshot of the current app state, save it as a tag (version)
// of the app. The name of the tag created will be returned in the response.
//
// https://wit.ai/docs/http/20200513/#post__apps__app_tags_link
func (c *Client) CreateAppTag(appID string, tag string) (*AppTag, error) {
	type appTag struct {
		Tag string `json:"tag"`
	}

	tagJSON, err := json.Marshal(appTag{Tag: tag})
	if err != nil {
		return nil, err
	}

	resp, err := c.request(http.MethodPost, fmt.Sprintf("/apps/%s/tags", url.PathEscape(tag)), "application/json", bytes.NewBuffer(tagJSON))
	if err != nil {
		return nil, err
	}

	defer resp.Close()

	// theresponse format is different than the one in get API.
	var tmp appTag
	decoder := json.NewDecoder(resp)
	if err := decoder.Decode(&tmp); err != nil {
		return nil, err
	}

	return &AppTag{Name: tmp.Tag}, nil
}

// UpdateAppTagRequest - https://wit.ai/docs/http/20200513/#put__apps__app_tags__tag_link
type UpdateAppTagRequest struct {
	Tag    string `json:"tag,omitempty"`
	Desc   string `json:"desc,omitempty"`
	MoveTo string `json:"move_to,omitempty"`
}

// UpdateAppTagResponse - https://wit.ai/docs/http/20200513/#put__apps__app_tags__tag_link
type UpdateAppTagResponse struct {
	Tag     string `json:"tag,omitempty"`
	Desc    string `json:"desc,omitempty"`
	MovedTo string `json:"moved_to,omitempty"`
}

// UpdateAppTag - Update the tag's name or description
//
// https://wit.ai/docs/http/20200513/#put__apps__app_tags__tag_link
func (c *Client) UpdateAppTag(appID, tagID string, updated AppTag) (*AppTag, error) {
	type tag struct {
		Tag  string `json:"tag,omitempty"`
		Desc string `json:"desc,omitempty"`
	}

	updateJSON, err := json.Marshal(tag{Tag: updated.Name, Desc: updated.Desc})
	if err != nil {
		return nil, err
	}

	resp, err := c.request(http.MethodPut, fmt.Sprintf("/apps/%s/tags/%s", url.PathEscape(appID), url.PathEscape(tagID)), "application/json", bytes.NewBuffer(updateJSON))
	if err != nil {
		return nil, err
	}

	defer resp.Close()

	var tagResp tag
	decoder := json.NewDecoder(resp)
	err = decoder.Decode(&tagResp)
	return &AppTag{Name: tagResp.Tag, Desc: tagResp.Desc}, err
}

type MovedAppTag struct {
	Tag     string `json:"tag"`
	Desc    string `json:"desc"`
	MovedTo string `json:"moved_to"`
}

// MoveAppTag - move the tag to point to another tag.
//
// https://wit.ai/docs/http/20200513/#put__apps__app_tags__tag_link
func (c *Client) MoveAppTag(appID, tagID string, to string, updated *AppTag) (*MovedAppTag, error) {
	type tag struct {
		Tag    string `json:"tag,omitempty"`
		Desc   string `json:"desc,omitempty"`
		MoveTo string `json:"move_to,omitempty"`
	}

	updateReq := tag{MoveTo: to}
	if updated != nil {
		updateReq.Tag = updated.Name
		updateReq.Desc = updated.Desc
	}

	updateJSON, err := json.Marshal(updateReq)
	if err != nil {
		return nil, err
	}

	resp, err := c.request(http.MethodPut, fmt.Sprintf("/apps/%s/tags/%s", url.PathEscape(appID), url.PathEscape(tagID)), "application/json", bytes.NewBuffer(updateJSON))
	if err != nil {
		return nil, err
	}

	defer resp.Close()

	var tagResp *MovedAppTag
	decoder := json.NewDecoder(resp)
	err = decoder.Decode(&tagResp)
	return tagResp, err
}

// DeleteAppTag - Permanently delete the tag.
//
// https://wit.ai/docs/http/20200513/#delete__apps__app_tags__tag_link
func (c *Client) DeleteAppTag(appID, tagID string) error {
	resp, err := c.request(http.MethodDelete, fmt.Sprintf("/apps/%s/tags/%s", url.PathEscape(appID), url.PathEscape(tagID)), "application/json", nil)
	if err == nil {
		resp.Close()
	}

	return err
}
