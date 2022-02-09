package gemini

import (
	"github.com/markbates/pkger"
	"github.com/pitr/gig"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c gig.Context) error {
	// Execute named template with data
	return t.templates.ExecuteTemplate(w, name, data)
}

func compileTemplates(dir string) (*template.Template, error) {
	tpl := template.New("")

	err := pkger.Walk(dir, func(path string, info os.FileInfo, _ error) error {
		if info.IsDir() || !strings.HasSuffix(path, ".gogmi") {
			return nil
		}
		f, err := pkger.Open(path)
		if err != nil {
			logger.Errorf("could not open pkger path %s: %s", path, err.Error())
			return err
		}
		// Now read it.
		sl, err := ioutil.ReadAll(f)
		if err != nil {
			logger.Errorf("could not read pkger file %s: %s", path, err.Error())
		}

		// It can now be parsed as a string.
		_, err = tpl.Parse(string(sl))
		if err != nil {
			logger.Errorf("could not open parse template %s: %s", path, err.Error())
			return err
		}

		return nil
	})

	return tpl, err
}
