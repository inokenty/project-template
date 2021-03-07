package icanhazip

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

var (
	ErrNoClient = errors.New("no client")
)

type Config struct {
	Endpoint string
}

type resource struct {
	config *Config
}

func NewResource(c *Config) Resource {
	return &resource{config: c}
}

type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

type GetIPInput struct {
	Client Doer
}

type GetIPReply struct {
	IP string
}

func (r *resource) GetIP(ctx context.Context, input GetIPInput) (*GetIPReply, error) {
	if input.Client == nil {
		return nil, ErrNoClient
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, r.config.Endpoint, nil)
	if err != nil {
		return nil, errors.Wrap(err, "http.NewRequestWithContext")
	}

	resp, err := input.Client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "client.Do")
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "ioutil.ReadAll")
	}

	ip := bytes.TrimSpace(bodyBytes)

	return &GetIPReply{IP: string(ip)}, nil
}
