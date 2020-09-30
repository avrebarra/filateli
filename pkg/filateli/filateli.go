package filateli

import (
	"bytes"

	"github.com/avrebarra/filateli/pkg/postman"
)

type Service interface {
	ConvertToHTML(c postman.Collection, lite bool) (buf *bytes.Buffer, err error)
	ConvertToMarkdown(c postman.Collection) (buf *bytes.Buffer, err error)
	ConvertToMarkdownHTML(c postman.Collection) (buf *bytes.Buffer, err error)
}
