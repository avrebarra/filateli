package main

import (
	"github.com/avrebarra/filateli/cmd"
)

//go:generate qtc

func main() {
	cmd.Initialize()
	cmd.Run()
}
