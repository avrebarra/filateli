package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/avrebarra/filateli/pkg/filateli"
	"github.com/avrebarra/filateli/storage"
	"github.com/spf13/cobra"
)

var (
	filatelisvc filateli.Service
)

var CommandRoot = &cobra.Command{
	Use:   "filateli",
	Short: "Generate documentation from Postman JSON collection",
	Long: strings.TrimSpace(`	 
.		                                                           
,7MM"""YMM   db  ,7MM             mm             ,7MM    db  
  MM    ,7         MM             MM               MM        
  MM   d   ,7MM    MM   ,6"Yb.  mmMMmm   .gP"Ya    MM  ,7MM  
  MM""MM     MM    MM  8)   MM    MM    ,M'   Yb   MM    MM  
  MM   Y     MM    MM   ,pm9MM    MM    8M""""""   MM    MM  
  MM         MM    MM  8M   MM    MM    YM.    ,   MM    MM  
.JMML.     .JMML..JMML.,Moo9^Yo.  ,Mbmo  ,Mbmmd' .JMML..JMML.
                                                                                                                         

Generate API documentation from Postman JSON collection
For more info visit: https://github.com/avrebarra/filateli
`),
}

func init() {
	// init dependencies
	filatelisvc = filateli.New(filateli.Config{
		IndexHTML:     readFileFromStorage("templates/html/index.html"),
		IndexHTMLLite: readFileFromStorage("templates/html/index-lite.html"),
		BootstrapCSS:  readFileFromStorage("templates/html/bootstrap.min.css"),
		BootstrapJS:   readFileFromStorage("templates/html/bootstrap.min.js"),
		JqueryJS:      readFileFromStorage("templates/html/jquery.min.js"),
		ScriptsJS:     readFileFromStorage("templates/html/scripts.js"),
		StylesCSS:     readFileFromStorage("templates/html/styles.css"),

		GithubMarkdownMinCSS: readFileFromStorage("templates/markdown/github-markdown.min.css"),
		IndexMarkdown:        readFileFromStorage("templates/markdown/index.md"),
		MarkdownHTML:         readFileFromStorage("templates/markdown/markdown.html"),
	})

	// register child commands
	CommandRoot.AddCommand(CommandConvert)
}

// Execute the root command
func Execute() error {
	return CommandRoot.Execute()
}

func handle(message string, err error) {
	if err != nil {
		err = fmt.Errorf("%s: %w", message, err)
		log.Fatal(err)
	}
}

func readFileFromStorage(a string) string {
	fp, err := storage.AssetFS.Open(a)
	if err != nil {
		log.Fatal(err)
	}
	bs, err := ioutil.ReadAll(fp)
	if err != nil {
		log.Fatal(err)
	}
	return string(bs)
}
