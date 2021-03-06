// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "HelloServerlessGoa": Application Controllers
//
// Command:
// $ goagen
// --design=github.com/tleyden/serverless-forms/design
// --out=$(GOPATH)/src/github.com/tleyden/serverless-forms/goa-generated
// --version=v1.2.0-dirty

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// HelloController is the controller interface for the Hello actions.
type HelloController interface {
	goa.Muxer
	Show(*ShowHelloContext) error
}

// MountHelloController "mounts" a Hello resource controller on the given service.
func MountHelloController(service *goa.Service, ctrl HelloController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowHelloContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/hello/:whatToSay", ctrl.MuxHandler("show", h, nil))
	service.LogInfo("mount", "ctrl", "Hello", "action", "Show", "route", "GET /hello/:whatToSay")
}
