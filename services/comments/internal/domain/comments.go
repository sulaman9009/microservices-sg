package domain

type Comment struct {
	Id      string `json:"id"`
	Content string `json:"content"`
}

type CreateCommentReq struct {
	Content string `json:"content"`
}

type GetCommentsResp struct {
	Comments []*Comment `json:"comments"`
}
