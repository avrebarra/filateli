package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/avrebarra/filateli/filateli"
	"github.com/avrebarra/filateli/templater"
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
	ctx := context.Background()

	// load target file
	fraw, err := ReadFile(c.config.SourcePath)
	if err != nil {
		return
	}

	// make templater
	tpl, err := templater.NewSimple(templater.ConfigSimple{})
	if err != nil {
		return
	}

	// execute
	f := filateli.New(filateli.Config{})
	docbts, err := f.Build(ctx, filateli.InputBuild{
		CollectionRaw: fraw,
		Templater:     tpl,
	})
	if err != nil {
		return
	}

	// write to target file
	err = WriteFile(c.config.OutputPath, docbts, false)
	if err != nil {
		return
	}

	c.Log(fmt.Sprintf("collection %s written to %s", c.config.SourcePath, c.config.OutputPath))

	return nil
}

func ReadFile(fpath string) (content []byte, err error) {
	content, err = ioutil.ReadFile(fpath)
	if err != nil {
		return
	}
	return
}

func CheckFile(fpath string) (exist bool, err error) {
	if _, err = os.Stat(fpath); os.IsNotExist(err) {
		return
	}

	exist = true
	return
}

func WriteFile(fpath string, content []byte, force bool) (err error) {
	if ok, _ := CheckFile(fpath); ok && !force {
		err = fmt.Errorf("file exist")
		return
	}

	os.MkdirAll(filepath.Dir(fpath), 0755)
	err = ioutil.WriteFile(fpath, content, 0644)
	if err != nil {
		return
	}
	return
}
