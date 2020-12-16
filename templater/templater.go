package templater

import (
	"context"

	"github.com/avrebarra/filateli/postman"
)

// Templater is defines collection template objects
type Templater interface {
	Apply(ctx context.Context, coll postman.Collection) (doc []byte, err error)
}
