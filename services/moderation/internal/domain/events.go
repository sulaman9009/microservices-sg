package domain

type CommentCreatedEventData struct {
	PostId  string `json:"postId"`
	Id      string `json:"id"`
	Content string `json:"content"`
	Status  string `json:"status"`
}
