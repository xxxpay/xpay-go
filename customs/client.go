package customs

import (
	"fmt"
	"log"

	"github.com/xxxpay/xpay-go"
)

type Client struct {
	backend xpay.Backend
}

func NewClient(backend xpay.Backend) Client {
	return Client{backend: backend}
}

func (c Client) New(params *xpay.CustomsParams) (*xpay.Customs, error) {
	paramsString, _ := xpay.JsonEncode(params)
	customs := &xpay.Customs{}
	err := c.backend.Call("POST", "/customs", nil, paramsString, customs)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("New Customs error: %v\n", err)
		}
	}
	return customs, err
}

func (c Client) Get(Id string) (*xpay.Customs, error) {
	customs := &xpay.Customs{}
	err := c.backend.Call("GET", fmt.Sprintf("/customs/%s", Id), nil, nil, customs)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("Get Customs error: %v\n", err)
		}
	}
	return customs, err
}
