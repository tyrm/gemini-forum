package gemini

import (
	"github.com/pitr/gig"
	"github.com/tyrm/gemini-forum/config"
	"github.com/tyrm/gemini-forum/db"
)

// Server is a GraphQL api server
type Server struct {
	db     db.DB
	server *gig.Gig
}

// NewServer will create a new GraphQL server
func NewServer(cfg *config.Config, d db.DB) (*Server, error) {
	server := Server{
		db:     d,
		server: gig.Default(),
	}

	server.server.Handle("/user", func(c gig.Context) error {
		query, err := c.QueryString()
		if err != nil {
			return err
		}
		return c.Gemini("# Hello, %s!", query)
	})

	return &server, nil
}

// Run starts the gemini server.
func (s *Server) Run() error {
	return s.server.Run("server.crt", "server.key")
}
