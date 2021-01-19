package templater

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/avrebarra/filateli/postman"
	"github.com/avrebarra/filateli/templater/templates"
)

type ConfigSimple struct{}

type Simple struct {
	config ConfigSimple
}

func NewSimple(cfg ConfigSimple) (t Templater, err error) {
	// if err = validator.New().Struct(cfg); err != nil {
	// 	return
	// }
	return &Simple{config: cfg}, nil
}

func (e *Simple) Apply(ctx context.Context, coll postman.Collection) (doc []byte, err error) {
	payload, err := e.makeTemplatePayload(coll)
	if err != nil {
		return
	}

	docstr := templates.MakePostmanSimple(payload)
	doc = []byte(docstr)

	return
}

func (e *Simple) makeTemplatePayload(collection postman.Collection) (payload templates.PostmanSimplePayload, err error) {
	if collection.Info.Name == "" {
		err = fmt.Errorf("nullish input: cannot get info name")
		return
	}

	payload = templates.PostmanSimplePayload{
		Name: collection.Info.Name,
		RequestDirectories: []struct {
			Name     string
			Requests []templates.PostmanSimpleRequestMarkupDataV1
		}{},
		Description: collection.Info.Description,
	}

	for _, group := range collection.Items {
		reqdir := struct {
			Name     string
			Requests []templates.PostmanSimpleRequestMarkupDataV1
		}{
			Name:     group.Name,
			Requests: []templates.PostmanSimpleRequestMarkupDataV1{},
		}

		for _, request := range group.Items {
			reqmarkupdata := templates.PostmanSimpleRequestMarkupDataV1{
				Directory:          group.Name,
				Name:               request.Name,
				Description:        request.Request.Description,
				HTTPVerb:           request.Request.Method,
				URL:                request.Request.URL.Raw,
				ExampleRequestBody: request.Request.Body.Raw,

				Responses: []templates.PostmanSimpleExampleResponse{},
				CURL:      "", // will be added below

				QueryParams: []templates.PostmanSimpleRequestQueryParam{},
				URLParams:   []templates.PostmanSimpleRequestURLParam{},
			}

			for _, response := range request.Responses {
				reqmarkupdata.Responses = append(reqmarkupdata.Responses, templates.PostmanSimpleExampleResponse{
					Code:         response.Code,
					Name:         response.Name,
					Status:       response.Status,
					ResponseBody: response.Body,
				})
			}

			// add curl string
			curlstr := ""
			mockrequest, _ := http.NewRequest(request.Request.Method, request.Request.URL.Raw, strings.NewReader(request.Request.Body.Raw))
			if mockrequest != nil {
				scurl, _ := CURLFromRequest(mockrequest)
				curlstr = string(scurl)
			}
			reqmarkupdata.CURL = curlstr

			// register to list
			reqdir.Requests = append(reqdir.Requests, reqmarkupdata)
		}

		payload.RequestDirectories = append(payload.RequestDirectories, reqdir)
	}

	return
}
