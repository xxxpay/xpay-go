package event

import (
	"log"
	"net/url"

	"github.com/xxxpay/xpay-go"
)

type Client struct {
	B   xpay.Backend
	Key string
}

func Get(id string) (*xpay.Event, error) {
	return getC().Get(id)
}

func (c Client) Get(id string) (*xpay.Event, error) {
	var body *url.Values
	body = &url.Values{}
	eve := &xpay.Event{}
	err := c.B.Call("GET", "/events/"+id, c.Key, body, nil, eve)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("Get Event error: %v\n", err)
		}
	}
	return eve, err
}

func getC() Client {
	return Client{xpay.GetBackend(xpay.APIBackend), xpay.Key}
}
