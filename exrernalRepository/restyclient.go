package exrernalRepository

import (
	jsonitr "github.com/json-iterator/go"
	"gopkg.in/resty.v1"
	"time"
)

// MakeDefaultRestyClient creates and initialises a resty instance.
func MakeDefaultRestyClient(hostURL string, timeout int64) *resty.Client {
	// Things we set on the client are applicable for all requests
	client := resty.New()
	// We can override all below settings and options at request level if we want to
	//--------------------------------------------------------------------------------
	// Host URL for all request. So we can use relative URL in the request
	client.SetHostURL(hostURL)

	if timeout == 0 {
		timeout = 60
	}

	timeOutDur := time.Duration(timeout) * time.Second
	client.SetTimeout(timeOutDur) // in secs

	// Enabling Content length value for all request
	client.SetContentLength(true)

	// Add header for `Accept application/json`
	client.SetHeaders(map[string]string{
		"Content-Type": "application/json; charset=utf-8",
		"Accept":       "application/json; charset=utf-8",
	})

	client.JSONMarshal = jsonitr.Marshal
	client.JSONUnmarshal = jsonitr.Unmarshal
	return client
}
