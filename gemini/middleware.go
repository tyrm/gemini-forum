package gemini

import "github.com/pitr/gig"

func (s *Server) middleware(next gig.HandlerFunc) gig.HandlerFunc {
	return func(c gig.Context) error {
		if c.Certificate() != nil {
			logger.Tracef("mid: looking up user for cert-hash %s", c.CertHash())
			user, err := s.db.ReadUserByCertHash(c.CertHash())
			if err != nil {
				c.Error(err)
				return err
			}
			if user != nil {
				c.Set("user", user)
			}
		}

		err := next(c)
		if err != nil {
			c.Error(err)
		}

		return err
	}
}
