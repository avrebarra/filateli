package filateli

import (
	"bytes"
	"context"

	"github.com/avrebarra/filateli/postman"
	"gopkg.in/go-playground/validator.v9"
)

type Config struct {
}

type FilateliStruct struct {
	config Config
}

func New(cfg Config) Filateli {
	if err := validator.New().Struct(cfg); err != nil {
		panic(err)
	}
	return &FilateliStruct{config: cfg}
}

func (e *FilateliStruct) Build(ctx context.Context, in InputBuild) (docbts []byte, err error) {
	// parse collection
	co := postman.Collection{}
	err = co.ParseFrom(bytes.NewBuffer(in.CollectionRaw))
	if err != nil {
		return
	}

	// apply templater
	docbts, err = in.Templater.Apply(ctx, co)
	if err != nil {
		return
	}

	return
}
