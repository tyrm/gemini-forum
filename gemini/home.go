package gemini

import (
	"github.com/pitr/gig"
)

type homeTemplate struct {
	templateCommon
}

func (s *Server) handleHome(c gig.Context) error {
	tmplVars := homeTemplate{}
	err := initTemplate(c, &tmplVars)
	if err != nil {
		return err
	}

	tmplVars.PageTitle = "Forum Home"

	return c.Render("home", &tmplVars)
}
