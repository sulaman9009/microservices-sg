package events

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Event struct {
	Type string `json:"type"`
	Data any    `json:"data"`
}

const (
	TypePostCreated    = "PostCreated"
	TypeCommentCreated = "CommentCreated"
)

func EmitEvent(eventType string, data any) error {
	event := Event{
		Type: eventType,
		Data: data,
	}
	body, err := json.Marshal(event)
	if err != nil {
		return err
	}
	_, err = http.Post("http://localhost:4005/events", "application/json", bytes.NewBuffer(body))
	return err
}
