package domain

type CommentModeratedEventData struct {
	PostId  string `json:"postId"`
	Id      string `json:"id"`
	Content string `json:"content"`
	Status  string `json:"status"`
}
