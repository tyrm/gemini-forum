package gemini

import (
	"crypto/tls"
	"github.com/markbates/pkger"
	"github.com/pitr/gig"
	"github.com/tyrm/gemini-forum/config"
	"github.com/tyrm/gemini-forum/db"
	"github.com/tyrm/gemini-forum/kv"
)

// Server is a GraphQL api server
type Server struct {
	db db.DB
	kv kv.KV

	server *gig.Gig
}

// NewServer will create a new GraphQL server
func NewServer(cfg *config.Config, d db.DB, k kv.KV) (*Server, error) {
	server := Server{
		db:     d,
		kv:     k,
		server: gig.Default(),
	}

	// High Security
	server.server.TLSConfig.MinVersion = tls.VersionTLS13

	// Load Templates
	templateDir := pkger.Include("/gemini/templates")
	t, err := compileTemplates(templateDir)
	if err != nil {
		return nil, err
	}
	server.server.Renderer = &templateRenderer{t}

	// Create Routes
	server.server.Use(server.middleware)

	server.server.Handle("/", server.handleHome)
	server.server.Handle("/register", server.handleRegister)

	return &server, nil
}

// Run starts the gemini server.
func (s *Server) Run() error {
	return s.server.Run("server.crt", "server.key")
}
