package filateli

import (
	"bytes"
	"log"
	"text/template"

	"github.com/avrebarra/filateli/pkg/postman"
)

type Config struct {
	BootstrapJS  string
	BootstrapCSS string
	IndexHTML    string
	JqueryJS     string
	ScriptsJS    string
	StylesCSS    string
	ExtraCSS     string

	IndexMarkdown        string
	MarkdownHTML         string
	GithubMarkdownMinCSS string
}

type DefaultService struct {
	Config
}

func New(conf Config) Service {
	return &DefaultService{conf}
}

func (s *DefaultService) ConvertToHTML(c postman.Collection, envCollection postman.Environment) (buf *bytes.Buffer, err error) {
	// populate envCollection with collection variables
	if len(c.Variables) > 0 {
		envCollection.SetCollectionVariables(c.Variables)
	}

	tm := template.New("main")
	tm.Delims("@{{", "}}@")
	tm.Funcs(template.FuncMap{
		"html":            htmlTemplate,
		"css":             cssTemplate,
		"js":              jsTemplate,
		"eHTML":           eHTML,
		"snake":           snake,
		"addOne":          addOne,
		"color":           color,
		"trimQueryParams": trimQueryParams,
		"date_time":       dateTime,
		"markdown":        markdown,
		"e":               e,
	})
	t, err := tm.Parse(s.IndexHTML)
	if err != nil {
		log.Fatal(err)
	}

	data := struct {
		Assets Config
		Data   postman.Collection
	}{
		Assets: s.Config,
		Data:   c,
	}
	buf = new(bytes.Buffer)
	if err := t.Execute(buf, data); err != nil {
		log.Fatal(err)
	}

	return
}

func (s *DefaultService) ConvertToMarkdown(c postman.Collection, envCollection postman.Environment) (buf *bytes.Buffer, err error) {
	// populate envCollection with collection variables
	if len(c.Variables) > 0 {
		envCollection.SetCollectionVariables(c.Variables)
	}

	tm := template.New("main")
	tm.Delims("@{{", "}}@")
	tm.Funcs(template.FuncMap{
		"snake":           snake,
		"addOne":          addOne,
		"trim":            trim,
		"lower":           lower,
		"upper":           upper,
		"glink":           githubLink,
		"glinkInc":        githubLinkIncrementer,
		"merge":           merge,
		"roman":           roman,
		"date_time":       dateTime,
		"trimQueryParams": trimQueryParams,
		"e":               e,
	})
	t, err := tm.Parse(s.IndexMarkdown)
	if err != nil {
		log.Fatal(err)
	}
	data := struct {
		Data postman.Collection
	}{
		Data: c,
	}
	buf = new(bytes.Buffer)
	if err := t.Execute(buf, data); err != nil {
		log.Fatal(err)
	}

	return
}

func (s *DefaultService) ConvertToMarkdownHTML(c postman.Collection, envCollection postman.Environment) (buf *bytes.Buffer, err error) {
	// populate envCollection with collection variables
	if len(c.Variables) > 0 {
		envCollection.SetCollectionVariables(c.Variables)
	}

	tm := template.New("main")
	tm.Delims("@{{", "}}@")
	tm.Funcs(template.FuncMap{
		"html":            htmlTemplate,
		"css":             cssTemplate,
		"js":              jsTemplate,
		"eHTML":           eHTML,
		"snake":           snake,
		"addOne":          addOne,
		"color":           color,
		"trimQueryParams": trimQueryParams,
		"date_time":       dateTime,
		"markdown":        markdown,
		"e":               e,
	})

	t, err := tm.Parse(s.MarkdownHTML)
	if err != nil {
		log.Fatal(err)
	}

	buf, err = s.ConvertToMarkdown(c, envCollection)
	mdHTML := markdown(buf.String())

	data := struct {
		Assets       Config
		Data         postman.Collection
		MarkdownHTML string
	}{
		Assets:       s.Config,
		Data:         c,
		MarkdownHTML: mdHTML,
	}
	buf = new(bytes.Buffer)
	if err := t.Execute(buf, data); err != nil {
		log.Fatal(err)
	}
	return
}
