package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/avrebarra/filateli/pkg/filateli"
	storeassets "github.com/avrebarra/filateli/storage/assets"
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
		BuiltinTemplateHTML:         readFileFromStorage("templates/index.html.filatpl"),
		BuiltinTemplateHTMLLite:     readFileFromStorage("templates/index-lite.html.filatpl"),
		BuiltinTemplateMarkdown:     readFileFromStorage("templates/markdown.md.filatpl"),
		BuiltinTemplateMarkdownHTML: readFileFromStorage("templates/markdown-web.html.filatpl"),
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
	fp, err := storeassets.AssetFS.Open(a)
	if err != nil {
		log.Fatal(err)
	}
	bs, err := ioutil.ReadAll(fp)
	if err != nil {
		log.Fatal(err)
	}
	return string(bs)
}
