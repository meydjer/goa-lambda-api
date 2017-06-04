// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "HelloServerlessGoa": hello Resource Client
//
// Command:
// $ goagen
// --design=github.com/tleyden/serverless-forms/design
// --out=$(GOPATH)/src/github.com/tleyden/serverless-forms/goa-generated
// --version=v1.2.0-dirty

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// ShowHelloPath computes a request path to the show action of hello.
func ShowHelloPath(whatToSay string) string {
	param0 := whatToSay

	return fmt.Sprintf("/hello/%s", param0)
}

// Say Hello
func (c *Client) ShowHello(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowHelloRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowHelloRequest create the request corresponding to the show action endpoint of the hello resource.
func (c *Client) NewShowHelloRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
