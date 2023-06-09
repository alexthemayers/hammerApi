package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-playground/validator/v10"
)

// ClientService is the interface for Client methods
type ClientService interface {
	ValidationFixtureEndpointCustomerDelete(context.Context, *ValidationFixtureEndpointCustomerDeleteParameters, *ValidationFixtureEndpointCustomerDeleteRequestBody) ([]byte, error)
	ValidationFixtureEndpointCustomerGet(context.Context, *ValidationFixtureEndpointCustomerGetParameters, *ValidationFixtureEndpointCustomerGetRequestBody) ([]byte, error)
	ValidationFixtureEndpointCustomerPatch(context.Context, *ValidationFixtureEndpointCustomerPatchParameters, *ValidationFixtureEndpointCustomerPatchRequestBody) ([]byte, error)
	ValidationFixtureEndpointCustomerPost(context.Context, *ValidationFixtureEndpointCustomerPostParameters, *ValidationFixtureEndpointCustomerPostRequestBody) ([]byte, error)
}

var _ ClientService = &Client{}

// AuthFunc needs to be passed in as a closure at client generation and allows for authentication to be applied to any
// request being made. It should contain any logic and data needed to apply and maintain authentication with a server.
type AuthFunc func(*http.Request) error

type Config struct {
	Port        string `validate:"required,hostname_port"`
	Host        string `validate:"required,hostname"` // Hostname RFC 952
	Protocol    string `validate:"required"`
	Token       string
	ContentType string `validate:"required"`

	base string
}

type Client struct {
	*http.Client

	cfg  *Config
	auth AuthFunc
}

func (c *Client) createUrl(pathPattern string, vals url.Values) url.URL {
	return url.URL{
		Scheme:   c.cfg.Protocol,
		Host:     c.cfg.Host + c.cfg.Port,
		Path:     pathPattern,
		RawQuery: vals.Encode(),
	}
}

func NewHTTPClient(cfg *Config, af AuthFunc) (*Client, error) {
	if cfg == nil {
		return nil, errors.New("cannot start client without config")
	}
	validate := validator.New()
	if err := validate.Struct(cfg); err != nil {
		return nil, err
	}
	if strings.Contains(cfg.Host, "://") {
		cfg.Host = strings.Split(cfg.Host, "://")[1]
		cfg.base = fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	} else {
		cfg.base = fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	}
	return &Client{
		Client: http.DefaultClient,
		cfg:    cfg,
		auth:   af,
	}, nil
}
