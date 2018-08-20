package witai

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestParse(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`{"msg_id": "msg1", "entities": {"e1": "d1"}}`))
	}))
	defer func() { testServer.Close() }()

	apiBase = testServer.URL

	c := NewClient("token")
	msg, _ := c.Parse(&MessageRequest{
		Query: "hello",
	})

	if msg == nil || msg.ID != "msg1" || len(msg.Entities) != 1 {
		t.Fatalf("msg is not parsed correctly")
	}
}
