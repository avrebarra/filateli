// +build generatestorage

package main

import (
	"log"
	"net/http"

	"github.com/shurcooL/vfsgen"
)

func main() {
	// from assets
	var fs http.FileSystem = http.Dir("./assets/")
	err := vfsgen.Generate(fs, vfsgen.Options{
		Filename:     "./storage/assets/store.go",
		PackageName:  "assets",
		VariableName: "AssetFS",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
