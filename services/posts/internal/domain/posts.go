package domain

type Post struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

type CreatePost struct {
	Title string `json:"title"`
}

type GetPostsResponse struct {
	Posts []*Post `json:"posts"`
}
