package gemini

import (
	"fmt"
	"github.com/pitr/gig"
	"github.com/tyrm/gemini-forum/models"
)

type registerErrorTemplate struct {
	templateCommon

	ErrorTest string
}

func (s *Server) handleRegister(c gig.Context) error {
	// if user already exists return home
	if c.Get("user") != nil {
		return c.NoContent(gig.StatusRedirectTemporary, "/")
	}

	// check if cert provided
	if c.Certificate() == nil {
		return c.NoContent(gig.StatusClientCertificateRequired, "A certificate is required to register")
	}

	// get query
	newUsername, err := c.QueryString()
	if err != nil {
		logger.Errorf("query string: %s", err)
		return err
	}

	// if empty prompt for input
	if newUsername == "" {
		return c.NoContent(gig.StatusInput, "Please choose a username.")
	}

	// validate username
	if !reValidUsername.MatchString(newUsername) {
		tmplVars := registerErrorTemplate{}
		tmplVars.PageTitle = "Invalid Username"
		tmplVars.ErrorTest = fmt.Sprintf("The username '%s' contains invalid characters.\nAllowed characters are upper and lowecase alpha, numbers, dash, uderscore, and period.", newUsername)

		return c.Render("register_error", &tmplVars)
	}

	// check if username exists
	existingUser, err := s.db.ReadUserByUsername(newUsername)
	if err != nil {
		logger.Errorf("db: %s", err)
		return c.NoContent(gig.StatusCGIError, "database error")
	}
	if existingUser != nil {
		tmplVars := registerErrorTemplate{}
		tmplVars.PageTitle = "Username Taken"
		tmplVars.ErrorTest = fmt.Sprintf("The username '%s' has been taken.", newUsername)

		return c.Render("register_error", &tmplVars)
	}

	// create user
	newUser := models.User{
		CertHash: c.CertHash(),
		Username: newUsername,
	}

	// add user to database
	err = s.db.Create(&newUser)
	if err != nil {
		logger.Errorf("db: %s", err)
		return c.NoContent(gig.StatusCGIError, "database error")
	}

	return c.NoContent(gig.StatusRedirectTemporary, "/")
}
