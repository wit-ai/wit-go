package witai

import "fmt"

// DefaultVersion - https://wit.ai/docs/http/20170307
const DefaultVersion = "20170307"

// Client - Wit.ai client type
type Client struct {
	Token        string
	Version      string
	headerAuth   string
	headerAccept string
}

// NewClient - returns Wit.ai client for default API version
func NewClient(token string) *Client {
	return NewClientWithVersion(token, DefaultVersion)
}

// NewClientWithVersion - returns Wit.ai client for specified API version
func NewClientWithVersion(token, version string) *Client {
	headerAuth := fmt.Sprintf("Bearer %s", token)
	headerAccept := fmt.Sprintf("application/vnd.wit.%s+json", version)

	return &Client{
		Token:        token,
		Version:      version,
		headerAuth:   headerAuth,
		headerAccept: headerAccept,
	}
}
