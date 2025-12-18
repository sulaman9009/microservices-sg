package domain

type Comment struct {
	Id      string `json:"id"`
	Content string `json:"content"`
	Status  string `json:"status"`
}

type Post struct {
	Id       string     `json:"id"`
	Title    string     `json:"title"`
	Comments []*Comment `json:"comments"`
}

type GetPostsResponse struct {
	Posts []*Post `json:"posts"`
}
