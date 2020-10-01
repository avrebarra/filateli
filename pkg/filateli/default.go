package filateli

import (
	"bytes"
	"log"
	"text/template"

	"github.com/avrebarra/filateli/pkg/postman"
)

type Config struct {
	BuiltinTemplateHTML         string
	BuiltinTemplateHTMLLite     string
	BuiltinTemplateMarkdown     string
	BuiltinTemplateMarkdownHTML string
}

type DefaultService struct {
	Config
}

func New(conf Config) Service {
	return &DefaultService{conf}
}

func (s *DefaultService) ConvertToHTML(c postman.Collection, lite bool) (buf *bytes.Buffer, err error) {
	// populate envCollection with collection variables
	if len(c.Variables) > 0 {
		envCollection.SetCollectionVariables(c.Variables)
	}

	tm := template.New("main")
	tm.Delims("<{", "}>")
	tm.Funcs(template.FuncMap{
		"FilaSanitizeHTML":            HTMLSanitize,
		"FilaEscapeHTML":              HTMLEscape,
		"FilaChangeCaseSnake":         formatSnakeCase,
		"FilaIncrement":               incrementOne,
		"FilaTemplateHTMLColorOfVerb": colorOfVerb,
		"FilaURLTrimQueryParams":      formatURLTrimQueryParams,
		"FilaDateTime":                formatDateTime,
		"FilaBuildMarkdown":           buildMarkdown,
		"FilaEnv":                     envvar,
	})

	base := s.BuiltinTemplateHTML
	if lite {
		base = s.BuiltinTemplateHTMLLite
	}

	t, err := tm.Parse(base)
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

func (s *DefaultService) ConvertToMarkdown(c postman.Collection) (buf *bytes.Buffer, err error) {
	// populate envCollection with collection variables
	if len(c.Variables) > 0 {
		envCollection.SetCollectionVariables(c.Variables)
	}

	tm := template.New("main")
	tm.Delims("<{", "}>")
	tm.Funcs(template.FuncMap{
		"FilaChangeCaseSnake":       formatSnakeCase,
		"FilaDateTime":              formatDateTime,
		"FilaEnv":                   envvar,
		"FilaIncrement":             incrementOne,
		"FilaURLTrimQueryParams":    formatURLTrimQueryParams,
		"FilaGithubLinkIncrementer": githubLinkIncrementer,
		"FilaGetGithubLink":         githubLink,
		"FilaLower":                 lower,
		"FilaMerge":                 merge,
		"FilaRoman":                 roman,
		"FilaTrim":                  trim,
		"FilaUpper":                 upper,
	})
	t, err := tm.Parse(s.BuiltinTemplateMarkdown)
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

func (s *DefaultService) ConvertToMarkdownHTML(c postman.Collection) (buf *bytes.Buffer, err error) {
	// populate envCollection with collection variables
	if len(c.Variables) > 0 {
		envCollection.SetCollectionVariables(c.Variables)
	}

	tm := template.New("main")
	tm.Delims("<{", "}>")
	tm.Funcs(template.FuncMap{
		"FilaSanitizeHTML":            HTMLSanitize,
		"FilaEscapeHTML":              HTMLEscape,
		"FilaChangeCaseSnake":         formatSnakeCase,
		"FilaIncrement":               incrementOne,
		"FilaTemplateHTMLColorOfVerb": colorOfVerb,
		"FilaURLTrimQueryParams":      formatURLTrimQueryParams,
		"FilaDateTime":                formatDateTime,
		"FilaBuildMarkdown":           buildMarkdown,
		"FilaEnv":                     envvar,
	})

	t, err := tm.Parse(s.BuiltinTemplateMarkdownHTML)
	if err != nil {
		log.Fatal(err)
	}

	buf, err = s.ConvertToMarkdown(c)
	mdHTML := buildMarkdown(buf.String())

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
