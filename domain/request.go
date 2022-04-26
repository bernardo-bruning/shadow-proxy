package domain

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	neturl "net/url"

	"cloud.google.com/go/pubsub"
)

type Request struct {
	Method  string              `json:"method"`
	Url     string              `json:"url"`
	Headers map[string][]string `json:"headers"`
	Body    []byte              `json:"body"`
}

func FromHttpRequest(r *http.Request) (*Request, error) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return nil, err
	}

	req := Request{
		Method:  r.Method,
		Url:     r.URL.String(),
		Headers: r.Header,
		Body:    body,
	}

	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return &req, nil
}

func FromJson(data []byte) *Request {
	var request Request
	json.Unmarshal(data, &request)
	return &request
}

func FromMessage(message *pubsub.Message) *Request {
	return FromJson(message.Data)
}

func (req *Request) ToHttpRequest(u string) (*http.Request, error) {
	url, err := neturl.Parse(u + req.Url)
	if err != nil {
		return nil, err
	}

	return &http.Request{
		Method: req.Method,
		URL:    url,
		Header: req.Headers,
		Body:   io.NopCloser(bytes.NewBuffer(req.Body)),
	}, nil
}

func (req *Request) ToJson() ([]byte, error) {
	data, err := json.Marshal(req)
	return data, err
}

func (req *Request) ToMessage() (*pubsub.Message, error) {
	data, err := req.ToJson()
	if err != nil {
		return nil, err
	}

	return &pubsub.Message{Data: data}, nil
}
