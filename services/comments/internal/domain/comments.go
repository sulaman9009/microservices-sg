package domain

type Comment struct {
	Id      string `json:"id"`
	Content string `json:"content"`
	Status  string `json:"status"`
}

type CreateCommentReq struct {
	Content string `json:"content"`
}

type GetCommentsResp struct {
	Comments []*Comment `json:"comments"`
}

type CommentWithPostId struct {
	PostId string `json:"postId"`
	Comment
}
