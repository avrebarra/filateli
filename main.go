package main

import (
	"fmt"

	"github.com/avrebarra/filateli/cmd"
)

//go:generate qtc

func main() {
	cmd.Initialize()
	err := cmd.Run()
	if err != nil {
		fmt.Println("unexpected error:", err.Error())
	}
}
