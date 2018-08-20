package witai

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// MessageResponse - https://wit.ai/docs/http/20170307#get__message_link
type MessageResponse struct {
	ID       string                 `json:"msg_id"`
	Text     string                 `json:"_text"`
	Entities map[string]interface{} `json:"entities"`
}

// MessageRequest - https://wit.ai/docs/http/20170307#get__message_link
type MessageRequest struct {
	Query    string `json:"q"`
	MsgID    string `json:"msg_id"`
	N        int    `json:"n"`
	ThreadID string `json:"thread_id"`
}

// MessageContext - https://wit.ai/docs/http/20170307#context_link
type MessageContext struct {
	TeferenceTime string        `json:"reference_time"` // "2014-10-30T12:18:45-07:00"
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
	q := fmt.Sprintf("?q=%s", url.QueryEscape(req.Query))
	if len(req.MsgID) != 0 {
		q += fmt.Sprintf("&msg_id=%s", req.MsgID)
	}
	if req.N != 0 {
		q += fmt.Sprintf("&n=%d", req.N)
	}
	if len(req.ThreadID) != 0 {
		q += fmt.Sprintf("&thread_id=%s", req.ThreadID)
	}

	resp, err := c.request(http.MethodGet, "/message"+q, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Close()

	var msgResp *MessageResponse
	decoder := json.NewDecoder(resp)
	err = decoder.Decode(&msgResp)
	if err != nil {
		return nil, err
	}

	return msgResp, nil
}
