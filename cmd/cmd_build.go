package cmd

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

type ConfigCommandBuild struct {
	Quiet      bool
	SourcePath string `validate:"required,ne="`
	OutputPath string `validate:"required,ne="`
}

type CommandBuild struct {
	config ConfigCommandBuild
}

func NewCommandBuild(cfg ConfigCommandBuild) CommandBuild {
	if err := validator.New().Struct(cfg); err != nil {
		panic(err)
	}
	return CommandBuild{config: cfg}
}

func (c CommandBuild) Log(msg string) {
	if !c.config.Quiet {
		fmt.Println(msg)
	}
}

func (c CommandBuild) Run() (err error) {
	return nil
}
