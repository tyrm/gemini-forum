package gemini

import "regexp"

var reValidUsername = regexp.MustCompile(`^[a-zA-Z0-9_\-\.]+$`)
