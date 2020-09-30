//go:generate go run generate_asset.go

package main

import (
	"github.com/avrebarra/filateli/cmd"
)

func main() {
	cmd.Execute()
}
