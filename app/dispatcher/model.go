package dispatcher

import "encoding/json"

type CallbackType func(map[string]any)

type FeishuEventRequestRaw struct {
	Header struct {
		EventType string `json:"event_type"`
		Token     string `json:"token"`
		EventId   string `json:"event_id"`
	} `json:"header"`
	Schema    string         `json:"schema"`
	Uuid      string         `json:"uuid"`
	Type      string         `json:"type"`
	Token     string         `json:"token"`
	Event     map[string]any `json:"event"`
	Challenge string         `json:"challenge"`
}

type FeishuEventRequest struct {
	EventId   string
	EventType string
	Token     string
	Event     map[string]any
	Challenge string
}

func deserializeRequest(dataStr string, request *FeishuEventRequest) {
	var data FeishuEventRequestRaw
	json.Unmarshal([]byte(dataStr), &data)

	request.Challenge = data.Challenge
	request.Event = data.Event

	if data.Schema != "" {
		// v2
		request.EventId = data.Header.EventId
		request.EventType = data.Header.EventType
		request.Token = data.Header.Token
	} else {
		// v1
		request.EventId = data.Uuid
		request.EventType = data.Type
		request.Token = data.Token
	}
}
