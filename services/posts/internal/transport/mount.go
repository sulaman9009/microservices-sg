package transport

func (s *server) mount() {
	s.mux.GET("/posts", s.getPosts)
	s.mux.POST("/post", s.createPost)
	s.mux.POST("/events", s.eventHandler)
}
