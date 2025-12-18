package domain

type PostCreatedEventData struct {
	Id    string
	Title string
}

type CommentCreatedEventData struct {
	PostId  string
	Id      string
	Content string
	Status  string
}

type CommentUpdatedEvent struct {
	CommentCreatedEventData
}
