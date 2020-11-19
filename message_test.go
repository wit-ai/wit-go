// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.

package witai

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`{
			"msg_id": "msg1",
			"text": "text",
			"intents": [
				{"id": "intent1", "name": "intent1_name", "confidence": 0.9},
				{"id": "intent2", "name": "intent2_name", "confidence": 0.7}
			],
			"entities": {
				"entity1": [{
					"id": "entity1-1",
					"name": "entity1",
					"role": "entity1",
					"start": 1,
					"end": 10,
					"body": "value1",
					"value": "value1",
					"confidence": 0.8
				}]
			},
			"traits": {
				"trait1": [{
					"id": "trait1-1",
					"value": "value1",
					"confidence": 0.8
				}]
			}
		}`))
	}))
	defer func() { testServer.Close() }()

	c := NewClient("token")
	c.APIBase = testServer.URL
	msg, _ := c.Parse(&MessageRequest{
		Query: "hello",
	})

	wantMessage := &MessageResponse{
		ID:   "msg1",
		Text: "text",
		Intents: []MessageIntent{
			{ID: "intent1", Name: "intent1_name", Confidence: 0.9},
			{ID: "intent2", Name: "intent2_name", Confidence: 0.7},
		},
		Entities: map[string][]MessageEntity{
			"entity1": {{
				ID:         "entity1-1",
				Name:       "entity1",
				Role:       "entity1",
				Start:      1,
				End:        10,
				Body:       "value1",
				Value:      "value1",
				Confidence: 0.8,
			}},
		},
		Traits: map[string][]MessageTrait{
			"trait1": {{ID: "trait1-1", Value: "value1", Confidence: 0.8}},
		},
	}

	if !reflect.DeepEqual(msg, wantMessage) {
		t.Fatalf("expected \n\tmsg %v \n\tgot %v", wantMessage, msg)
	}
}

func TestSpeech(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`{
			"msg_id": "msg1",
			"text": "text",
			"intents": [
				{"id": "intent1", "name": "intent1_name", "confidence": 0.9},
				{"id": "intent2", "name": "intent2_name", "confidence": 0.7}
			],
			"entities": {
				"entity1": [{
					"id": "entity1-1",
					"name": "entity1",
					"role": "entity1",
					"start": 1,
					"end": 10,
					"body": "value1",
					"value": "value1",
					"confidence": 0.8
				}]
			},
			"traits": {
				"trait1": [{
					"id": "trait1-1",
					"value": "value1",
					"confidence": 0.8
				}]
			}
		}`))
	}))
	defer func() { testServer.Close() }()

	c := NewClient("token")
	c.APIBase = testServer.URL
	msg, _ := c.Speech(&MessageRequest{
		Speech: &Speech{
			File:        bytes.NewReader([]byte{}),
			ContentType: "audio/raw;encoding=unsigned-integer;bits=16;rate=16k;endian=little",
		},
	})

	wantMessage := &MessageResponse{
		ID:   "msg1",
		Text: "text",
		Intents: []MessageIntent{
			{ID: "intent1", Name: "intent1_name", Confidence: 0.9},
			{ID: "intent2", Name: "intent2_name", Confidence: 0.7},
		},
		Entities: map[string][]MessageEntity{
			"entity1": {{
				ID:         "entity1-1",
				Name:       "entity1",
				Role:       "entity1",
				Start:      1,
				End:        10,
				Body:       "value1",
				Value:      "value1",
				Confidence: 0.8,
			}},
		},
		Traits: map[string][]MessageTrait{
			"trait1": {{ID: "trait1-1", Value: "value1", Confidence: 0.8}},
		},
	}

	if !reflect.DeepEqual(msg, wantMessage) {
		t.Fatalf("expected \n\tmsg %v \n\tgot %v", wantMessage, msg)
	}
}
