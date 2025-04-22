package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gojektech/heimdall/v6"

	"github.com/gojektech/heimdall/v6/httpclient"
)

const (
	// DefaultMaxRetry is number of maximum retry when the request was fail.
	DefaultMaxRetry = 5

	// HTTPMethodGET represent http get method.
	HTTPMethodGET HTTPMethod = iota

	// HTTPMethodPOST represent http post method.
	HTTPMethodPOST

	// HTTPMethodPUT represent http put method.
	HTTPMethodPUT

	// HTTPMethodDELETE represent http delete method.
	HTTPMethodDELETE
)

// HTTPMethod represent method of http request (get/post/put/delete) in a number
type HTTPMethod int

// Request is a struct represent dependencies needed to be initialize.
type Request struct {
	client *httpclient.Client
	Method HTTPMethod
	Data   RequestData
}

// RequestData is a struct to sets the request data.
type RequestData struct {
	URL    string
	Body   interface{}
	Header http.Header
}

// NewHTTPRequest will initialize the HTTPRequest.
func NewHTTPRequest(options ...Option) *Request {
	linearRetry := heimdall.NewRetrierFunc(func(retry int) time.Duration {
		if retry <= 0 {
			return 0 * time.Millisecond
		}
		return time.Duration(retry) * time.Millisecond
	})

	timeout := 30000 * time.Millisecond
	client := httpclient.NewClient(
		httpclient.WithHTTPTimeout(timeout),
		httpclient.WithRetrier(linearRetry),
		httpclient.WithRetryCount(DefaultMaxRetry),
	)

	httpRequest := &Request{
		client: client,
	}

	for _, opt := range options {
		opt(httpRequest)
	}

	return httpRequest
}

// Exec is a method uses to execute the http request.
func (hr *Request) Exec(ctx context.Context) (*http.Response, error) {
	var response *http.Response
	var err error

	switch hr.Method {
	case HTTPMethodGET:
		response, err = hr.client.Get(hr.Data.URL, hr.Data.Header)

	case HTTPMethodPOST:

		response, err = hr.client.Post(hr.Data.URL, hr.TransformRequestBody(hr.Data.Body), hr.Data.Header)

	case HTTPMethodPUT:
		response, err = hr.client.Put(hr.Data.URL, hr.TransformRequestBody(hr.Data.Body), hr.Data.Header)

	case HTTPMethodDELETE:
		response, err = hr.client.Delete(hr.Data.URL, hr.Data.Header)
	}

	if err != nil {
		log.Printf("Unable to perform HTTP request, err: %s", err)

		return nil, err
	}

	return response, nil
}

// TransformRequestBody is a method uses to transform the request body into io.Reader.
func (hr *Request) TransformRequestBody(body interface{}) io.Reader {
	if body == nil {
		return bytes.NewReader(nil)
	}
	if b, ok := body.([]byte); ok {
		return bytes.NewReader(b)
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Printf("Unable to marshal request body, err: %v", err)
	}

	return bytes.NewReader(bodyBytes)
}

// ReadResponseBody is a method uses to read http.Response and return as bytes.
func (hr *Request) ReadResponseBody(response *http.Response) []byte {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Unable to read response body: %s", err)
	}

	return body
}
