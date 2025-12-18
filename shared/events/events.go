package events

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Event struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

const (
	TypePostCreated      = "PostCreated"
	TypeCommentCreated   = "CommentCreated"
	TypeCommentModerated = "CommentModerated"
	TypeCommentUpdated   = "CommentUpdated"
)

func EmitEvent(eventType string, data any) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	event := Event{
		Type: eventType,
		Data: jsonData,
	}
	body, err := json.Marshal(event)
	if err != nil {
		return err
	}
	_, err = http.Post("http://localhost:4005/events", "application/json", bytes.NewBuffer(body))
	return err
}
