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

func Test_buildParseQuery(t *testing.T) {
	want := "?q=" + "hello+world%26foo" +
		"&n=1&tag=tag" +
		"&context=" +
		"%7B" +
		"%22reference_time%22%3A%222014-10-30T12%3A18%3A45-07%3A00%22%2C" +
		"%22timezone%22%3A%22America%2FLos_Angeles%22%2C" +
		"%22locale%22%3A%22en_US%22%2C" +
		"%22coords%22%3A%7B%22lat%22%3A32.47104%2C%22long%22%3A-122.14703%7D" +
		"%7D"

	got := buildParseQuery(&MessageRequest{
		Query: "hello world&foo",
		N:     1,
		Tag:   "tag",
		Context: &MessageContext{
			Locale: "en_US",
			Coords: MessageCoords{
				Lat:  32.47104,
				Long: -122.14703,
			},
			Timezone:      "America/Los_Angeles",
			ReferenceTime: "2014-10-30T12:18:45-07:00",
		},
	})

	if got != want {
		t.Fatalf("expected \n\tquery = %v \n\tgot = %v", want, got)
	}
}
