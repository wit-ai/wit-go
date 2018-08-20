package witai

import (
	"bytes"
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
		t.Fatalf("expected message id: msg1 and 1 entity, got %v", msg)
	}
}

func TestSpeech(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`{"msg_id": "msg1", "entities": {"e1": "d1"}}`))
	}))
	defer func() { testServer.Close() }()

	apiBase = testServer.URL

	c := NewClient("token")
	msg, _ := c.Speech(&MessageRequest{
		Speech: &Speech{
			File:        bytes.NewReader([]byte{}),
			ContentType: "audio/raw;encoding=unsigned-integer;bits=16;rate=16k;endian=little",
		},
	})

	if msg == nil || msg.ID != "msg1" || len(msg.Entities) != 1 {
		t.Fatalf("expected message id: msg1 and 1 entity, got %v", msg)
	}
}
