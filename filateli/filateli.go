package filateli

import (
	"context"

	"github.com/avrebarra/filateli/templater"
)

// Filateli defines filateli instance
type Filateli interface {
	Build(ctx context.Context, in InputBuild) (doc []byte, err error)
}

// **

// InputBuild defines input for build method
type InputBuild struct {
	CollectionRaw []byte
	Templater     templater.Templater
}

// **
