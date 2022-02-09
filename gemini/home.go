package gemini

import "github.com/pitr/gig"

func (s *Server) handleHome(c gig.Context) error {
	return c.Render("home", nil)
}
