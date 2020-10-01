//go:generate go run generate_storage.go

package main

import (
	"github.com/avrebarra/filateli/cmd"
)

func main() {
	cmd.Execute()
}
