package postman

import (
	"fmt"
	"os"
	"testing"
)

func TestDocumentation_Open(t *testing.T) {
	file, err := os.Open("collection.json")
	if err != nil {
		t.Error(err)
	}
	d := &Collection{}
	err = d.ParseFrom(file)
	if err != nil {
		t.Error(err)
	}

	for _, c := range d.Items {
		for _, i := range c.Items {
			for _, r := range i.Responses {
				fmt.Println("x", r)
			}
		}
	}
}
