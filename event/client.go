package event

import (
	"log"
	"net/url"

	"github.com/xxxpay/xpay-go"
)

type Client struct {
	backend xpay.Backend
}

func NewClient(backend xpay.Backend) Client {
	return Client{backend: backend}
}

func (c Client) Get(id string) (*xpay.Event, error) {
	var body *url.Values
	body = &url.Values{}
	eve := &xpay.Event{}
	err := c.backend.Call("GET", "/events/"+id, body, nil, eve)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("Get Event error: %v\n", err)
		}
	}
	return eve, err
}
