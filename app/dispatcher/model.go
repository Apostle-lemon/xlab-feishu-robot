package dispatcher

type CallbackType func(map[string]any)

type FeishuEventRequest struct {
	EventId   string
	EventType string
	Token     string
	Event     map[string]any
}

func readFromDict(data map[string]any) FeishuEventRequest {
	eventType := ""
	token := ""
	eventId := ""
	event, _ := data["event"].(map[string]any)

	if _, exist := data["schema"]; exist {
		// v2
		header, _ := data["header"].(map[string]any)
		eventType, _ = header["event_type"].(string)
		token, _ = header["token"].(string)
		eventId, _ = header["event_id"].(string)
	} else {
		// v1
		eventType, _ = event["type"].(string)
		token, _ = event["token"].(string)
		eventId, _ = event["uuid"].(string)
	}

	return FeishuEventRequest{
		EventId:   eventId,
		EventType: eventType,
		Token:     token,
		Event:     event,
	}
}
