package client

import (
	"context"
	"fmt"
	"strings"

	"go-vrf/src/nsxt"
)

const (
	XSRF_TOKEN = "X-XSRF-TOKEN"
)

type NSXTClient struct {
	cfg           Configuration
	client        *Client
	Context       context.Context
	common        Service
	defaultHeader map[string]string

	// Services
	ApiServiceNSXT *nsxt.ApiServiceNSXT
}

type Service struct {
	client *NSXTClient
}

type Configuration struct {
	BasePath string
	Username string
	Password string
	Insecure bool
}

type ErrorMessage struct {
	ModuleName   string         `json:"module_name,omitempty"`
	Message      string         `json:"error_message,omitempty"`
	RelatedError []RelatedError `json:"related_errors,omitempty"`
}

type RelatedError struct {
	HttpStatus   string `json:"httpStatus,omitempty"`
	ErrorCode    int    `json:"error_code,omitempty"`
	ModuleName   string `json:"module_name,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
}

func SetContext(cfg Configuration) context.Context {
	return context.WithValue(
		context.Background(), ContextBaseAuth, BasicAuth{
			Username: cfg.Username,
			Password: cfg.Password,
		},
	)
}
func (c *NSXTClient) initSession() (err error) {
	var (
		path     string
		response *Response
	)

	path = fmt.Sprintf("%s/api/session/create", c.cfg.BasePath)

	if response, err = c.client.Request(
		POST, path, &Options{
			Headers: map[string]string{
				"Accept":       "application/json",
				"Content-Type": "application/x-www-form-urlencoded",
				"User-Agent":   "",
			},
			Ctx: c.Context,
		},
	); err != nil {
		return
	}

	if c.defaultHeader == nil {
		c.defaultHeader = make(map[string]string)
	}

	for k, v := range response.Header {
		if strings.ToLower(XSRF_TOKEN) == strings.ToLower(k) {
			c.defaultHeader[XSRF_TOKEN] = v[0]
		}
	}
	return
}

func (c *NSXTClient) RequestAPI(method, path string, options *Options) (r *Response, err error) {
	path = fmt.Sprintf("%s%s", c.cfg.BasePath, path)
	return c.client.Request(method, path, options)
}

func NewNSXTClient(cfg Configuration) (c *NSXTClient, err error) {
	c = &NSXTClient{}

	if c.client, err = NewClient(); err != nil {
		return
	}

	c.cfg = cfg
	c.Context = SetContext(cfg)
	c.common.client = c

	if err = c.initSession(); err != nil {
		return
	}

	c.ApiServiceNSXT = (*nsxt.ApiServiceNSXT)(&c.common)

	return
}
