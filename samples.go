package witai

// Sample - https://wit.ai/docs/http/20170307#get__samples_link
type Sample struct {
	Text     string         `json:"text"`
	Entities []SampleEntity `json:"entities"`
}

// SampleEntity - https://wit.ai/docs/http/20170307#get__samples_link
type SampleEntity struct {
	Entity       string         `json:"entity"`
	Value        string         `json:"value"`
	Role         string         `json:"role"`
	Start        int            `json:"start"`
	End          int            `json:"end"`
	Subentitites []SampleEntity `json:"subentities"`
}
