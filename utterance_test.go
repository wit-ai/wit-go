package witai

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetUtterances(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`[
			{
				"text": "text 1",
				"intent": {
					"id": "i1",
					"name": "intent1"
				},
				"entities": [
					{
						"id": "e1",
						"name": "entity1",
						"role": "role1",
						"start": 0,
						"end": 10,
						"body": "entity1",
						"entities": [
							{
								"id": "sub-e1",
								"name": "sub-entity1",
								"role": "sub-role1",
								"start": 5,
								"end": 7,
								"body": "sub-entity1",
								"entities": null
							}
						]
					},
					{
						"id": "e2",
						"name": "entity2",
						"role": "role2",
						"start": 10,
						"end": 20,
						"body": "entity2",
						"entities": [
							{
								"id": "sub-e2",
								"name": "sub-entity2",
								"role": "sub-role2",
								"start": 15,
								"end": 17,
								"body": "sub-entity2",
								"entities": null
							}
						]
					}
				],
				"traits": [
					{
						"id": "t1",
						"name": "trait1",
						"value": "value1"
					},
					{
						"id": "t2",
						"name": "trait2",
						"value": "value2"
					}
				]
			},
			{
				"text": "text 2",
				"intent": {
					"id": "i2",
					"name": "intent2"
				},
				"entities": [
					{
						"id": "e1",
						"name": "entity1",
						"role": "role1",
						"start": 0,
						"end": 10,
						"body": "entity1",
						"entities": [
							{
								"id": "sub-e1",
								"name": "sub-entity1",
								"role": "sub-role1",
								"start": 5,
								"end": 7,
								"body": "sub-entity1",
								"entities": null
							}
						]
					},
					{
						"id": "e2",
						"name": "entity2",
						"role": "role2",
						"start": 10,
						"end": 20,
						"body": "entity2",
						"entities": [
							{
								"id": "sub-e2",
								"name": "sub-entity2",
								"role": "sub-role2",
								"start": 15,
								"end": 17,
								"body": "sub-entity2",
								"entities": null
							}
						]
					}
				],
				"traits": [
					{
						"id": "t1",
						"name": "trait1",
						"value": "value1"
					},
					{
						"id": "t2",
						"name": "trait2",
						"value": "value2"
					}
				]
			}
		]`))
	}))
	defer testServer.Close()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL
	utterances, _ := c.GetUtterances(5, 0)

	wantUtterances := []Utterance{
		{
			Text:   "text 1",
			Intent: UtteranceIntent{ID: "i1", Name: "intent1"},
			Entities: []UtteranceEntity{
				{
					ID:    "e1",
					Name:  "entity1",
					Role:  "role1",
					Start: 0,
					End:   10,
					Body:  "entity1",
					Entities: []UtteranceEntity{
						{
							ID:    "sub-e1",
							Name:  "sub-entity1",
							Role:  "sub-role1",
							Start: 5,
							End:   7,
							Body:  "sub-entity1",
						},
					},
				},
				{
					ID:    "e2",
					Name:  "entity2",
					Role:  "role2",
					Start: 10,
					End:   20,
					Body:  "entity2",
					Entities: []UtteranceEntity{
						{
							ID:    "sub-e2",
							Name:  "sub-entity2",
							Role:  "sub-role2",
							Start: 15,
							End:   17,
							Body:  "sub-entity2",
						},
					},
				},
			},
			Traits: []UtteranceTrait{
				{ID: "t1", Name: "trait1", Value: "value1"},
				{ID: "t2", Name: "trait2", Value: "value2"},
			},
		},
		{
			Text:   "text 2",
			Intent: UtteranceIntent{ID: "i2", Name: "intent2"},
			Entities: []UtteranceEntity{
				{
					ID:    "e1",
					Name:  "entity1",
					Role:  "role1",
					Start: 0,
					End:   10,
					Body:  "entity1",
					Entities: []UtteranceEntity{
						{
							ID:    "sub-e1",
							Name:  "sub-entity1",
							Role:  "sub-role1",
							Start: 5,
							End:   7,
							Body:  "sub-entity1",
						},
					},
				},
				{
					ID:    "e2",
					Name:  "entity2",
					Role:  "role2",
					Start: 10,
					End:   20,
					Body:  "entity2",
					Entities: []UtteranceEntity{
						{
							ID:    "sub-e2",
							Name:  "sub-entity2",
							Role:  "sub-role2",
							Start: 15,
							End:   17,
							Body:  "sub-entity2",
						},
					},
				},
			},
			Traits: []UtteranceTrait{
				{ID: "t1", Name: "trait1", Value: "value1"},
				{ID: "t2", Name: "trait2", Value: "value2"},
			},
		},
	}

	if !reflect.DeepEqual(utterances, wantUtterances) {
		t.Fatalf("expected \n\tmsg %+v \n\tgot %+v", wantUtterances, utterances)
	}
}

func TestDeleteUtterances(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"sent": true, "n": 2}`))
	}))
	defer testServer.Close()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL
	resp, _ := c.DeleteUtterances([]string{"text1", "text2"})

	wantResp := &TrainingResponse{
		Sent: true,
		N:    2,
	}

	if !reflect.DeepEqual(resp, wantResp) {
		t.Fatalf("expected \n\tresp %+v \n\tgot %+v", wantResp, resp)
	}
}

func TestTrainUtterances(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"sent": true, "n": 2}`))
	}))
	defer testServer.Close()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL
	resp, _ := c.TrainUtterances([]Training{
		{Text: "text1"},
	})

	wantResp := &TrainingResponse{
		Sent: true,
		N:    2,
	}

	if !reflect.DeepEqual(resp, wantResp) {
		t.Fatalf("expected \n\tresp %+v \n\tgot %+v", wantResp, resp)
	}
}
