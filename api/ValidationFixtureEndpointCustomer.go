package api

import (
	"bytes"
	"context"

	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"strings"

	"github.com/google/go-querystring/query"
)

func (c *Client) ValidationFixtureEndpointCustomerDelete(ctx context.Context, params *ValidationFixtureEndpointCustomerDeleteParameters, body *ValidationFixtureEndpointCustomerDeleteRequestBody) ([]byte, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	var payload io.Reader = nil
	var vals url.Values
	var err error
	if params != nil {
		vals, err = query.Values(params)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("(c *Client)  :: nil params")
	}
	pathPattern := "/endpoint/{customer}"
	pathPattern = strings.Replace(pathPattern, `{customer}`, params.Customer, -1)
	URL := c.createUrl(pathPattern, vals)
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, URL.String(), payload)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", c.cfg.ContentType) // TODO this feels hacky
	if c.auth != nil {
		err = c.auth(req)
		if err != nil {
			return nil, err
		}
	}

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode < 200 || res.StatusCode > 399 {
		var body []byte
		body, err := io.ReadAll(res.Body)
		if err == nil { // Close body if it is not nil and is read successfully
			defer res.Body.Close() // TODO figure out how to bubble this error up for checking
		}
		return nil, fmt.Errorf("status code check failed :: url %s :: returned %d :: body %v", URL.String(), res.StatusCode, string(body))
	}

	defer res.Body.Close() // TODO figure out how to bubble this error up for checking

	responseBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return responseBytes, nil
}
func (c *Client) ValidationFixtureEndpointCustomerGet(ctx context.Context, params *ValidationFixtureEndpointCustomerGetParameters, body *ValidationFixtureEndpointCustomerGetRequestBody) ([]byte, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	var payload io.Reader = nil
	var vals url.Values
	var err error
	if params != nil {
		vals, err = query.Values(params)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("(c *Client)  :: nil params")
	}
	pathPattern := "/endpoint/{customer}"
	pathPattern = strings.Replace(pathPattern, `{customer}`, params.Customer, -1)
	URL := c.createUrl(pathPattern, vals)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL.String(), payload)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", c.cfg.ContentType) // TODO this feels hacky
	if c.auth != nil {
		err = c.auth(req)
		if err != nil {
			return nil, err
		}
	}

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode < 200 || res.StatusCode > 399 {
		var body []byte
		body, err := io.ReadAll(res.Body)
		if err == nil { // Close body if it is not nil and is read successfully
			defer res.Body.Close() // TODO figure out how to bubble this error up for checking
		}
		return nil, fmt.Errorf("status code check failed :: url %s :: returned %d :: body %v", URL.String(), res.StatusCode, string(body))
	}

	defer res.Body.Close() // TODO figure out how to bubble this error up for checking

	responseBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return responseBytes, nil
}
func (c *Client) ValidationFixtureEndpointCustomerPatch(ctx context.Context, params *ValidationFixtureEndpointCustomerPatchParameters, body *ValidationFixtureEndpointCustomerPatchRequestBody) ([]byte, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	var payload io.Reader = nil
	var vals url.Values
	var err error
	if params != nil {
		vals, err = query.Values(params)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("(c *Client)  :: nil params")
	}
	pathPattern := "/endpoint/{customer}"
	pathPattern = strings.Replace(pathPattern, `{customer}`, params.Customer, -1)
	URL := c.createUrl(pathPattern, vals)
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, URL.String(), payload)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", c.cfg.ContentType) // TODO this feels hacky
	if c.auth != nil {
		err = c.auth(req)
		if err != nil {
			return nil, err
		}
	}

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode < 200 || res.StatusCode > 399 {
		var body []byte
		body, err := io.ReadAll(res.Body)
		if err == nil { // Close body if it is not nil and is read successfully
			defer res.Body.Close() // TODO figure out how to bubble this error up for checking
		}
		return nil, fmt.Errorf("status code check failed :: url %s :: returned %d :: body %v", URL.String(), res.StatusCode, string(body))
	}

	defer res.Body.Close() // TODO figure out how to bubble this error up for checking

	responseBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return responseBytes, nil
}
func (c *Client) ValidationFixtureEndpointCustomerPost(ctx context.Context, params *ValidationFixtureEndpointCustomerPostParameters, body *ValidationFixtureEndpointCustomerPostRequestBody) ([]byte, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	var payload io.Reader = nil
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil { // This serves as our validation, for now
			return nil, err
		}
		payload = bytes.NewReader(data)
	} else {
		return nil, errors.New("(c *Client)  :: post operation with nil body")
	}
	var vals url.Values
	var err error
	if params != nil {
		vals, err = query.Values(params)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("(c *Client)  :: nil params")
	}
	pathPattern := "/endpoint/{customer}"
	pathPattern = strings.Replace(pathPattern, `{customer}`, params.Customer, -1)
	URL := c.createUrl(pathPattern, vals)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, URL.String(), payload)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", c.cfg.ContentType) // TODO this feels hacky
	if c.auth != nil {
		err = c.auth(req)
		if err != nil {
			return nil, err
		}
	}

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode < 200 || res.StatusCode > 399 {
		var body []byte
		body, err := io.ReadAll(res.Body)
		if err == nil { // Close body if it is not nil and is read successfully
			defer res.Body.Close() // TODO figure out how to bubble this error up for checking
		}
		return nil, fmt.Errorf("status code check failed :: url %s :: returned %d :: body %v", URL.String(), res.StatusCode, string(body))
	}

	defer res.Body.Close() // TODO figure out how to bubble this error up for checking

	responseBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return responseBytes, nil
}
