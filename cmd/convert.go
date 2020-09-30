package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/avrebarra/filateli/pkg/postman"
	"github.com/spf13/cobra"
)

var (
	argsTargetFile string
	argsOut        string
	argsEnv        string
	argsMode       string
)

var CommandConvert = &cobra.Command{
	Use:   "convert",
	Short: "Convert postman collection to html/markdown documentation",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			handle("target file required", fmt.Errorf("no target file defined"))
		}

		argsTargetFile = args[0]

		env := &postman.Environment{}
		if argsEnv != "" {
			if _, err := os.Stat(argsEnv); os.IsNotExist(err) {
				handle("invalid environment file path", err)
			}
			f, err := os.Open(argsEnv)
			if err != nil {
				handle("unable to open file", err)
			}
			if err := env.ParseFrom(f); err != nil {
				handle("unable to parse environment file", err)
			}
		}

		switch argsMode {
		case "html":
			convertToHTML()

		case "litehtml":
			convertToLiteHTML()

		case "markdown":
			convertToMarkdown()

		case "webmarkdown":
			convertToMarkdownHTML()

		default:
			log.Fatalf("unknown mode: %s", argsMode)
		}
	},
}

func init() {
	CommandConvert.PersistentFlags().StringVarP(&argsOut, "out", "o", "", "output file path")
	CommandConvert.PersistentFlags().StringVarP(&argsEnv, "env", "e", "", "environment file path")
	CommandConvert.PersistentFlags().StringVarP(&argsMode, "mode", "m", "html", "convertion mode [html, litehtml, webmarkdown, markdown]")
}

func convertToMarkdown() {
	buffile, err := readFile(argsTargetFile)
	handle("cannot read target file", err)

	collection := postman.Collection{}
	err = collection.ParseFrom(buffile)
	handle("cannot parse target file", err)

	bufmd, err := filatelisvc.ConvertToMarkdown(collection)
	handle("conversion to markdown failed", err)

	lines := strings.Split(bufmd.String(), "\n")
	for i, l := range lines {
		if strings.HasPrefix(l, "<!---") && strings.HasSuffix(l, "-->") {
			lines = append(lines[:i], lines[i+1:]...)
		}
	}

	// contents := strings.Join(lines, "\n")
	var contents string
	var ws int
	for _, l := range lines {
		if l == "" {
			ws++
		} else {
			ws = 0
		}
		if ws <= 3 {
			contents += "\n" + l
		}
	}

	err = ioutil.WriteFile(argsOut, []byte(contents), 0644)
	handle("failed writing file", err)

	log.Printf("markdown generated in %s", argsOut)
}

func convertToMarkdownHTML() {
	buffile, err := readFile(argsTargetFile)
	handle("cannot read target file", err)

	collection := postman.Collection{}
	err = collection.ParseFrom(buffile)
	handle("cannot parse target file", err)

	bufmd, err := filatelisvc.ConvertToMarkdownHTML(collection)
	handle("conversion to markdown failed", err)

	lines := strings.Split(bufmd.String(), "\n")
	for i, l := range lines {
		if strings.HasPrefix(l, "<!---") && strings.HasSuffix(l, "-->") {
			lines = append(lines[:i], lines[i+1:]...)
		}
	}

	// contents := strings.Join(lines, "\n")
	var contents string
	var ws int
	for _, l := range lines {
		if l == "" {
			ws++
		} else {
			ws = 0
		}
		if ws <= 3 {
			contents += "\n" + l
		}
	}

	err = ioutil.WriteFile(argsOut, []byte(contents), 0644)
	handle("failed writing file", err)

	log.Printf("markdown generated in %s", argsOut)
}

func convertToHTML() {
	buffile, err := readFile(argsTargetFile)
	handle("cannot read target file", err)

	collection := postman.Collection{}
	err = collection.ParseFrom(buffile)
	handle("cannot parse target file", err)

	bufmd, err := filatelisvc.ConvertToHTML(collection, false)
	handle("conversion to html failed", err)

	lines := strings.Split(bufmd.String(), "\n")
	for i, l := range lines {
		if strings.HasPrefix(l, "<!---") && strings.HasSuffix(l, "-->") {
			lines = append(lines[:i], lines[i+1:]...)
		}
	}

	// contents := strings.Join(lines, "\n")
	var contents string
	var ws int
	for _, l := range lines {
		if l == "" {
			ws++
		} else {
			ws = 0
		}
		if ws <= 3 {
			contents += "\n" + l
		}
	}

	err = ioutil.WriteFile(argsOut, []byte(contents), 0644)
	handle("failed writing file", err)

	log.Printf("html file generated in %s", argsOut)
}

func readFile(filepath string) (buf *bytes.Buffer, err error) {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return
	}

	buf = bytes.NewBuffer(content)
	return
}

func convertToLiteHTML() {
	buffile, err := readFile(argsTargetFile)
	handle("cannot read target file", err)

	collection := postman.Collection{}
	err = collection.ParseFrom(buffile)
	handle("cannot parse target file", err)

	bufmd, err := filatelisvc.ConvertToHTML(collection, true)
	handle("conversion to html failed", err)

	lines := strings.Split(bufmd.String(), "\n")
	for i, l := range lines {
		if strings.HasPrefix(l, "<!---") && strings.HasSuffix(l, "-->") {
			lines = append(lines[:i], lines[i+1:]...)
		}
	}

	// contents := strings.Join(lines, "\n")
	var contents string
	var ws int
	for _, l := range lines {
		if l == "" {
			ws++
		} else {
			ws = 0
		}
		if ws <= 3 {
			contents += "\n" + l
		}
	}

	err = ioutil.WriteFile(argsOut, []byte(contents), 0644)
	handle("failed writing file", err)

	log.Printf("html file generated in %s", argsOut)
}
