package main

func (s *Server) routes() {
	s.router.Get("/", s.handleHome())
	s.router.Get("/posts", s.handlePosts())
}
